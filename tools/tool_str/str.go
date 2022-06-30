package tool_str

import (
	"math/rand"
	"regexp"
	"strings"
	"unicode"

	"github.com/gofrs/uuid"
)

// RandString .
func RandString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := rand.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

// UUID .
func UUID() string {
	u1 := uuid.Must(uuid.NewV6())
	return u1.String()
}

func LastPwdStr(pwd string) string {
	pwdTmp := strings.Split(pwd, "/")
	return pwdTmp[len(pwdTmp)-1]
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func ToCamelCase(str string) string {
	convertedParts := strings.Split(str, "_")
	for index, part := range convertedParts {
		if index > 0 {
			convertedParts[index] = ToUFirst(part)
			continue
		}
		convertedParts[index] = part
	}
	return strings.Join(convertedParts, "")
}

func SnakeCaseToCamelCaseUF(str string) string {
	convertedParts := strings.Split(str, "_")
	for index, part := range convertedParts {
		convertedParts[index] = ToUFirst(part)
	}
	return strings.Join(convertedParts, "")
}

func UFirst(str string) bool {
	for _, v := range str {
		if unicode.IsUpper(v) {
			return true
		}
		return false
	}
	return false
}

func ToUFirst(str string) string {
	for _, v := range str {
		return string(unicode.ToUpper(v)) + str[1:]
	}
	return ""
}

func LFirst(str string) bool {
	for _, v := range str {
		if unicode.IsLower(v) {
			return true
		}
		return false
	}
	return false
}

func ToLFirst(str string) string {
	for _, v := range str {
		return string(unicode.ToLower(v)) + str[1:]
	}
	return ""
}
