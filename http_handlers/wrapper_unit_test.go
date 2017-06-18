package http_handlers_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	. "github.com/eggsbenjamin/piemapping/http_handlers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type panicker struct{}

func (p *panicker) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	panic("argh!")
}

var _ = Describe("Wrapper", func() {
	It("recovers and returns a 500 response on panic", func() {
		By("setup")
		hw := NewHandlerWrapper(logr)
		p := &panicker{}
		wr := hw.Init(p)
		expectedStatusCode := http.StatusInternalServerError
		expectedBody := `{ "message": "internal server error" }`
		req, _ := http.NewRequest("GET", "", nil)
		res := httptest.NewRecorder()

		By("making call")
		wr.ServeHTTP(res, req)
		actual := res.Result()
		actualBody, _ := ioutil.ReadAll(actual.Body)

		By("assert")
		Expect(actual.StatusCode).To(Equal(expectedStatusCode))
		Expect(actualBody).To(MatchJSON(expectedBody))
	})
})
