package tool_json

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
	"reflect"
)

func Int64ToStringCodec() jsoniter.Extension {
	return jsoniter.EncoderExtension{
		reflect2.DefaultTypeOfKind(reflect.Int64): &int64StringCodec{},
	}
}
