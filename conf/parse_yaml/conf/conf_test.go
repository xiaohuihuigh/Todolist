package conf_test
//
//import (
//	"io"
//	"os"
//	"path/filepath"
//	"testing"
//
//	todo/gopher/qms/pkg/conf"
//	"github.com/go-chassis/go-archaius"
//	"github.com/stretchr/testify/assert"
//)
//
//var filename2 string
//
//func TestInit(t *testing.T) {
//	f1Bytes := []byte(`
//age: 14
//name: peter
//`)
//	f2Bytes := []byte(`
//addr: somewhere
//number: 1
//exist: true
//`)
//	d, _ := os.Getwd()
//	filename1 := filepath.Join(d, "f1.yaml")
//	filename2 = filepath.Join(d, "f2.yaml")
//	f1, err := os.Create(filename1)
//	assert.NoError(t, err)
//	defer f1.Close()
//	defer os.Remove(filename1)
//	f2, err := os.Create(filename2)
//	assert.NoError(t, err)
//	defer f2.Close()
//	defer os.Remove(filename2)
//	_, err = io.WriteString(f1, string(f1Bytes))
//	assert.NoError(t, err)
//	_, err = io.WriteString(f2, string(f2Bytes))
//	assert.NoError(t, err)
//	os.Setenv("age", "15")
//
//	err = archaius.Init(
//		archaius.WithRequiredFiles([]string{filename1}),
//		archaius.WithOptionalFiles([]string{filename2}),
//		archaius.WithENVSource(),
//		archaius.WithMemorySource())
//	assert.NoError(t, err)
//	assert.Equal(t, "15", archaius.Get("age"))
//	t.Run("add mem config", func(t *testing.T) {
//		archaius.Set("age", "16")
//		assert.Equal(t, "16", archaius.Get("age"))
//	})
//	t.Run("delete mem config", func(t *testing.T) {
//		archaius.Delete("age")
//		assert.Equal(t, "15", archaius.Get("age"))
//	})
//
//}
//
//func TestConfig_Get(t *testing.T) {
//	TestInit(t)
//
//	s := conf.Get("number")
//	assert.Equal(t, 1, s)
//
//	e := conf.GetBool("exist", false)
//	assert.Equal(t, true, e)
//
//	n := conf.Get("name")
//	assert.Equal(t, "peter", n)
//
//	n3 := conf.GetString("name", "")
//	assert.Equal(t, "peter", n3)
//
//	b := conf.Exist("name")
//	assert.True(t, b)
//
//	b = conf.Exist("none")
//	assert.False(t, b)
//
//	m := conf.GetConfigs()
//	//t.Log(m)
//	assert.Equal(t, 1, m["number"])
//
//	s = conf.GetInt("number", 0)
//	assert.Equal(t, 1, s)
//}
//
//func TestUnmarshalConfig(t *testing.T) {
//	TestInit(t)
//	b := []byte(`
//key: peter
//info:
//  address: a
//  number: 8
//metadata:
//  a: b
//`)
//	d, _ := os.Getwd()
//	os.Setenv("QMS_CONF_DIR", d)
//	filename1 := filepath.Join(d, "app.yaml")
//	f1, err := os.Create(filename1)
//	assert.NoError(t, err)
//	defer f1.Close()
//	defer os.Remove(filename1)
//	_, err = io.WriteString(f1, string(b))
//	assert.NoError(t, err)
//
//	type Info struct {
//		Addr   string `yaml:"address"`
//		Number int    `yaml:"number"`
//	}
//	type Person struct {
//		Name string            `yaml:"key"`
//		MD   map[string]string `yaml:"metadata"`
//		Info *Info             `yaml:"info"`
//	}
//	err = archaius.AddFile(filename1)
//	assert.NoError(t, err)
//	p := &Person{}
//	err = conf.Unmarshal(p)
//	assert.NoError(t, err)
//	assert.Equal(t, "peter", p.Name)
//	assert.Equal(t, "b", p.MD["a"])
//	assert.Equal(t, "a", p.Info.Addr)
//	assert.Equal(t, 8, p.Info.Number)
//}
