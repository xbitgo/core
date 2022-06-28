package tool_mock

import (
	"reflect"
	"strings"
	"time"

	"github.com/xbitgo/core/tools/tool_convert"

	"github.com/brianvoe/gofakeit/v6"
)

type DataMocker struct {
	faker *gofakeit.Faker
}

func NewDataMocker() *DataMocker {
	return &DataMocker{
		faker: gofakeit.New(time.Now().Unix()),
	}
}

func (mocker *DataMocker) Struct(input interface{}) {
	mocker.faker.Struct(input)

	mocker.FixMockData(reflect.Indirect(reflect.ValueOf(input)))
}

func (mocker *DataMocker) FixMockData(v reflect.Value) {
	switch v.Kind() {
	case reflect.Struct:
		mocker.FixMockDataStruct(v)
	case reflect.Slice:
		mocker.FixMockDataSlice(v)
	}
}

func (mocker *DataMocker) FixMockDataStruct(v reflect.Value) {
	for i := 0; i != v.NumField(); i++ {
		field := v.Field(i)
		fieldType := v.Type().Field(i)
		if !field.CanSet() {
			continue
		}

		mocker.FixMockDataField(v, field, fieldType)
	}
}

func (mocker *DataMocker) FixMockDataSlice(v reflect.Value) {
	for i := 0; i != v.Len(); i++ {
		mocker.FixMockData(v.Index(i))
	}
}

func (mocker *DataMocker) FixMockDataField(v reflect.Value, field reflect.Value, structField reflect.StructField) {
	if len(structField.Tag.Get("fake")) != 0 {
		return
	}

	switch field.Kind() {
	case reflect.String:
		if strings.HasSuffix(structField.Name, "Name") {
			field.SetString(mocker.faker.Name())
		}

		if strings.HasSuffix(structField.Name, "Id") {
			field.SetString(tool_convert.ToString(mocker.faker.Number(1, 100)))
		}

		if strings.HasSuffix(structField.Name, "Date") {
			field.SetString(time.Now().Format("2006-01-02"))
		}

		if strings.HasSuffix(structField.Name, "Time") {
			field.SetString(time.Now().Format("2006-01-02 15:04:05"))
		}

		if strings.HasSuffix(structField.Name, "Url") {
			field.SetString(mocker.faker.URL())
		}

		if strings.HasSuffix(structField.Name, "Email") {
			field.SetString(mocker.faker.Email())
		}

		if strings.HasSuffix(structField.Name, "ErrorMessage") {
			field.SetString("success")
		}
	case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
		if strings.HasSuffix(structField.Name, "Id") {
			field.SetInt(int64(mocker.faker.Number(1, 100)))
		}

		if strings.HasSuffix(structField.Name, "Status") {
			field.SetInt(int64(mocker.faker.Number(1, 5)))
		}

		if strings.HasSuffix(structField.Name, "Type") {
			field.SetInt(int64(mocker.faker.Number(1, 20)))
		}

		if strings.HasSuffix(structField.Name, "Code") {
			field.SetInt(0)
		}
	case reflect.Ptr:
		mocker.FixMockData(field.Elem())
	case reflect.Struct:
		if _, ok := field.Interface().(time.Time); ok {
			field.Set(reflect.ValueOf(time.Now()))
		}
		mocker.FixMockData(field)
	case reflect.Slice:
		mocker.FixMockData(field)
	}
}
