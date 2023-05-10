package s3_test

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/genevieve/leftovers/aws/s3"
	"github.com/genevieve/leftovers/aws/s3/fakes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bucket", func() {
	var (
		bucket s3.Bucket
		client *fakes.BucketsClient
		name   *string
	)

	BeforeEach(func() {
		client = &fakes.BucketsClient{}
		name = aws.String("the-name")

		bucket = s3.NewBucket(client, name)
	})

	Describe("Delete", func() {
		It("deletes the bucket", func() {
			err := bucket.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DeleteBucketCall.CallCount).To(Equal(1))
			Expect(client.DeleteBucketCall.Receives.DeleteBucketInput.Bucket).To(Equal(name))
		})

		Context("the client fails", func() {
			BeforeEach(func() {
				client.DeleteBucketCall.Returns.Error = errors.New("banana")
			})

			It("returns the error", func() {
				err := bucket.Delete()
				Expect(err).To(MatchError("Delete: banana"))
			})
		})
	})

	Describe("Name", func() {
		It("returns the identifier", func() {
			Expect(bucket.Name()).To(Equal("the-name"))
		})
	})

	Describe("Type", func() {
		It("returns the type", func() {
			Expect(bucket.Type()).To(Equal("S3 Bucket"))
		})
	})
})
