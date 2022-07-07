package tool_json

import (
	"strconv"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

type int64StringCodec struct{}

func (codec *int64StringCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	if !iter.ReadNil() {
		*((*int64)(ptr)) = iter.ReadInt64()
	}
}

func (codec *int64StringCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	x := *((*int64)(ptr))
	stream.WriteString(strconv.FormatInt(x, 10))
}

func (codec *int64StringCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return *((*int64)(ptr)) == 0
}
