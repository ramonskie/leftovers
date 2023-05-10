package compute_test

import (
	"errors"

	"github.com/genevieve/leftovers/gcp/compute"
	"github.com/genevieve/leftovers/gcp/compute/fakes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GlobalHealthCheck", func() {
	var (
		client *fakes.GlobalHealthChecksClient
		name   string

		globalHealthCheck compute.GlobalHealthCheck
	)

	BeforeEach(func() {
		client = &fakes.GlobalHealthChecksClient{}
		name = "banana"

		globalHealthCheck = compute.NewGlobalHealthCheck(client, name)
	})

	Describe("Delete", func() {
		It("deletes the global health check", func() {
			err := globalHealthCheck.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DeleteGlobalHealthCheckCall.CallCount).To(Equal(1))
			Expect(client.DeleteGlobalHealthCheckCall.Receives.GlobalHealthCheck).To(Equal(name))
		})

		Context("when the client fails to delete", func() {
			BeforeEach(func() {
				client.DeleteGlobalHealthCheckCall.Returns.Error = errors.New("the-error")
			})

			It("returns the error", func() {
				err := globalHealthCheck.Delete()
				Expect(err).To(MatchError("Delete: the-error"))
			})
		})
	})

	Describe("Name", func() {
		It("returns the name", func() {
			Expect(globalHealthCheck.Name()).To(Equal(name))
		})
	})

	Describe("Type", func() {
		It("returns the type", func() {
			Expect(globalHealthCheck.Type()).To(Equal("Global Health Check"))
		})
	})
})
