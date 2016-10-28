// ************************************************************
// DO NOT EDIT.
// THIS FILE IS AUTO-GENERATED BY codecgen.
// ************************************************************

package types

import (
	"errors"
	"fmt"
	codec1978 "github.com/ugorji/go/codec"
	pkg2_api "k8s.io/client-go/1.5/pkg/api"
	pkg1_unversioned "k8s.io/client-go/1.5/pkg/api/unversioned"
	pkg3_types "k8s.io/client-go/1.5/pkg/types"
	"reflect"
	"runtime"
	time "time"
)

const (
	// ----- content types ----
	codecSelferC_UTF86855 = 1
	codecSelferC_RAW6855  = 0
	// ----- value types used ----
	codecSelferValueTypeArray6855 = 10
	codecSelferValueTypeMap6855   = 9
	// ----- containerStateValues ----
	codecSelfer_containerMapKey6855    = 2
	codecSelfer_containerMapValue6855  = 3
	codecSelfer_containerMapEnd6855    = 4
	codecSelfer_containerArrayElem6855 = 6
	codecSelfer_containerArrayEnd6855  = 7
)

var (
	codecSelferBitsize6855                         = uint8(reflect.TypeOf(uint(0)).Bits())
	codecSelferOnlyMapOrArrayEncodeToStructErr6855 = errors.New(`only encoded map or array can be decoded into a struct`)
)

type codecSelfer6855 struct{}

func init() {
	if codec1978.GenVersion != 5 {
		_, file, _, _ := runtime.Caller(0)
		err := fmt.Errorf("codecgen version mismatch: current: %v, need %v. Re-generate file: %v",
			5, codec1978.GenVersion, file)
		panic(err)
	}
	if false { // reference the types, but skip this branch at build/run time
		var v0 pkg2_api.ObjectMeta
		var v1 pkg1_unversioned.TypeMeta
		var v2 pkg3_types.UID
		var v3 time.Time
		_, _, _, _ = v0, v1, v2, v3
	}
}

func (x *RunItemList) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer6855
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			var yyq2 [4]bool
			_, _, _ = yysep2, yyq2, yy2arr2
			const yyr2 bool = false
			yyq2[0] = x.Kind != ""
			yyq2[1] = x.APIVersion != ""
			yyq2[2] = true
			var yynn2 int
			if yyr2 || yy2arr2 {
				r.EncodeArrayStart(4)
			} else {
				yynn2 = 1
				for _, b := range yyq2 {
					if b {
						yynn2++
					}
				}
				r.EncodeMapStart(yynn2)
				yynn2 = 0
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem6855)
				if yyq2[0] {
					yym4 := z.EncBinary()
					_ = yym4
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF86855, string(x.Kind))
					}
				} else {
					r.EncodeString(codecSelferC_UTF86855, "")
				}
			} else {
				if yyq2[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey6855)
					r.EncodeString(codecSelferC_UTF86855, string("kind"))
					z.EncSendContainerState(codecSelfer_containerMapValue6855)
					yym5 := z.EncBinary()
					_ = yym5
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF86855, string(x.Kind))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem6855)
				if yyq2[1] {
					yym7 := z.EncBinary()
					_ = yym7
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF86855, string(x.APIVersion))
					}
				} else {
					r.EncodeString(codecSelferC_UTF86855, "")
				}
			} else {
				if yyq2[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey6855)
					r.EncodeString(codecSelferC_UTF86855, string("apiVersion"))
					z.EncSendContainerState(codecSelfer_containerMapValue6855)
					yym8 := z.EncBinary()
					_ = yym8
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF86855, string(x.APIVersion))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem6855)
				if yyq2[2] {
					yy10 := &x.ListMeta
					yym11 := z.EncBinary()
					_ = yym11
					if false {
					} else if z.HasExtensions() && z.EncExt(yy10) {
					} else {
						z.EncFallback(yy10)
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[2] {
					z.EncSendContainerState(codecSelfer_containerMapKey6855)
					r.EncodeString(codecSelferC_UTF86855, string("metadata"))
					z.EncSendContainerState(codecSelfer_containerMapValue6855)
					yy12 := &x.ListMeta
					yym13 := z.EncBinary()
					_ = yym13
					if false {
					} else if z.HasExtensions() && z.EncExt(yy12) {
					} else {
						z.EncFallback(yy12)
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem6855)
				if x.Items == nil {
					r.EncodeNil()
				} else {
					yym15 := z.EncBinary()
					_ = yym15
					if false {
					} else {
						h.encSliceRunItem(([]RunItem)(x.Items), e)
					}
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey6855)
				r.EncodeString(codecSelferC_UTF86855, string("items"))
				z.EncSendContainerState(codecSelfer_containerMapValue6855)
				if x.Items == nil {
					r.EncodeNil()
				} else {
					yym16 := z.EncBinary()
					_ = yym16
					if false {
					} else {
						h.encSliceRunItem(([]RunItem)(x.Items), e)
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd6855)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd6855)
			}
		}
	}
}

func (x *RunItemList) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer6855
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap6855 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd6855)
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray6855 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd6855)
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr6855)
		}
	}
}

func (x *RunItemList) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer6855
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys3Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys3Slc
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey6855)
		yys3Slc = r.DecodeBytes(yys3Slc, true, true)
		yys3 := string(yys3Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue6855)
		switch yys3 {
		case "kind":
			if r.TryDecodeAsNil() {
				x.Kind = ""
			} else {
				yyv4 := &x.Kind
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} else {
					*((*string)(yyv4)) = r.DecodeString()
				}
			}
		case "apiVersion":
			if r.TryDecodeAsNil() {
				x.APIVersion = ""
			} else {
				yyv6 := &x.APIVersion
				yym7 := z.DecBinary()
				_ = yym7
				if false {
				} else {
					*((*string)(yyv6)) = r.DecodeString()
				}
			}
		case "metadata":
			if r.TryDecodeAsNil() {
				x.ListMeta = pkg1_unversioned.ListMeta{}
			} else {
				yyv8 := &x.ListMeta
				yym9 := z.DecBinary()
				_ = yym9
				if false {
				} else if z.HasExtensions() && z.DecExt(yyv8) {
				} else {
					z.DecFallback(yyv8, false)
				}
			}
		case "items":
			if r.TryDecodeAsNil() {
				x.Items = nil
			} else {
				yyv10 := &x.Items
				yym11 := z.DecBinary()
				_ = yym11
				if false {
				} else {
					h.decSliceRunItem((*[]RunItem)(yyv10), d)
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		} // end switch yys3
	} // end for yyj3
	z.DecSendContainerState(codecSelfer_containerMapEnd6855)
}

func (x *RunItemList) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer6855
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj12 int
	var yyb12 bool
	var yyhl12 bool = l >= 0
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} else {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd6855)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem6855)
	if r.TryDecodeAsNil() {
		x.Kind = ""
	} else {
		yyv13 := &x.Kind
		yym14 := z.DecBinary()
		_ = yym14
		if false {
		} else {
			*((*string)(yyv13)) = r.DecodeString()
		}
	}
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} else {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd6855)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem6855)
	if r.TryDecodeAsNil() {
		x.APIVersion = ""
	} else {
		yyv15 := &x.APIVersion
		yym16 := z.DecBinary()
		_ = yym16
		if false {
		} else {
			*((*string)(yyv15)) = r.DecodeString()
		}
	}
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} else {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd6855)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem6855)
	if r.TryDecodeAsNil() {
		x.ListMeta = pkg1_unversioned.ListMeta{}
	} else {
		yyv17 := &x.ListMeta
		yym18 := z.DecBinary()
		_ = yym18
		if false {
		} else if z.HasExtensions() && z.DecExt(yyv17) {
		} else {
			z.DecFallback(yyv17, false)
		}
	}
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} else {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd6855)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem6855)
	if r.TryDecodeAsNil() {
		x.Items = nil
	} else {
		yyv19 := &x.Items
		yym20 := z.DecBinary()
		_ = yym20
		if false {
		} else {
			h.decSliceRunItem((*[]RunItem)(yyv19), d)
		}
	}
	for {
		yyj12++
		if yyhl12 {
			yyb12 = yyj12 > l
		} else {
			yyb12 = r.CheckBreak()
		}
		if yyb12 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem6855)
		z.DecStructFieldNotFound(yyj12-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd6855)
}

func (x *RunItem) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer6855
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			var yyq2 [5]bool
			_, _, _ = yysep2, yyq2, yy2arr2
			const yyr2 bool = false
			yyq2[0] = x.Kind != ""
			yyq2[1] = x.APIVersion != ""
			yyq2[2] = true
			yyq2[4] = len(x.Arguments) != 0
			var yynn2 int
			if yyr2 || yy2arr2 {
				r.EncodeArrayStart(5)
			} else {
				yynn2 = 1
				for _, b := range yyq2 {
					if b {
						yynn2++
					}
				}
				r.EncodeMapStart(yynn2)
				yynn2 = 0
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem6855)
				if yyq2[0] {
					yym4 := z.EncBinary()
					_ = yym4
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF86855, string(x.Kind))
					}
				} else {
					r.EncodeString(codecSelferC_UTF86855, "")
				}
			} else {
				if yyq2[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey6855)
					r.EncodeString(codecSelferC_UTF86855, string("kind"))
					z.EncSendContainerState(codecSelfer_containerMapValue6855)
					yym5 := z.EncBinary()
					_ = yym5
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF86855, string(x.Kind))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem6855)
				if yyq2[1] {
					yym7 := z.EncBinary()
					_ = yym7
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF86855, string(x.APIVersion))
					}
				} else {
					r.EncodeString(codecSelferC_UTF86855, "")
				}
			} else {
				if yyq2[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey6855)
					r.EncodeString(codecSelferC_UTF86855, string("apiVersion"))
					z.EncSendContainerState(codecSelfer_containerMapValue6855)
					yym8 := z.EncBinary()
					_ = yym8
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF86855, string(x.APIVersion))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem6855)
				if yyq2[2] {
					yy10 := &x.ObjectMeta
					yym11 := z.EncBinary()
					_ = yym11
					if false {
					} else if z.HasExtensions() && z.EncExt(yy10) {
					} else {
						z.EncFallback(yy10)
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[2] {
					z.EncSendContainerState(codecSelfer_containerMapKey6855)
					r.EncodeString(codecSelferC_UTF86855, string("metadata"))
					z.EncSendContainerState(codecSelfer_containerMapValue6855)
					yy12 := &x.ObjectMeta
					yym13 := z.EncBinary()
					_ = yym13
					if false {
					} else if z.HasExtensions() && z.EncExt(yy12) {
					} else {
						z.EncFallback(yy12)
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem6855)
				yym15 := z.EncBinary()
				_ = yym15
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF86855, string(x.Script))
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey6855)
				r.EncodeString(codecSelferC_UTF86855, string("script"))
				z.EncSendContainerState(codecSelfer_containerMapValue6855)
				yym16 := z.EncBinary()
				_ = yym16
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF86855, string(x.Script))
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem6855)
				if yyq2[4] {
					if x.Arguments == nil {
						r.EncodeNil()
					} else {
						yym18 := z.EncBinary()
						_ = yym18
						if false {
						} else {
							z.F.EncSliceStringV(x.Arguments, false, e)
						}
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[4] {
					z.EncSendContainerState(codecSelfer_containerMapKey6855)
					r.EncodeString(codecSelferC_UTF86855, string("arguments"))
					z.EncSendContainerState(codecSelfer_containerMapValue6855)
					if x.Arguments == nil {
						r.EncodeNil()
					} else {
						yym19 := z.EncBinary()
						_ = yym19
						if false {
						} else {
							z.F.EncSliceStringV(x.Arguments, false, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd6855)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd6855)
			}
		}
	}
}

func (x *RunItem) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer6855
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap6855 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd6855)
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray6855 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd6855)
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr6855)
		}
	}
}

func (x *RunItem) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer6855
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys3Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys3Slc
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey6855)
		yys3Slc = r.DecodeBytes(yys3Slc, true, true)
		yys3 := string(yys3Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue6855)
		switch yys3 {
		case "kind":
			if r.TryDecodeAsNil() {
				x.Kind = ""
			} else {
				yyv4 := &x.Kind
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} else {
					*((*string)(yyv4)) = r.DecodeString()
				}
			}
		case "apiVersion":
			if r.TryDecodeAsNil() {
				x.APIVersion = ""
			} else {
				yyv6 := &x.APIVersion
				yym7 := z.DecBinary()
				_ = yym7
				if false {
				} else {
					*((*string)(yyv6)) = r.DecodeString()
				}
			}
		case "metadata":
			if r.TryDecodeAsNil() {
				x.ObjectMeta = pkg2_api.ObjectMeta{}
			} else {
				yyv8 := &x.ObjectMeta
				yym9 := z.DecBinary()
				_ = yym9
				if false {
				} else if z.HasExtensions() && z.DecExt(yyv8) {
				} else {
					z.DecFallback(yyv8, false)
				}
			}
		case "script":
			if r.TryDecodeAsNil() {
				x.Script = ""
			} else {
				yyv10 := &x.Script
				yym11 := z.DecBinary()
				_ = yym11
				if false {
				} else {
					*((*string)(yyv10)) = r.DecodeString()
				}
			}
		case "arguments":
			if r.TryDecodeAsNil() {
				x.Arguments = nil
			} else {
				yyv12 := &x.Arguments
				yym13 := z.DecBinary()
				_ = yym13
				if false {
				} else {
					z.F.DecSliceStringX(yyv12, false, d)
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		} // end switch yys3
	} // end for yyj3
	z.DecSendContainerState(codecSelfer_containerMapEnd6855)
}

func (x *RunItem) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer6855
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj14 int
	var yyb14 bool
	var yyhl14 bool = l >= 0
	yyj14++
	if yyhl14 {
		yyb14 = yyj14 > l
	} else {
		yyb14 = r.CheckBreak()
	}
	if yyb14 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd6855)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem6855)
	if r.TryDecodeAsNil() {
		x.Kind = ""
	} else {
		yyv15 := &x.Kind
		yym16 := z.DecBinary()
		_ = yym16
		if false {
		} else {
			*((*string)(yyv15)) = r.DecodeString()
		}
	}
	yyj14++
	if yyhl14 {
		yyb14 = yyj14 > l
	} else {
		yyb14 = r.CheckBreak()
	}
	if yyb14 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd6855)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem6855)
	if r.TryDecodeAsNil() {
		x.APIVersion = ""
	} else {
		yyv17 := &x.APIVersion
		yym18 := z.DecBinary()
		_ = yym18
		if false {
		} else {
			*((*string)(yyv17)) = r.DecodeString()
		}
	}
	yyj14++
	if yyhl14 {
		yyb14 = yyj14 > l
	} else {
		yyb14 = r.CheckBreak()
	}
	if yyb14 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd6855)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem6855)
	if r.TryDecodeAsNil() {
		x.ObjectMeta = pkg2_api.ObjectMeta{}
	} else {
		yyv19 := &x.ObjectMeta
		yym20 := z.DecBinary()
		_ = yym20
		if false {
		} else if z.HasExtensions() && z.DecExt(yyv19) {
		} else {
			z.DecFallback(yyv19, false)
		}
	}
	yyj14++
	if yyhl14 {
		yyb14 = yyj14 > l
	} else {
		yyb14 = r.CheckBreak()
	}
	if yyb14 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd6855)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem6855)
	if r.TryDecodeAsNil() {
		x.Script = ""
	} else {
		yyv21 := &x.Script
		yym22 := z.DecBinary()
		_ = yym22
		if false {
		} else {
			*((*string)(yyv21)) = r.DecodeString()
		}
	}
	yyj14++
	if yyhl14 {
		yyb14 = yyj14 > l
	} else {
		yyb14 = r.CheckBreak()
	}
	if yyb14 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd6855)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem6855)
	if r.TryDecodeAsNil() {
		x.Arguments = nil
	} else {
		yyv23 := &x.Arguments
		yym24 := z.DecBinary()
		_ = yym24
		if false {
		} else {
			z.F.DecSliceStringX(yyv23, false, d)
		}
	}
	for {
		yyj14++
		if yyhl14 {
			yyb14 = yyj14 > l
		} else {
			yyb14 = r.CheckBreak()
		}
		if yyb14 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem6855)
		z.DecStructFieldNotFound(yyj14-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd6855)
}

func (x codecSelfer6855) encSliceRunItem(v []RunItem, e *codec1978.Encoder) {
	var h codecSelfer6855
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	r.EncodeArrayStart(len(v))
	for _, yyv1 := range v {
		z.EncSendContainerState(codecSelfer_containerArrayElem6855)
		yy2 := &yyv1
		yy2.CodecEncodeSelf(e)
	}
	z.EncSendContainerState(codecSelfer_containerArrayEnd6855)
}

func (x codecSelfer6855) decSliceRunItem(v *[]RunItem, d *codec1978.Decoder) {
	var h codecSelfer6855
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r

	yyv1 := *v
	yyh1, yyl1 := z.DecSliceHelperStart()
	var yyc1 bool
	_ = yyc1
	if yyl1 == 0 {
		if yyv1 == nil {
			yyv1 = []RunItem{}
			yyc1 = true
		} else if len(yyv1) != 0 {
			yyv1 = yyv1[:0]
			yyc1 = true
		}
	} else if yyl1 > 0 {
		var yyrr1, yyrl1 int
		var yyrt1 bool
		_, _ = yyrl1, yyrt1
		yyrr1 = yyl1 // len(yyv1)
		if yyl1 > cap(yyv1) {

			yyrg1 := len(yyv1) > 0
			yyv21 := yyv1
			yyrl1, yyrt1 = z.DecInferLen(yyl1, z.DecBasicHandle().MaxInitLen, 296)
			if yyrt1 {
				if yyrl1 <= cap(yyv1) {
					yyv1 = yyv1[:yyrl1]
				} else {
					yyv1 = make([]RunItem, yyrl1)
				}
			} else {
				yyv1 = make([]RunItem, yyrl1)
			}
			yyc1 = true
			yyrr1 = len(yyv1)
			if yyrg1 {
				copy(yyv1, yyv21)
			}
		} else if yyl1 != len(yyv1) {
			yyv1 = yyv1[:yyl1]
			yyc1 = true
		}
		yyj1 := 0
		for ; yyj1 < yyrr1; yyj1++ {
			yyh1.ElemContainerState(yyj1)
			if r.TryDecodeAsNil() {
				yyv1[yyj1] = RunItem{}
			} else {
				yyv2 := &yyv1[yyj1]
				yyv2.CodecDecodeSelf(d)
			}

		}
		if yyrt1 {
			for ; yyj1 < yyl1; yyj1++ {
				yyv1 = append(yyv1, RunItem{})
				yyh1.ElemContainerState(yyj1)
				if r.TryDecodeAsNil() {
					yyv1[yyj1] = RunItem{}
				} else {
					yyv3 := &yyv1[yyj1]
					yyv3.CodecDecodeSelf(d)
				}

			}
		}

	} else {
		yyj1 := 0
		for ; !r.CheckBreak(); yyj1++ {

			if yyj1 >= len(yyv1) {
				yyv1 = append(yyv1, RunItem{}) // var yyz1 RunItem
				yyc1 = true
			}
			yyh1.ElemContainerState(yyj1)
			if yyj1 < len(yyv1) {
				if r.TryDecodeAsNil() {
					yyv1[yyj1] = RunItem{}
				} else {
					yyv4 := &yyv1[yyj1]
					yyv4.CodecDecodeSelf(d)
				}

			} else {
				z.DecSwallow()
			}

		}
		if yyj1 < len(yyv1) {
			yyv1 = yyv1[:yyj1]
			yyc1 = true
		} else if yyj1 == 0 && yyv1 == nil {
			yyv1 = []RunItem{}
			yyc1 = true
		}
	}
	yyh1.End()
	if yyc1 {
		*v = yyv1
	}
}