package tool_reflect

import (
	"fmt"
	"reflect"

	"github.com/xbitgo/core/tools/tool_convert"
)

func IndirectValue(value interface{}) reflect.Value {
	return indirect(reflect.ValueOf(value))
}

func IsSlice(value interface{}) error {
	if indirect(reflect.ValueOf(value)).Kind() != reflect.Slice {
		return fmt.Errorf("value is not slice")
	}
	return nil
}

func IsString(value interface{}) error {
	if indirect(reflect.ValueOf(value)).Kind() != reflect.String {
		return fmt.Errorf("value is not string")
	}
	return nil
}

func GetSlice(value interface{}) []interface{} {
	if err := IsSlice(value); err != nil {
		return []interface{}{}
	}
	v := indirect(reflect.ValueOf(value))
	result := make([]interface{}, 0, v.Len())

	for i := 0; i != v.Len(); i++ {
		result = append(result, v.Index(i).Interface())
	}

	return result
}

func AppendSlice(slice []interface{}, value interface{}) []interface{} {
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Slice {
		if slice == nil {
			slice = make([]interface{}, 0)
		}
		for i := 0; i != v.Len(); i++ {
			slice = append(slice, v.Index(i).Interface())
		}
	} else {
		slice = append(slice, v.Interface())
	}
	return slice
}

func EqualTo(value1 interface{}, value2 interface{}) (bool, error) {
	v1 := indirect(reflect.ValueOf(value1))
	v2 := indirect(reflect.ValueOf(value2))

	switch v1.Kind() {
	case reflect.String:
		v2Str := tool_convert.ToString(v2.Interface())
		return v2Str == v1.String(), nil
	case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
		v2Int := tool_convert.MustToInt64(v2.Interface())
		return v2Int == v1.Int(), nil
	case reflect.Float32, reflect.Float64:
		v1Val := tool_convert.MustToFloat64(v1.Interface())
		v2Val := tool_convert.MustToFloat64(v2.Interface())
		return v1Val == v2Val, nil
	default:
		return false, fmt.Errorf("EqualTo can not process type:%s", v1.Type())
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
