package tool_slice

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/xbitgo/core/tools/tool_convert"
)

func ImplodeIntSlice(slice []int, delimiter string) string {
	strApp := make([]string, 0)
	for _, singleId := range slice {
		strApp = append(strApp, tool_convert.IntToString(singleId))
	}
	return strings.Join(strApp, delimiter)
}

func ImplodeInt32Slice(slice []int32, delimiter string) string {
	strApp := make([]string, 0)
	for _, singleId := range slice {
		strApp = append(strApp, tool_convert.Int32ToString(singleId))
	}
	return strings.Join(strApp, delimiter)
}

func ImplodeInt64Slice(slice []int64, delimiter string) string {
	strApp := make([]string, 0)
	for _, singleId := range slice {
		strApp = append(strApp, tool_convert.Int64ToString(singleId))
	}
	return strings.Join(strApp, delimiter)
}

// ExplodeInt64 将一个 string 根据 seq 拆分为 int64数组
func ExplodeInt64(rawStr string, seq string) []int64 {
	res := make([]int64, 0)
	if len(rawStr) == 0 {
		return res
	}

	strArr := strings.Split(rawStr, seq)
	for _, str := range strArr {
		if val, err := strconv.Atoi(str); err == nil {
			res = append(res, int64(val))
		} else {
			return make([]int64, 0)
		}
	}
	return res
}

// ExplodeInt32 .
func ExplodeInt32(rawStr string, seq string) []int32 {
	res := make([]int32, 0)
	if len(rawStr) == 0 {
		return res
	}

	strArr := strings.Split(rawStr, seq)
	for _, str := range strArr {
		if val, err := strconv.Atoi(str); err == nil {
			res = append(res, int32(val))
		} else {
			return make([]int32, 0)
		}
	}
	return res
}

// ExplodeStr .
func ExplodeStr(rawStr string, seq string) []string {
	res := make([]string, 0)
	if len(rawStr) == 0 {
		return res
	}

	strArr := strings.Split(rawStr, seq)
	for _, str := range strArr {
		if str != "" {
			res = append(res, str)
		}
	}
	return res
}

// Implode .
func Implode(list interface{}, seq string) string {
	listValue := reflect.Indirect(reflect.ValueOf(list))
	if listValue.Kind() != reflect.Slice {
		return ""
	}

	count := listValue.Len()
	listStr := make([]string, 0, count)
	for i := 0; i < count; i++ {
		v := listValue.Index(i)
		if str, err := getValue(v); err == nil {
			listStr = append(listStr, str)
		}
	}
	return strings.Join(listStr, seq)
}

func getValue(value reflect.Value) (res string, err error) {
	switch value.Kind() {
	case reflect.Ptr:
		res, err = getValue(value.Elem())
	default:
		res = fmt.Sprint(value.Interface())
	}
	return
}
