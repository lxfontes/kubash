/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

		Unless required by applicable law or agreed to in writing, software
		distributed under the License is distributed on an "AS IS" BASIS,
		WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
		See the License for the specific language governing permissions and
		limitations under the License.
*/

package main

import (
	"encoding/json"
	"net/http"
	"net/http/httputil"
	"os"
	"path"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/lxfontes/kubash/types"
	"k8s.io/client-go/1.5/pkg/api"
	//"k8s.io/client-go/1.5/pkg/fields"
	"k8s.io/client-go/1.5/pkg/runtime"
	"k8s.io/client-go/1.5/pkg/runtime/serializer"
	"k8s.io/client-go/1.5/pkg/watch"
	_ "k8s.io/client-go/1.5/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/1.5/rest"
	"k8s.io/client-go/1.5/tools/cache"
	"k8s.io/client-go/1.5/tools/clientcmd"
)

type loggingRoundTripper struct {
	inner http.RoundTripper
}

func (d *loggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	d.dumpRequest(req)
	res, err := d.inner.RoundTrip(req)
	d.dumpResponse(res)
	return res, err
}

func (d *loggingRoundTripper) dumpRequest(r *http.Request) {
	dump, err := httputil.DumpRequestOut(r, true)
	if err != nil {
		log.WithError(err).Error("oops")
	}
	log.WithField("section", "request").Infof(string(dump))
}

func (d *loggingRoundTripper) dumpResponse(r *http.Response) {
	dump, err := httputil.DumpResponse(r, true)
	if err != nil {
		log.WithError(err).Error("oops")
	}
	log.WithField("section", "response").Infof(string(dump))
}

type testLW struct {
	ListFunc  func(options api.ListOptions) (runtime.Object, error)
	WatchFunc func(options api.ListOptions) (watch.Interface, error)
}

func (t *testLW) List(options api.ListOptions) (runtime.Object, error) {
	return t.ListFunc(options)
}
func (t *testLW) Watch(options api.ListOptions) (watch.Interface, error) {
	return t.WatchFunc(options)
}

func main() {
	// creates the in-cluster config
	config, err := clientcmd.BuildConfigFromFlags("", path.Join(os.Getenv("HOME"), ".kube", "config"))
	if err != nil {
		panic(err.Error())
	}

	config.WrapTransport = func(rt http.RoundTripper) http.RoundTripper {
		return &loggingRoundTripper{rt}
	}

	config.GroupVersion = &types.GroupVersion
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: api.Codecs}
	config.APIPath = "/apis"

	client, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err.Error())
	}

	source := &testLW{
		WatchFunc: func(options api.ListOptions) (watch.Interface, error) {
			log.WithField("stage", "watch").Warnf("%#v", options)
			return watch.NewFake(), nil
		},
		ListFunc: func(options api.ListOptions) (runtime.Object, error) {
			log.WithField("stage", "list").Warn(options)
			l := types.RunItemList{}
			raw, err := client.Get().Resource("runitems").Do().Raw()
			if err != nil {
				return nil, err
			}
			err = json.Unmarshal(raw, &l)
			return &l, err
		},
	}

	handler := func(obj interface{}) {
		ev := obj.(*types.RunItem)
		log.WithField("script", ev.Script).Warnf("%v", ev.Arguments)
	}

	store, c := cache.NewInformer(
		source,
		&types.RunItem{},
		30*time.Second,
		cache.ResourceEventHandlerFuncs{
			AddFunc:    handler,
			DeleteFunc: handler,
		})

	stop := make(chan struct{})
	c.Run(stop)
	log.Warn(store, c)
}
