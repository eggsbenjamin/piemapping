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

//	build binary
var _ = BeforeSuite(func() {
	var err error
	viper.AutomaticEnv()
	viper.SetEnvPrefix("piemapping")
	sourcePath := "github.com/eggsbenjamin/piemapping/cmd"
	binaryPath, err = gexec.Build(sourcePath)
	Expect(err).NotTo(HaveOccurred())
})

//	clean up binary
var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

func verifyServerIsListening() error {
	_, err := net.Dial("tcp", baseUrl)
	return err
}

//	spin up server
var _ = BeforeEach(func() {
	var err error
	session, err = gexec.Start(exec.Command(binaryPath, "run"), GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())

	host := viper.GetString("host")
	port := viper.GetString("port")
	baseUrl = fmt.Sprintf("%s:%s", host, port)
	Eventually(verifyServerIsListening).Should(Succeed())
})

//	tear down server
var _ = AfterEach(func() {
	session.Interrupt()
	Eventually(session).Should(gexec.Exit())
})

func TestSystemtest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "System Test Suite")
}
