package tool_convert

import (
	"reflect"
	"strconv"
	"unsafe"
)

// StringToInt @Converter implicit
func StringToInt(str string) int {
	result, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return int(result)
}

// StringToInt32 @Converter implicit
func StringToInt32(str string) int32 {
	result, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return int32(result)
}

// StringToInt64 @Converter implicit
func StringToInt64(str string) int64 {
	result, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return result
}

// UnsafeStringToByte .
func UnsafeStringToByte(s string) []byte {
	return (*[0x7fff0000]byte)(unsafe.Pointer(
		(*reflect.StringHeader)(unsafe.Pointer(&s)).Data),
	)[:len(s):len(s)]
}

// StringToInt16 @Converter implicit
func StringToInt16(s string) int16 {
	if s == "" {
		return 0
	}
	TmpInt, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}

	return int16(TmpInt)
}

// StringToUint @Converter implicit
func StringToUint(s string) uint {
	if s == "" {
		return 0
	}

	TmpInt, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}

	return uint(TmpInt)
}

// StringToUint16 @Converter implicit
func StringToUint16(s string) uint16 {
	if s == "" {
		return 0
	}

	TmpInt, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}

	return uint16(TmpInt)
}

// StringToUint32 @Converter implicit
func StringToUint32(s string) uint32 {
	if s == "" {
		return 0
	}

	TmpInt, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}

	return uint32(TmpInt)
}

// StringToUint64 @Converter implicit
func StringToUint64(s string) uint64 {
	if s == "" {
		return 0
	}

	TmpInt, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}

	return uint64(TmpInt)
}

// StringToFloat64 @Converter implicit
func StringToFloat64(s string) float64 {
	if s == "" {
		return 0
	}
	tmp, _ := strconv.ParseFloat(s, 64)
	return tmp
}

// IntToString @Converter implicit
func IntToString(val int) string {
	return strconv.FormatInt(int64(val), 10)
}

// Int32ToString @Converter implicit
func Int32ToString(val int32) string {
	return strconv.FormatInt(int64(val), 10)
}

// Int64ToString @Converter implicit
func Int64ToString(val int64) string {
	return strconv.FormatInt(val, 10)
}

// Int16ToString @Converter implicit
func Int16ToString(a int16) string {
	return strconv.FormatInt(int64(a), 10)
}

// Uint64ToString @Converter implicit
func Uint64ToString(a uint64) string {
	return strconv.FormatUint(a, 10)
}

// Uint32ToString @Converter implicit
func Uint32ToString(a uint32) string {
	return strconv.FormatUint(uint64(a), 10)
}

// Uint16ToString @Converter implicit
func Uint16ToString(a uint16) string {
	return strconv.FormatUint(uint64(a), 10)
}

// Float64ToString @Converter implicit
func Float64ToString(a float64) string {
	return strconv.FormatFloat(a, 'f', -1, 64)
}
