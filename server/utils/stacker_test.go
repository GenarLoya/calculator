package utils_test

import (
	"testing"

	"genarold/calculator/utils"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestStacker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stacker Suite")
}

var _ = Describe("Stacker", func() {

	stackerPushTests := []struct {
		name     string
		input    utils.Stack[int]
		toPush   []int
		expected utils.Stack[int]
	}{
		{
			name:     "Push 1",
			input:    utils.Stack[int]{},
			toPush:   []int{1},
			expected: utils.Stack[int]{1},
		},
		{
			name:     "Push 2",
			input:    utils.Stack[int]{1},
			toPush:   []int{2},
			expected: utils.Stack[int]{1, 2},
		},
		{
			name:     "Push 2 and 3",
			input:    utils.Stack[int]{1},
			toPush:   []int{2, 3},
			expected: utils.Stack[int]{1, 2, 3},
		},
	}

	for _, test := range stackerPushTests {
		test := test

		It(test.name, func() {
			for _, val := range test.toPush {
				test.input.Push(val)
			}

			Expect(test.input).To(Equal(test.expected))
		})
	}

	stackerPopTests := []struct {
		name     string
		input    utils.Stack[int]
		popValue int
		expected utils.Stack[int]
	}{
		{
			name:     "Pop 1",
			input:    utils.Stack[int]{1, 2, 10},
			popValue: 10,
			expected: utils.Stack[int]{1, 2},
		},
		{
			name:     "Pop 3",
			input:    utils.Stack[int]{1, 2, 3},
			popValue: 3,
			expected: utils.Stack[int]{1, 2},
		},
		{
			name:     "Pop 4",
			input:    utils.Stack[int]{1, 2},
			popValue: 2,
			expected: utils.Stack[int]{1},
		},
	}

	for _, test := range stackerPopTests {
		test := test

		It(test.name, func() {
			val, _ := test.input.Pop()

			Expect(val).To(Equal(test.popValue))
			Expect(test.input).To(Equal(test.expected))
		})
	}

	It("Return false when try pop from empty stack", func() {
		stack := utils.Stack[int]{}

		_, ok := stack.Pop()

		Expect(stack.IsEmpty()).To(Equal(true))
		Expect(ok).To(Equal(false))
	})
})
