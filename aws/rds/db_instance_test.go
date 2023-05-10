package rds_test

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/genevieve/leftovers/aws/rds"
	"github.com/genevieve/leftovers/aws/rds/fakes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DBInstance", func() {
	var (
		dbInstance   rds.DBInstance
		client       *fakes.DbInstancesClient
		name         *string
		skipSnapshot *bool
	)

	BeforeEach(func() {
		client = &fakes.DbInstancesClient{}
		name = aws.String("the-name")
		skipSnapshot = aws.Bool(true)

		dbInstance = rds.NewDBInstance(client, name)
	})

	Describe("Delete", func() {
		It("deletes the db instance", func() {
			err := dbInstance.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DeleteDBInstanceCall.CallCount).To(Equal(1))
			Expect(client.DeleteDBInstanceCall.Receives.DeleteDBInstanceInput.DBInstanceIdentifier).To(Equal(name))
			Expect(client.DeleteDBInstanceCall.Receives.DeleteDBInstanceInput.SkipFinalSnapshot).To(Equal(skipSnapshot))
		})

		Context("when the client fails", func() {
			BeforeEach(func() {
				client.DeleteDBInstanceCall.Returns.Error = errors.New("banana")
			})

			It("returns the error", func() {
				err := dbInstance.Delete()
				Expect(err).To(MatchError("Delete: banana"))
			})
		})
	})

	Describe("Name", func() {
		It("returns the identifier", func() {
			Expect(dbInstance.Name()).To(Equal("the-name"))
		})
	})

	Describe("Type", func() {
		It("returns the type", func() {
			Expect(dbInstance.Type()).To(Equal("RDS DB Instance"))
		})
	})
})
