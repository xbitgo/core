package config

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"github.com/xbitgo/core/tools/tool_file"
	"io/ioutil"
	"log"
	"reflect"
)

type Json struct {
	ConfigFile string
}

func (j Json) Apply(rs interface{}) error {
	if j.ConfigFile == "" {
		return errors.New("[cfg] ConfigFile is empty")
	}
	if rs == nil {
		return errors.New("[cfg] param rs is nil")
	}
	typ := reflect.TypeOf(rs)
	if typ.Kind() != reflect.Ptr {
		return errors.New("[cfg] cannot apply to non-pointer struct")
	}
	if tool_file.Exists(j.ConfigFile) {
		buf, _ := ioutil.ReadFile(j.ConfigFile)
		err := jsoniter.Unmarshal(buf, rs)
		if err != nil {
			log.Printf("%s 配置文件解析异常！", j.ConfigFile)
			return err
		}
	}
	return nil
}
