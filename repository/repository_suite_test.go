package repository_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/viper"

	"testing"
)

var _ = BeforeSuite(func() {
	viper.AddConfigPath("../config")
	viper.SetConfigName("app")
	err := viper.ReadInConfig()
	Expect(err).NotTo(HaveOccurred())
})

func TestRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repository Test Suite")
}
