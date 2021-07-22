package httpclient

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gojuno/minimock/v3"
	fuzz "github.com/google/gofuzz"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Doer Decorators", func() {
	var (
		doerMock *DoerMock
		doer     Doer
		mc       *minimock.Controller
		fuzzer   = fuzz.New()

		req         *http.Request
		resp        *http.Response
		serviceName string
	)

	BeforeEach(func() {
		mc = minimock.NewController(mc)
		doerMock = NewDoerMock(mc)

		req = &http.Request{}
		resp = &http.Response{}
		fuzzer.Fuzz(&serviceName)

		doer = DecorateDoer(doerMock, ErrorDoerDecorator(serviceName))
	})

	Context("Error decorator", func() {
		It("should call doer", func() {
			doerMock.DoMock.
				Expect(req).
				Return(resp, nil)

			got, err := doer.Do(req)
			Expect(err).NotTo(HaveOccurred())
			Expect(got).To(BeEquivalentTo(resp))
		})

		When("doer returns error", func() {
			doerErr := errors.New("err")

			BeforeEach(func() {
				doerMock.DoMock.
					Return(nil, doerErr)
			})

			It("should wrap err", func() {
				someNil, gotErr := doer.Do(req)

				Expect(someNil).To(BeNil())
				Expect(gotErr).
					To(MatchError(fmt.Errorf("httpclient_doer_error [%s]: %w", serviceName, doerErr)))
			})
		})

		When("doer returns no error", func() {
			BeforeEach(func() {
				doerMock.DoMock.
					Return(resp, nil)
			})

			It("should modify nothing", func() {
				got, err := doer.Do(req)
				Expect(err).NotTo(HaveOccurred())
				Expect(got).To(BeEquivalentTo(resp))
			})
		})
	})
})
