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

	It("responds to GET '/journeys/start/{depLoc}/range/{start}/to/{end}' with correct payload", func() {
		By("setup")
		depLoc := "London"
		st := "2016-03-11 19:00:00"
		end := "2016-03-11 22:00:00"
		url := fmt.Sprintf("http://%s/journeys/start/%s/range/%s/to/%s", baseUrl, depLoc, st, end)
		fixture, err := ioutil.ReadFile("../fixtures/london_2016-03-11 19:00:00_2016-03-11 22:00:00_overlapping_journeys.json")
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
