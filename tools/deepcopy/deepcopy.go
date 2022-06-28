package deepcopy

import (
	"fmt"
	"reflect"
	"time"
)

func DeepCopy(input interface{}) (interface{}, error) {
	if input == nil {
		return nil, nil
	}

	inputV := reflect.ValueOf(input)
	outputV := reflect.New(inputV.Type()).Elem()

	if err := deepCopyValue(inputV, outputV); err != nil {
		return nil, err
	}

	return outputV.Interface(), nil
}

func deepCopyValue(input reflect.Value, output reflect.Value) error {

	switch input.Kind() {
	case reflect.Ptr:
		originalValue := input.Elem()
		if !originalValue.IsValid() {
			return nil
		}

		output.Set(reflect.New(originalValue.Type()))
		return deepCopyValue(originalValue, output.Elem())

	case reflect.Interface:
		if input.IsNil() {
			return nil
		}

		originalValue := input.Elem()
		copyValue := reflect.New(originalValue.Type()).Elem()
		if err := deepCopyValue(originalValue, copyValue); err != nil {
			return err
		}
		output.Set(copyValue)

	case reflect.Slice:
		if input.IsNil() {
			return nil
		}

		output.Set(reflect.MakeSlice(input.Type(), input.Len(), input.Len()))
		for idx := 0; idx != input.Len(); idx++ {
			if err := deepCopyValue(input.Index(idx), output.Index(idx)); err != nil {
				return err
			}
		}
	case reflect.Map:
		if input.IsNil() {
			return nil
		}

		output.Set(reflect.MakeMap(input.Type()))

		iter := input.MapRange()
		for iter.Next() {
			originalValue := iter.Value()
			copyValue := reflect.New(originalValue.Type()).Elem()
			if err := deepCopyValue(originalValue, copyValue); err != nil {
				return err
			}

			copyKey, err := DeepCopy(iter.Key().Interface())
			if err != nil {
				return err
			}

			output.SetMapIndex(reflect.ValueOf(copyKey), copyValue)
		}

	case reflect.Struct:
		t, ok := input.Interface().(time.Time)
		if ok {
			output.Set(reflect.ValueOf(t))
			return nil
		}

		// Go through each field of the struct and copy it.
		for i := 0; i < input.NumField(); i++ {
			// The Type's StructField for a given field is checked to see if StructField.PkgPath
			// is set to determine if the field is exported or not because CanSet() returns false
			// for settable fields.  I'm not sure why.  -mohae
			if input.Type().Field(i).PkgPath != "" {
				continue
			}
			if input.Type().Field(i).Tag.Get("copy") == "-" {
				continue
			}

			if err := deepCopyValue(input.Field(i), output.Field(i)); err != nil {
				return err
			}
		}

	case reflect.Invalid:
		return fmt.Errorf("value invalid in deepcopy")
	case reflect.Chan:
		return fmt.Errorf("channel member is not copied in deepcopy, need to exclude it explicitly")
	case reflect.Func:
		return fmt.Errorf("function member is not copied in deepcopy, need to exclude it explicitly")
	default:
		output.Set(input)
	}

	return nil
}
