package iam_test

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/genevieve/leftovers/aws/iam"
	"github.com/genevieve/leftovers/aws/iam/fakes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {
	var (
		user       iam.User
		client     *fakes.UsersClient
		policies   *fakes.UserPolicies
		accessKeys *fakes.AccessKeys
		name       *string
	)

	BeforeEach(func() {
		client = &fakes.UsersClient{}
		policies = &fakes.UserPolicies{}
		accessKeys = &fakes.AccessKeys{}
		name = aws.String("the-name")

		user = iam.NewUser(client, policies, accessKeys, name)
	})

	Describe("Delete", func() {
		It("deletes the user", func() {
			err := user.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(policies.DeleteCall.CallCount).To(Equal(1))
			Expect(policies.DeleteCall.Receives.UserName).To(Equal(*name))

			Expect(accessKeys.DeleteCall.CallCount).To(Equal(1))
			Expect(accessKeys.DeleteCall.Receives.UserName).To(Equal(*name))

			Expect(client.DeleteUserCall.CallCount).To(Equal(1))
			Expect(client.DeleteUserCall.Receives.DeleteUserInput.UserName).To(Equal(name))
		})

		Context("when deleting the user's access keys fails", func() {
			BeforeEach(func() {
				accessKeys.DeleteCall.Returns.Error = errors.New("banana")
			})

			It("returns the error", func() {
				err := user.Delete()
				Expect(err).To(MatchError("Delete access keys: banana"))
			})
		})

		Context("when deleting the user's policies fails", func() {
			BeforeEach(func() {
				policies.DeleteCall.Returns.Error = errors.New("banana")
			})

			It("returns the error", func() {
				err := user.Delete()
				Expect(err).To(MatchError("Delete policies: banana"))
			})
		})

		Context("when the client fails", func() {
			BeforeEach(func() {
				client.DeleteUserCall.Returns.Error = errors.New("banana")
			})

			It("returns the error", func() {
				err := user.Delete()
				Expect(err).To(MatchError("Delete: banana"))
			})
		})
	})

	Describe("Name", func() {
		It("returns the identifier", func() {
			Expect(user.Name()).To(Equal("the-name"))
		})
	})

	Describe("Type", func() {
		It("returns the type", func() {
			Expect(user.Type()).To(Equal("IAM User"))
		})
	})
})
