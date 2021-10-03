package main_test

import (
	// . "newApp/main_test"

	"fmt"
	"io/ioutil"
	"net/http"
	. "newApp/controller"
	. "newApp/demo"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewApp", func() {
	Describe("testing check and Demos", func() {
		Context("check", func() {
			It("run test", func() {
				Expect(1).Should(Equal(1))
				Expect(Add(5, 6)).Should(Equal(11))
			})
		})

		var (
			p, q, m, n, sum1, sum2 int
		)
		BeforeEach(func() {
			p, q, sum1 = 5, 6, 11
			// Putting wrong value of sum2 intentionally
			m, n, sum2 = 8, 7, 16
		})
		Context("Addition of two digits", func() {
			It("should return sum of the two digits", func() {
				addition_of_two_digits := Add(p, q)
				Expect(addition_of_two_digits).Should(Equal(sum1))
			})
			It("should not return the sum provided", func() {
				addition_of_two_digits := Add(m, n)
				Expect(addition_of_two_digits).ShouldNot(Equal(sum2))
			})
		})

	})
	Describe("EndPoint Testing", func() {

		Context("Addition of two digits", func() {
			It("should return sum of the two digits", func() {
				// req, err := http.NewRequest("GET", "/api/books/", nil)
				req, err := http.Get("https://api.github.com/users/pradeepneosoft")
				if err != nil {
					print(err)
				}

				// rr := httptest.NewRecorder()
				// handler := http.HandlerFunc()
				// handler.ServeHTTP(rr, req)
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					print(err)
				}
				fmt.Print(string(body))
				Expect(1).Should(Equal(1))

			})

		})

	})

})
