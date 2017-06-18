package http_handlers_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	. "github.com/eggsbenjamin/piemapping/http_handlers"
	"github.com/eggsbenjamin/piemapping/models"
	"github.com/eggsbenjamin/piemapping/repository/mocks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Journey", func() {
	var _ = Describe("GetByDriverAvailability", func() {
		It("returns a 500 if an error occurs when getting the journeys", func() {
			By("setup")
			repo := &mocks.JourneyRepositor{}
			repo.On("GetByDriverAvailability", "").Return(nil, errors.New("error"))
			hdlr := NewDriverAvailabilityHandler(repo, logr)
			expectedStatusCode := http.StatusInternalServerError
			expectedBody := `{ "message": "internal server error" }`
			req, err := http.NewRequest("GET", "", nil)
			res := httptest.NewRecorder()
			Expect(err).NotTo(HaveOccurred())

			By("making call")
			hdlr.Handle(res, req)
			actual := res.Result()
			actualBody, _ := ioutil.ReadAll(actual.Body)

			By("assert")
			Expect(actual.StatusCode).To(Equal(expectedStatusCode))
			Expect(actualBody).To(MatchJSON(expectedBody))
		})

		It("returns a 200 and the journeys if no error occurs when getting the journeys", func() {
			By("setup")
			jrnys := []*models.Journey{}
			fixture, err := ioutil.ReadFile("../fixtures/driver_2_journeys.json")
			err = json.Unmarshal(fixture, &jrnys)
			repo := &mocks.JourneyRepositor{}
			repo.On("GetByDriverAvailability", "").Return(jrnys, nil)
			hdlr := NewDriverAvailabilityHandler(repo, logr)
			expectedStatusCode := http.StatusOK
			expectedBody := string(fixture)
			req, err := http.NewRequest("GET", "", nil)
			res := httptest.NewRecorder()
			Expect(err).NotTo(HaveOccurred())

			By("making call")
			hdlr.Handle(res, req)
			actualStatusCode := res.Code
			actualBody := res.Body.String()
			Expect(err).NotTo(HaveOccurred())

			By("assert")
			Expect(actualStatusCode).To(Equal(expectedStatusCode))
			Expect(actualBody).To(MatchJSON(expectedBody))
		})
	})
})
