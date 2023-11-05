package utils_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "genarold/calculator/utils"
)

func TestTokenizer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tokenizer Suite")
}

var _ = Describe("Tokenizer", func() {
	It("Return list of simple tokens", func() {
		tokens := Tokenizer("1 + 2")
		expectTokens := []string{"1", "+", "2"}

		Expect(tokens).To(Equal(expectTokens))
	})

	It("Return list of simple tokens when number contains more than 2 digits", func() {
		tokens := Tokenizer("14 + 24")
		expectTokens := []string{"14", "+", "24"}

		Expect(tokens).To(Equal(expectTokens))
	})

	It("Return list of simple tokens when contains *", func() {
		tokens := Tokenizer("14 * 4")
		expectTokens := []string{"14", "*", "4"}

		Expect(tokens).To(Equal(expectTokens))
	})

	It("Return list of simple tokens when contains ^", func() {
		tokens := Tokenizer("44 ^ 47")
		expectTokens := []string{"44", "^", "47"}

		Expect(tokens).To(Equal(expectTokens))
	})
})

var _ = Describe("TokenizerValidator", func() {
	testValidate := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Validate 1 + 2",
			input:    "1 + 2",
			expected: true,
		},
		{
			name:     "Validate 1 + 2 * 3",
			input:    "1 + 2 * 3",
			expected: true,
		},
		{
			name:     "Validate 1 + 2 * 3 / 4",
			input:    "1 + 2 * 3 / 4",
			expected: true,
		},
		{
			name:     "Validate 1 + 2 * 3 / (4 ^ 5)",
			input:    "1 + 2 * 3 / (4 ^ 5)",
			expected: true,
		},
		{
			name:     "Validate 1 + 2 * 3 / (4 ^ 5) * 6",
			input:    "1 + 2 * 3 / (4 ^ 5) * 6",
			expected: true,
		},
		//false cases
		{
			name:     "Validate 1 + 2 * 3 / (4 ^ 5) * 6 )",
			input:    "1 + 2 * 3 / (4 ^ 5) * 6 )",
			expected: false,
		},
		{
			name:     "Validate 1 + 2 * 3 / (4 ^ 5) * 6 (",
			input:    "1 + 2 * 3 / (4 ^ 5) * 6 (",
			expected: false,
		},
		{
			name:     "Validate 1 + 2 * 3 / (4 ^ 5) * 6 (7)",
			input:    "1 + 2 * 3 / (4 ^ 5) * 6 (7)",
			expected: false,
		},
		{
			name:     "Validate 1 + 2",
			input:    "1 + 2",
			expected: true,
		},
		{
			name:     "Validate 1 + 2 * 3",
			input:    "1 + 2 * 3",
			expected: true,
		},
		{
			name:     "Validate 1 + 2 * 3 / 4",
			input:    "1 + 2 * 3 / 4",
			expected: true,
		},
		{
			name:     "Validate 1 + 2 * 3 / (4 ^ 5)",
			input:    "1 + 2 * 3 / (4 ^ 5)",
			expected: true,
		},
		{
			name:     "Validate 1 + 2 * 3 / (4 ^ 5) * 6",
			input:    "1 + 2 * 3 / (4 ^ 5) * 6",
			expected: true,
		},
		{
			name:     "Validate 1 + 2 * 3 / (4 ^ 5) * 6 )",
			input:    "1 + 2 * 3 / ((4 ^ 5) * 6 )",
			expected: true,
		},
		{
			name:     "Validate 1 + 2 * 3 / (4 ^ 5) * 6 (",
			input:    "1 + 2 * 3 / (4 ^ 5) * 6 (eee",
			expected: false,
		},
		{
			name:     "Validate 1 + 2 * 3 / (4 ^ 5) * 6 (7)",
			input:    "1 + 2 * 3 / eee(4 ^ 5) * 6 (7)",
			expected: false,
		},
	}

	for _, test := range testValidate {
		test := test

		It(test.name, func() {
			Expect(TokenizerValidator(test.input)).To(Equal(test.expected))
		})
	}
})
