package system_test

import (
	"fmt"
	"io/ioutil"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("System", func() {
	It("responds to GET '/journeys/driver/{driverId}/availability' with correct payload", func() {
		By("setup")
		driverId := "DRIVER_2_ID"
		url := fmt.Sprintf("http://%s/journeys/driver/%s/availability", baseUrl, driverId)
		fixture, err := ioutil.ReadFile("../fixtures/driver_2_journeys.json")
		expected := string(fixture)
		Expect(err).NotTo(HaveOccurred())

		By("making request")
		resp, err := http.Get(url)
		Expect(err).NotTo(HaveOccurred())

		By("assertions")
		defer resp.Body.Close()
		raw, err := ioutil.ReadAll(resp.Body)
		Expect(err).NotTo(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(http.StatusOK))
		Expect(raw).To(MatchJSON(expected))
	})
})
