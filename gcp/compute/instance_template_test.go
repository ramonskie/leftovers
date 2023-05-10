package compute_test

import (
	"errors"

	"github.com/genevieve/leftovers/gcp/compute"
	"github.com/genevieve/leftovers/gcp/compute/fakes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("InstanceTemplate", func() {
	var (
		client *fakes.InstanceTemplatesClient
		name   string

		instanceTemplate compute.InstanceTemplate
	)

	BeforeEach(func() {
		client = &fakes.InstanceTemplatesClient{}
		name = "banana"

		instanceTemplate = compute.NewInstanceTemplate(client, name)
	})

	Describe("Delete", func() {
		It("deletes the instance template", func() {
			err := instanceTemplate.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DeleteInstanceTemplateCall.CallCount).To(Equal(1))
			Expect(client.DeleteInstanceTemplateCall.Receives.Template).To(Equal(name))
		})

		Context("when the client fails to delete", func() {
			BeforeEach(func() {
				client.DeleteInstanceTemplateCall.Returns.Error = errors.New("the-error")
			})

			It("returns the error", func() {
				err := instanceTemplate.Delete()
				Expect(err).To(MatchError("Delete: the-error"))
			})
		})
	})

	Describe("Name", func() {
		It("returns the name", func() {
			Expect(instanceTemplate.Name()).To(Equal(name))
		})
	})

	Describe("Type", func() {
		It("returns the type", func() {
			Expect(instanceTemplate.Type()).To(Equal("Instance Template"))
		})
	})
})
