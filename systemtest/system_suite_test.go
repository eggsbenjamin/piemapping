package system_test

import (
	"fmt"
	"net"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/spf13/viper"

	"testing"
)

var (
	baseUrl    string
	binaryPath string
	session    *gexec.Session
)

var _ = BeforeSuite(func() {
	var err error
	viper.AddConfigPath("../config")
	viper.SetConfigName("app")
	err = viper.ReadInConfig()
	Expect(err).NotTo(HaveOccurred())

	sourcePath := "github.com/eggsbenjamin/piemapping/cmd"
	binaryPath, err = gexec.Build(sourcePath)
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

func verifyServerIsListening() error {
	fmt.Println(baseUrl)
	_, err := net.Dial("tcp", baseUrl)
	return err
}

var _ = BeforeEach(func() {
	var err error

	session, err = gexec.Start(exec.Command(binaryPath), GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())

	host := viper.GetString("development.host")
	port := viper.GetString("development.port")
	baseUrl = fmt.Sprintf("%s:%s", host, port)
	Eventually(verifyServerIsListening).Should(Succeed())
})

var _ = AfterEach(func() {
	session.Interrupt()
	Eventually(session).Should(gexec.Exit())
})

func TestSystemtest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "System Test Suite")
}
