package iam_test

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/genevieve/leftovers/aws/iam"
	"github.com/genevieve/leftovers/aws/iam/fakes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ServerCertificate", func() {
	var (
		serverCertificate iam.ServerCertificate
		client            *fakes.ServerCertificatesClient
		name              *string
	)

	BeforeEach(func() {
		client = &fakes.ServerCertificatesClient{}
		name = aws.String("the-name")

		serverCertificate = iam.NewServerCertificate(client, name)
	})

	Describe("Delete", func() {
		It("deletes the server certificate", func() {
			err := serverCertificate.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DeleteServerCertificateCall.CallCount).To(Equal(1))
			Expect(client.DeleteServerCertificateCall.Receives.DeleteServerCertificateInput.ServerCertificateName).To(Equal(name))
		})

		Context("when the client fails", func() {
			BeforeEach(func() {
				client.DeleteServerCertificateCall.Returns.Error = errors.New("banana")
			})

			It("returns the error", func() {
				err := serverCertificate.Delete()
				Expect(err).To(MatchError("Delete: banana"))
			})
		})
	})

	Describe("Name", func() {
		It("returns the identifier", func() {
			Expect(serverCertificate.Name()).To(Equal("the-name"))
		})
	})

	Describe("Type", func() {
		It("returns \"server certificate\"", func() {
			Expect(serverCertificate.Type()).To(Equal("IAM Server Certificate"))
		})
	})
})
