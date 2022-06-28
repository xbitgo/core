package tool_reflect

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/xbitgo/core/tools/tool_str"
)

func MustExtractData(field string, data interface{}) interface{} {
	result, err := ExtractData(field, data)
	if err != nil {
		panic(err)
	}
	return result
}

func ExtractData(field string, data interface{}) (interface{}, error) {
	if field == "." || field == "" {
		return data, nil
	}

	field = strings.TrimLeft(field, ".")

	parts := strings.SplitN(field, ".", 2)
	field = parts[0]
	fieldToGo := ""
	if len(parts) > 1 {
		fieldToGo = parts[1]
	}

	indexField, err := strconv.ParseInt(field, 0, 32)
	if err == nil {
		result, err := ExtractDataIndex(int(indexField), data)
		if err != nil {
			return nil, err
		}

		return ExtractData(fieldToGo, result)
	} else {
		result, err := ExtractDataField(field, data)
		if err != nil {
			return nil, err
		}

		return ExtractData(fieldToGo, result)
	}
}

func ExtractDataIndex(index int, data interface{}) (interface{}, error) {
	v := indirect(reflect.ValueOf(data))

	if !v.IsValid() {
		return nil, fmt.Errorf("try to get index on nil value")
	}

	if v.Kind() != reflect.Slice {
		return nil, fmt.Errorf("try to get index field for non slice type:%s", v.Type())
	}

	if v.IsNil() {
		return nil, fmt.Errorf("try to get index from nil slice")
	}

	if v.Len() <= index || index < 0 {
		return nil, fmt.Errorf("try to get index out of bound, len:%d, index:%d", v.Len(), index)
	}

	return v.Index(index).Interface(), nil
}

func ExtractDataField(field string, data interface{}) (interface{}, error) {
	v := indirect(reflect.ValueOf(data))

	if !v.IsValid() {
		return nil, fmt.Errorf("try to get field on nil value")
	}

	switch v.Kind() {
	case reflect.Map:
		if v.IsNil() {
			return nil, fmt.Errorf("try to get field from nil map")
		}

		mv := v.MapIndex(reflect.ValueOf(field))
		if !mv.IsValid() {
			mv = v.MapIndex(reflect.ValueOf(tool_str.ToLFirst(field)))
			if !mv.IsValid() {
				return ExtractDataByGetMethod(field, data)
			}
		}

		return mv.Interface(), nil
	case reflect.Slice:
		if v.IsNil() {
			return nil, fmt.Errorf("try to get field from nil slice")
		}

		return ExtractDataByGetMethod(field, data)
	case reflect.Struct:
		mv := v.FieldByName(field)
		if !mv.IsValid() {
			return ExtractDataByGetMethod(field, data)
		}

		return mv.Interface(), nil
	default:
		return nil, fmt.Errorf("can not get field for type:%s", v.Type())
	}
}

func ExtractDataByGetMethod(field string, data interface{}) (interface{}, error) {
	v := reflect.ValueOf(data)

	if !v.IsValid() {
		return nil, fmt.Errorf("try to get by method on nil value")
	}

	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Slice || v.Kind() == reflect.Map {
		if v.IsNil() {
			return nil, fmt.Errorf("try to get by method on nil ptr")
		}
	} else if v.Kind() == reflect.Struct {

	} else {
		return nil, fmt.Errorf("try to get by method on wrong type:%s", v.Type())
	}

	methodName := fmt.Sprintf("%s", field)
	m := v.MethodByName(methodName)
	if !m.IsValid() {
		methodName = fmt.Sprintf("Get%s", field)
		m = v.MethodByName(methodName)
		if !m.IsValid() {
			return nil, fmt.Errorf("method %s, no exist in type:%s", methodName, v.Type())
		}
	}

	results := m.Call([]reflect.Value{})
	if len(results) != 1 {
		return nil, fmt.Errorf("method %s, result count err:%d, should be one", methodName, len(results))
	}

	return results[0].Interface(), nil
}
