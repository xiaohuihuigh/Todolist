package parse_yaml

import (
	"bytes"
	"todoList/conf/parse_yaml/fileutil"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Yaml struct {
}

func NewYaml()*Yaml{
	return &Yaml{}
}


// Unmarshal unmarshals the config into a Struct. Make sure that the tags
// on the fields of the structure are properly set.
func Unmarshal(obj interface{}) error {
	content, err := ioutil.ReadFile(fileutil.AppConfigPath())
	if err != nil {
		return err
	}
	decoder := yaml.NewDecoder(bytes.NewReader(content))
	return decoder.Decode(obj)
}
