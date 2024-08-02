package stack_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestStack(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stack Suite")
}

//var _ = Describe("Test Stack", func() {
//	When("Using integers", func() {
//		var integerStack *stack.Stack[int]
//
//		BeforeEach(func() {
//			integerStack = stack.From(1)
//		})
//	})
//})
