package deepcopy

import (
	"reflect"
	"testing"
)

func TestDeepCopy_Int(t *testing.T) {
	result, err := DeepCopy(1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}

func TestDeepCopy_IntPtr(t *testing.T) {
	i := 1
	result, err := DeepCopy(&i)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result, reflect.TypeOf(result))
}

func TestDeepCopy_IntInterface(t *testing.T) {
	var i interface{} = 1

	result, err := DeepCopy(&i)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result, reflect.TypeOf(result))
}

func TestDeepCopy_Slice(t *testing.T) {
	result, err := DeepCopy([]interface{}{1, "test"})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}

func TestDeepCopy_Structure(t *testing.T) {
	type TestStructure struct {
		Int    int
		String string `copy:"-"`
	}

	result, err := DeepCopy([]interface{}{TestStructure{Int: 1, String: "2"}})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}
