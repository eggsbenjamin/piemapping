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
			conn := NewConnection(logr, nil)
			defer conn.Close()
			db := NewDBWrapper(conn)
			jRepo := NewJourneyRepository(db, logr)
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

	var _ = Describe("GetByDepLocWithinRange", func() {
		It("should return the correct journeys when found", func() {
			By("setup")
			conn := NewConnection(logr, nil)
			defer conn.Close()
			db := NewDBWrapper(conn)
			jRepo := NewJourneyRepository(db, logr)
			depLoc := "London"
			st := "2016-03-11 19:00:00"
			end := "2016-03-11 22:00:00"
			expected := []*models.Journey{}
			fixture, err := ioutil.ReadFile("../fixtures/london_2016-03-11 19:00:00_2016-03-11 22:00:00_overlapping_journeys.json")
			err = json.Unmarshal(fixture, &expected)
			Expect(err).NotTo(HaveOccurred())

			By("making call")
			actual, err := jRepo.GetByDepLocWithinRange(depLoc, st, end)

			By("assertions")
			Expect(err).NotTo(HaveOccurred())
			Expect(actual).To(Equal(expected))
		})
	})
})
