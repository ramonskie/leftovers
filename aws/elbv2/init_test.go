package elbv2_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestELBV2(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "aws/elbv2")
}
