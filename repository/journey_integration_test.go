//	+build integration

package repository_test

import (
	"encoding/json"
	"io/ioutil"

	"github.com/eggsbenjamin/piemapping/models"
	. "github.com/eggsbenjamin/piemapping/repository"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Journey Repository", func() {
	var _ = Describe("GetByDriverAvailability", func() {
		It("should return the correct journeys when found", func() {
			By("setup")
			conn := NewConnection()
			db := NewDBWrapper(conn)
			jRepo := NewJourneyRepository(db)
			driverId := "DRIVER_2_ID"
			expected := []*models.Journey{}
			fixture, err := ioutil.ReadFile("../fixtures/driver_2_journeys.json")
			err = json.Unmarshal(fixture, &expected)
			Expect(err).NotTo(HaveOccurred())

			By("making call")
			actual, err := jRepo.GetByDriverAvailability(driverId)

			By("assertions")
			Expect(err).NotTo(HaveOccurred())
			Expect(actual).To(Equal(expected))
		})
	})
})
