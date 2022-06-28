package tool_validator

import (
	"regexp"
	"testing"
)

type TestData struct {
	Id            int64  `validate:"required"`
	Name          string `validate:"min=2"`
	Password      string `validate:"required,password"`
	PasswordAgain string `validate:"required,eqfield=Password"`
}

func TestValidate(t *testing.T) {
	MustRegisterStringValidation("password", func(s string) bool {
		reg := regexp.MustCompile(`^[a-zA-Z0-9._\-@!#$%^&*]+$`)
		return reg.MatchString(s)
	})

	testData := &TestData{
		Id:            1,
		Name:          "1",
		Password:      "1234",
		PasswordAgain: "12345",
	}

	err := ValidateStruct(testData)

	t.Log(err)
}
