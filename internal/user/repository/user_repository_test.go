package repository_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/thefuga/go-template/internal/user/repository"
)

var _ = Describe("UserRepository", func() {
	var userRepository repository.UserRepository

	Describe("FindByFirstName", func() {
		It("returns a use with firstname", func() {
			Expect(userRepository.FindByFirstName("jon").FirstName).
				To(Equal("jon"))
		})
	})
})
