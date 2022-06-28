package tool_mock

import (
	"encoding/json"
	"testing"
	"time"
)

type Foo struct {
	Id            int
	Name          string
	ProjectIdList []int32
	Status        int32
	Bar           *Bar
	CreateTime    time.Time
}

type Bar struct {
	Code         int32
	ErrorMessage string
	BarId        int
	BarData      string
	CreateDate   string
	UpdateTime   string
}

func (foo *Foo) String() string {
	data, _ := json.Marshal(foo)
	return string(data)
}

func TestDataMocker_Struct(t *testing.T) {
	mocker := NewDataMocker()
	test := &Foo{}
	mocker.Struct(test)

	t.Log(test)
}

func TestDataMocker_StructRecursive(t *testing.T) {
	type Test struct {
		Value int   `json:"value"`
		Test  *Test `json:"test,omitempty" fake:"skip"`
	}

	mocker := NewDataMocker()
	test := &Test{}
	mocker.Struct(test)

	t.Log(test)
}
