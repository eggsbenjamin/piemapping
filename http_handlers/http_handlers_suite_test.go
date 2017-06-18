package http_handlers_test

import (
	"github.com/eggsbenjamin/piemapping/commons"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/viper"

	"testing"
)

var logr commons.LevelledLogWriter

var _ = BeforeSuite(func() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("piemapping")
	logs := viper.GetBool("logging")
	if logs == true {
		logr = commons.NewLogger("Repository Test", 1)
		return
	}
	logr = &commons.NoopLogger{}
})

func TestHttpHandlers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HttpHandlers Suite")
}
