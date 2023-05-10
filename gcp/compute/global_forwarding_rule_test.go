package compute_test

import (
	"errors"

	"github.com/genevieve/leftovers/gcp/compute"
	"github.com/genevieve/leftovers/gcp/compute/fakes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GlobalForwardingRule", func() {
	var (
		client *fakes.GlobalForwardingRulesClient
		name   string

		globalForwardingRule compute.GlobalForwardingRule
	)

	BeforeEach(func() {
		client = &fakes.GlobalForwardingRulesClient{}
		name = "banana"

		globalForwardingRule = compute.NewGlobalForwardingRule(client, name)
	})

	Describe("Delete", func() {
		It("deletes the global forwarding rule", func() {
			err := globalForwardingRule.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DeleteGlobalForwardingRuleCall.CallCount).To(Equal(1))
			Expect(client.DeleteGlobalForwardingRuleCall.Receives.Rule).To(Equal(name))
		})

		Context("when the client fails to delete", func() {
			BeforeEach(func() {
				client.DeleteGlobalForwardingRuleCall.Returns.Error = errors.New("the-error")
			})

			It("returns the error", func() {
				err := globalForwardingRule.Delete()
				Expect(err).To(MatchError("Delete: the-error"))
			})
		})
	})

	Describe("Name", func() {
		It("returns the name", func() {
			Expect(globalForwardingRule.Name()).To(Equal(name))
		})
	})

	Describe("Type", func() {
		It("returns the type", func() {
			Expect(globalForwardingRule.Type()).To(Equal("Global Forwarding Rule"))
		})
	})
})
