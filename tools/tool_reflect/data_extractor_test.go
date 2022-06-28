package tool_reflect

import "testing"

type StringMap map[string]string

type TestData struct {
	IntValue    int
	StringValue string
	IntSlice    []int
	StringMap   StringMap
	TestData    *TestData
}

func (data *TestData) GetWhatever() int {
	return 123
}

func (data StringMap) GetWhatever() int {
	return 123
}

func getTestData() interface{} {

	result := &TestData{
		IntValue:    1,
		StringValue: "stringTest",
		IntSlice:    []int{0, 1, 2, 3, 4, 5},
		StringMap: StringMap{
			"f1": "1",
			"f2": "2",
		},
	}

	result.TestData = result
	return result
}

func TestExtractData(t *testing.T) {
	if value := MustExtractData("IntValue", getTestData()); value == nil || value.(int) != 1 {
		t.Fatal(value)
	}

	if value := MustExtractData("StringValue", getTestData()); value == nil || value.(string) != "stringTest" {
		t.Fatal(value)
	}

	if value := MustExtractData("IntSlice.1", getTestData()); value == nil || value.(int) != 1 {
		t.Fatal(value)
	}

	if value := MustExtractData("StringMap.f2", getTestData()); value == nil || value.(string) != "2" {
		t.Fatal(value)
	}

	if value := MustExtractData("TestData.TestData.TestData.StringMap.f2", getTestData()); value == nil || value.(string) != "2" {
		t.Fatal(value)
	}

	if value := MustExtractData(".TestData.TestData.TestData.StringMap.f2", getTestData()); value == nil || value.(string) != "2" {
		t.Fatal(value)
	}

	if value := MustExtractData(".TestData.TestData.TestData.Whatever", getTestData()); value == nil || value.(int) != 123 {
		t.Fatal(value)
	}

	if value := MustExtractData(".TestData.TestData.TestData.StringMap.Whatever", getTestData()); value == nil || value.(int) != 123 {
		t.Fatal(value)
	}
}
