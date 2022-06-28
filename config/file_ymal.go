package config

import (
	"github.com/pkg/errors"
	"github.com/xbitgo/core/tools/tool_file"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"reflect"
)

type Yaml struct {
	ConfigFile string
}

func (y Yaml) Apply(rs interface{}) error {
	if y.ConfigFile == "" {
		return errors.New("[cfg] ConfigFile is empty")
	}
	if rs == nil {
		return errors.New("[cfg] param rs is nil")
	}
	typ := reflect.TypeOf(rs)
	if typ.Kind() != reflect.Ptr {
		return errors.New("[cfg] cannot apply to non-pointer struct")
	}
	if tool_file.Exists(y.ConfigFile) {
		buf, _ := ioutil.ReadFile(y.ConfigFile)
		err := yaml.Unmarshal(buf, rs)
		if err != nil {
			log.Printf("%s 配置文件解析异常！", y.ConfigFile)
			return err
		}
	}
	return nil
}
