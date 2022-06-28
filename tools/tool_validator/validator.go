package tool_validator

import (
	"fmt"
	"strings"
	"sync"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	v "github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	validate *v.Validate
	trans    ut.Translator
	once     sync.Once
)

func init() {
	MustRegisterStringValidation("date", CheckDate, false)
	MustRegisterStringValidation("datetime", CheckDatetime, false)
}

func initValidator() {
	//实例化需要转换的语言
	zhT := zh.New()
	enT := en.New()
	uni := ut.New(enT, zhT, enT)
	var ok = false
	trans, ok = uni.GetTranslator("zh")
	if !ok {
		panic("validator get zh fail")
	}
	validate = v.New()

	if err := zhTranslations.RegisterDefaultTranslations(validate, trans); err != nil {
		panic(err)
	}

	fmt.Println("validator init")
}

func GetValidator() *v.Validate {
	once.Do(initValidator)
	return validate
}

func GetTranslator() ut.Translator {
	once.Do(initValidator)
	return trans
}

func ValidateStruct(input interface{}) error {
	return TranslateValidateError(GetValidator().Struct(input))
}

func TranslateValidateError(err error) error {
	if err == nil {
		return err
	}
	errs, ok := err.(v.ValidationErrors)
	if ok {
		return &ValidateErrorWithTranslation{errs.Translate(GetTranslator())}
	}

	return err
}

func MustRegisterStringValidation(tag string, f func(string) bool, callValidationEvenIfNull ...bool) {
	if err := GetValidator().RegisterValidation(tag, func(fl v.FieldLevel) bool {
		return f(fl.Field().String())
	}, callValidationEvenIfNull...); err != nil {
		panic(err)
	}
}

func MustRegisterIntValidation(tag string, f func(int64) bool, callValidationEvenIfNull ...bool) {
	if err := GetValidator().RegisterValidation(tag, func(fl v.FieldLevel) bool {
		return f(fl.Field().Int())
	}, callValidationEvenIfNull...); err != nil {
		panic(err)
	}
}

type ValidateErrorWithTranslation struct {
	v.ValidationErrorsTranslations
}

func (e ValidateErrorWithTranslation) Error() string {
	result := make([]string, 0, len(e.ValidationErrorsTranslations))
	for _, v := range e.ValidationErrorsTranslations {
		result = append(result, fmt.Sprintf("%s", v))
	}
	return strings.Join(result, ",")
}

func AsValidateError(err error) *ValidateErrorWithTranslation {
	if err == nil {
		return nil
	}

	errs, ok := err.(*ValidateErrorWithTranslation)
	if ok {
		return errs
	}

	return nil
}
