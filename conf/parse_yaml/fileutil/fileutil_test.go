package fileutil_test

//import (
//	"os"
//	"path/filepath"
//	"testing"
//
//	"git.qutoutiao.net/gopher/qms/internal/pkg/utils/fileutil"
//
//	"github.com/stretchr/testify/assert"
//)
//
//func TestGetWorkDirHmNotSet(t *testing.T) {
//	os.Setenv("QMS_HOME", "test")
//	assert.Equal(t, "conf", filepath.Base(fileutil.GetConfDir()))
//
//}
//func TestGetWorkDir(t *testing.T) {
//	os.Setenv("QMS_HOME", "test")
//	assert.Equal(t, "conf", filepath.Base(fileutil.GetConfDir()))
//}
//
//func TestHystricDefinaiton(t *testing.T) {
//	os.Setenv("QMS_HOME", "test")
//	def := fileutil.CircuitBreakerConfigPath()
//	assert.Equal(t, filepath.Join("test", "conf", fileutil.Hystric), def)
//}
//func TestMicroserviceDefinition(t *testing.T) {
//	def := fileutil.MicroserviceDefinition("micro")
//	assert.Equal(t, filepath.Join("test", "conf", "micro", fileutil.Definition), def)
//}
//func TestGlobalDefinition(t *testing.T) {
//	def := fileutil.GlobalConfigPath()
//	assert.Equal(t, filepath.Join("test", "conf", fileutil.Global), def)
//}
//func TestPassLagerDefinition(t *testing.T) {
//	def := fileutil.LogConfigPath()
//	assert.Equal(t, filepath.Join("test", "conf", fileutil.PaasLager), def)
//}
//func TestSchemaDir(t *testing.T) {
//	def := fileutil.SchemaDir("micro")
//	assert.Equal(t, filepath.Join("test", "conf", "micro", fileutil.SchemaDirectory), def)
//}
//func TestGetDefinition(t *testing.T) {
//	def := fileutil.GetDefinition()
//	assert.Equal(t, filepath.Join("test", "conf", fileutil.Definition), def)
//}
//func TestGetWorkDirConfSet(t *testing.T) {
//	os.Setenv("QMS_CONF_DIR", "conf")
//	assert.Equal(t, "conf", filepath.Base(fileutil.GetConfDir()))
//}
//
//func TestGetWorkDirValid(t *testing.T) {
//	_, err := fileutil.GetWorkDir()
//	assert.Nil(t, err)
//}
//
//func TestQmsHomeDir(t *testing.T) {
//	assert.NotEmpty(t, fileutil.QmsHomeDir())
//}
//
//func TestGetLoadBalancing(t *testing.T) {
//	assert.NotEmpty(t, fileutil.LoadBalancingConfigPath())
//}
//
//func TestGetRateLimiting(t *testing.T) {
//	assert.NotEmpty(t, fileutil.RateLimitingFile())
//}
//
//func TestGetTLS(t *testing.T) {
//	assert.NotEmpty(t, fileutil.TLSConfigPath())
//}
//
//func TestGetMonitoring(t *testing.T) {
//	assert.NotEmpty(t, fileutil.MonitoringConfigPath())
//}
//
//func TestGetRouter(t *testing.T) {
//	assert.NotEmpty(t, fileutil.RouterConfigPath())
//}
//
//func TestGetMicroserviceDesc(t *testing.T) {
//	assert.NotEmpty(t, fileutil.MicroServiceConfigPath())
//}
//
//func TestGetAuth(t *testing.T) {
//	assert.NotEmpty(t, fileutil.AuthConfigPath())
//}
//
//func TestGetTracing(t *testing.T) {
//	assert.NotEmpty(t, fileutil.TracingPath())
//}
