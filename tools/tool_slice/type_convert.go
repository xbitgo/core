package tool_slice

import (
	"fmt"

	"github.com/xbitgo/core/tools/tool_convert"
)

// StringSliceToInterfaceSlice @Converter implicit
func StringSliceToInterfaceSlice(in []string) []interface{} {
	if in == nil {
		return nil
	}

	result := make([]interface{}, len(in))
	for i, v := range in {
		result[i] = v
	}
	return result
}

// StringSliceToIntSlice @Converter implicit
func StringSliceToIntSlice(in []string) []int {
	if in == nil {
		return nil
	}

	result := make([]int, len(in))
	for i, v := range in {
		result[i] = tool_convert.StringToInt(v)
	}
	return result
}

// StringSliceToInt32Slice @Converter implicit
func StringSliceToInt32Slice(in []string) []int32 {
	if in == nil {
		return nil
	}

	result := make([]int32, len(in))
	for i, v := range in {
		result[i] = tool_convert.StringToInt32(v)
	}
	return result
}

// StringSliceToInt64Slice @Converter implicit
func StringSliceToInt64Slice(in []string) []int64 {
	if in == nil {
		return nil
	}

	result := make([]int64, len(in))
	for i, v := range in {
		result[i] = tool_convert.StringToInt64(v)
	}
	return result
}

// IntSliceToInterfaceSlice @Converter implicit
func IntSliceToInterfaceSlice(in []int) []interface{} {
	if in == nil {
		return nil
	}

	result := make([]interface{}, len(in))
	for i, v := range in {
		result[i] = v
	}
	return result
}

// IntSliceToStringSlice @Converter implicit
func IntSliceToStringSlice(in []int) []string {
	if in == nil {
		return nil
	}

	result := make([]string, len(in))
	for i, v := range in {
		result[i] = fmt.Sprintf("%d", v)
	}
	return result
}

// IntSliceToInt32Slice @Converter implicit
func IntSliceToInt32Slice(in []int) []int32 {
	if in == nil {
		return nil
	}

	result := make([]int32, len(in))
	for i, v := range in {
		result[i] = int32(v)
	}
	return result
}

// IntSliceToInt64Slice @Converter implicit
func IntSliceToInt64Slice(in []int) []int64 {
	if in == nil {
		return nil
	}

	result := make([]int64, len(in))
	for i, v := range in {
		result[i] = int64(v)
	}
	return result
}

// Int32SliceToInterfaceSlice @Converter implicit
func Int32SliceToInterfaceSlice(in []int32) []interface{} {
	if in == nil {
		return nil
	}

	result := make([]interface{}, len(in))
	for i, v := range in {
		result[i] = v
	}
	return result
}

// Int32SliceToStringSlice @Converter implicit
func Int32SliceToStringSlice(in []int32) []string {
	if in == nil {
		return nil
	}

	result := make([]string, len(in))
	for i, v := range in {
		result[i] = fmt.Sprintf("%d", v)
	}
	return result
}

// Int32SliceToIntSlice @Converter implicit
func Int32SliceToIntSlice(in []int32) []int {
	if in == nil {
		return nil
	}

	result := make([]int, len(in))
	for i, v := range in {
		result[i] = int(v)
	}
	return result
}

// Int32SliceToInt64Slice @Converter implicit
func Int32SliceToInt64Slice(in []int32) []int64 {
	if in == nil {
		return nil
	}

	result := make([]int64, len(in))
	for i, v := range in {
		result[i] = int64(v)
	}
	return result
}

// Int64SliceToInterfaceSlice @Converter implicit
func Int64SliceToInterfaceSlice(in []int64) []interface{} {
	if in == nil {
		return nil
	}

	result := make([]interface{}, len(in))
	for i, v := range in {
		result[i] = v
	}
	return result
}

// Int64SliceToStringSlice @Converter implicit
func Int64SliceToStringSlice(in []int64) []string {
	if in == nil {
		return nil
	}

	result := make([]string, len(in))
	for i, v := range in {
		result[i] = fmt.Sprintf("%d", v)
	}
	return result
}

// Int64SliceToIntSlice @Converter implicit
func Int64SliceToIntSlice(in []int64) []int {
	if in == nil {
		return nil
	}

	result := make([]int, len(in))
	for i, v := range in {
		result[i] = int(v)
	}
	return result
}

// Int64SliceToInt32Slice @Converter implicit
func Int64SliceToInt32Slice(in []int64) []int32 {
	if in == nil {
		return nil
	}

	result := make([]int32, len(in))
	for i, v := range in {
		result[i] = int32(v)
	}
	return result
}
