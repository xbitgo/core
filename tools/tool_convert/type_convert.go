package tool_convert

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

func MustToFloat64(data interface{}) float64 {
	var f float64
	if str, ok := data.(string); ok {
		float, err := strconv.ParseFloat(str, 64)
		if err == nil {
			return float
		}
	} else if intValue, ok := data.(int); ok {
		return float64(intValue)
	} else if intValue, ok := data.(int32); ok {
		return float64(intValue)
	} else if intValue, ok := data.(float64); ok {
		return intValue
	} else if intValue, ok := data.(int64); ok {
		return float64(intValue)
	}
	return f
}

func MustToInt64(value interface{}) int64 {
	v := indirect(reflect.ValueOf(value))
	switch v.Kind() {
	case reflect.String:
		x, _ := strconv.ParseInt(v.String(), 10, 64)
		return x
	case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
		return v.Int()
	case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
		return int64(v.Uint())
	case reflect.Float32, reflect.Float64:
		return int64(v.Float())
	case reflect.Bool:
		boolValue := v.Bool()
		if boolValue {
			return 1
		}
		return 0
	default:
		return 0
	}
}

func MustToBool(value interface{}) bool {
	v := indirect(reflect.ValueOf(value))
	switch v.Kind() {
	case reflect.String:
		x, _ := strconv.ParseBool(v.String())
		return x
	case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
		return v.Int() != 0
	case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
		return v.Uint() != 0
	case reflect.Float32, reflect.Float64:
		return int64(v.Float()) != 0
	case reflect.Bool:
		return v.Bool()
	default:
		return false
	}
}

func indirect(value reflect.Value) reflect.Value {
	for {
		if value.Kind() != reflect.Ptr && value.Kind() != reflect.Interface {
			return value
		}
		if value.IsNil() {
			return value
		}
		value = value.Elem()
	}
}

// StringToInt32Slice :
func StringToInt32Slice(s string, sep string) (ret []int32) {
	tokens := strings.Split(s, sep)
	for _, k := range tokens {
		i, err := strconv.ParseInt(k, 10, 64)
		if err != nil {
			return nil
		}
		ret = append(ret, int32(i))
	}
	return
}

// ToString convert some type to string
// []string{"a","b"} => "a,b"
// []int{1,2} => "1,2"
// []string{} => ""
func ToString(v interface{}) string {
	switch x := v.(type) {
	case string:
		return x
	case int64:
		return strconv.FormatInt(x, 10)
	case int:
		return strconv.FormatInt(int64(x), 10)
	case int32:
		return strconv.FormatInt(int64(x), 10)
	case int16:
		return strconv.FormatInt(int64(x), 10)
	case int8:
		return strconv.FormatInt(int64(x), 10)
	case uint:
		return strconv.FormatInt(int64(x), 10)
	case uint64:
		return strconv.FormatInt(int64(x), 10)
	case uint32:
		return strconv.FormatInt(int64(x), 10)
	case uint16:
		return strconv.FormatInt(int64(x), 10)
	case uint8:
		return strconv.FormatInt(int64(x), 10)
	case float64:
		return strconv.FormatFloat(x, 'f', -1, 64)
	case json.Number:
		return x.String()
	case []string:
		return strings.Join(x, ",")
	case []int, []int64, []int32, []int16, []int8, []uint, []uint64, []uint32, []uint16, []uint8, []float64:
		return ToString(ToStringSlice(x))
	case map[string]interface{}:
		str, _ := jsoniter.MarshalToString(x)
		return str
	default:
		return fmt.Sprint(v)
	}
}

// ToStringSlice .
func ToStringSlice(v interface{}) []string {
	r := reflect.ValueOf(v)
	if r.Kind() == reflect.Slice {
		s := make([]string, r.Len())
		for i := 0; i < r.Len(); i++ {
			s[i] = ToString(r.Index(i).Interface())
		}
		return s
	}
	return nil
}
