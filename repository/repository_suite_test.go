package repository_test

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
		lf := viper.GetString("log_format")
		df := commons.GetLogLevelId(lf)
		logr = commons.NewLogger("Repository Test", df)
		return
	}
	logr = &commons.NoopLogger{}
})

func TestRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repository Test Suite")
}
