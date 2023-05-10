package compute_test

import (
	"errors"

	"github.com/genevieve/leftovers/gcp/compute"
	"github.com/genevieve/leftovers/gcp/compute/fakes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("HttpHealthCheck", func() {
	var (
		client *fakes.HttpHealthChecksClient
		name   string

		httpHealthCheck compute.HttpHealthCheck
	)

	BeforeEach(func() {
		client = &fakes.HttpHealthChecksClient{}
		name = "banana"

		httpHealthCheck = compute.NewHttpHealthCheck(client, name)
	})

	Describe("Delete", func() {
		It("deletes the http health check", func() {
			err := httpHealthCheck.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DeleteHttpHealthCheckCall.CallCount).To(Equal(1))
			Expect(client.DeleteHttpHealthCheckCall.Receives.HttpHealthCheck).To(Equal(name))
		})

		Context("when the client fails to delete", func() {
			BeforeEach(func() {
				client.DeleteHttpHealthCheckCall.Returns.Error = errors.New("the-error")
			})

			It("returns the error", func() {
				err := httpHealthCheck.Delete()
				Expect(err).To(MatchError("Delete: the-error"))
			})
		})
	})

	Describe("Name", func() {
		It("returns the name", func() {
			Expect(httpHealthCheck.Name()).To(Equal(name))
		})
	})

	Describe("Type", func() {
		It("returns the type", func() {
			Expect(httpHealthCheck.Type()).To(Equal("Http Health Check"))
		})
	})
})
