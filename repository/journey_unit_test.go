//	+build unit

package repository_test

import (
	"errors"

	"github.com/eggsbenjamin/piemapping/models"
	. "github.com/eggsbenjamin/piemapping/repository"
	"github.com/eggsbenjamin/piemapping/repository/mocks"
	testify "github.com/stretchr/testify/mock"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Journey Repository", func() {
	var _ = Describe("GetByDriverAvailability", func() {
		It("returns the correct error if an error occurs when querying", func() {
			By("setup")
			qryr := &mocks.Queryer{}
			mockErr := errors.New("error")
			expected := "unable to get journeys"
			qryr.On("Query",
				testify.AnythingOfType("string"),
				testify.AnythingOfType("string"),
			).Return(nil, mockErr)
			jRepo := NewJourneyRepository(qryr, logr)

			By("making call")
			actual, err := jRepo.GetByDriverAvailability("")

			By("assert")
			Expect(actual).To(BeNil())
			Expect(err.Error()).To(Equal(expected))
		})

		It("returns the correct error if an error occurs when scanning results", func() {
			By("setup")
			qryr := &mocks.Queryer{}
			rIt := &mocks.RowIterator{}
			mockErr := errors.New("error")
			expected := "unable to get journeys"
			rIt.On("Close").Return(nil)
			rIt.On("Next").Return(true).Once()
			rIt.On("Scan",
				testify.Anything,
				testify.Anything,
				testify.Anything,
				testify.Anything,
				testify.Anything,
			).Return(mockErr)
			qryr.On("Query",
				testify.AnythingOfType("string"),
				testify.AnythingOfType("string"),
			).Return(rIt, nil)
			jRepo := NewJourneyRepository(qryr, logr)

			By("making call")
			actual, err := jRepo.GetByDriverAvailability("")

			By("assert")
			Expect(actual).To(BeNil())
			Expect(err.Error()).To(Equal(expected))
		})

		It("returns the correct error if an error occurs when iterating through results", func() {
			By("setup")
			qryr := &mocks.Queryer{}
			rIt := &mocks.RowIterator{}
			mockErr := errors.New("error")
			expected := "unable to get journeys"
			rIt.On("Close").Return(nil)
			rIt.On("Next").Return(false).Once()
			rIt.On("Err").Return(mockErr)
			qryr.On("Query",
				testify.AnythingOfType("string"),
				testify.AnythingOfType("string"),
			).Return(rIt, nil)
			jRepo := NewJourneyRepository(qryr, logr)

			By("making call")
			actual, err := jRepo.GetByDriverAvailability("")

			By("assert")
			Expect(actual).To(BeNil())
			Expect(err.Error()).To(Equal(expected))
		})

		It("returns an empty array if no results are returned", func() {
			By("setup")
			qryr := &mocks.Queryer{}
			rIt := &mocks.RowIterator{}
			expected := []*models.Journey{}
			rIt.On("Close").Return(nil)
			rIt.On("Next").Return(false).Once()
			rIt.On("Err").Return(nil)
			qryr.On("Query",
				testify.AnythingOfType("string"),
				testify.AnythingOfType("string"),
			).Return(rIt, nil)
			jRepo := NewJourneyRepository(qryr, logr)

			By("making call")
			actual, err := jRepo.GetByDriverAvailability("")

			By("assert")
			Expect(err).NotTo(HaveOccurred())
			Expect(actual).To(Equal(expected))
		})
	})
})
