package postfix_generator_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	pGen "genarold/calculator/src/postfixGenerator"
	tManager "genarold/calculator/src/tokenManager"
)

func TestPostFixGenerator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Token manager Suite")
}

var _ = Describe("PostfixGenarator", func() {
	It("Return postfix expression '1+2' => '12+'", func() {
		tokens, _ := tManager.TokenManager("1+2")
		postfix := pGen.PostfixGenerator(tokens)

		expectTokens := []string{"1", "2", "+"}

		for i, token := range postfix {
			Expect(token.Unit).To(Equal(expectTokens[i]))
		}
	})

	It("Return postfix expression '3+4*(2/2/2)' => '34222//*+'", func() {
		tokens, _ := tManager.TokenManager("3+4*(2/2/2)")
		postfix := pGen.PostfixGenerator(tokens)

		expectTokens := []string{"3", "4", "2", "2", "/", "2", "/", "*", "+"}

		for i, token := range postfix {
			Expect(token.Unit).To(Equal(expectTokens[i]))
		}
	})

	It("Return postfix expression '34*(3+1)' => '34*31+'", func() {
		tokens, _ := tManager.TokenManager("3*4/ (3+1)")
		postfix := pGen.PostfixGenerator(tokens)

		expectTokens := []string{"3", "4", "*", "3", "1", "+", "/"}

		for i, token := range postfix {
			Expect(token.Unit).To(Equal(expectTokens[i]))
		}
	})
})

var _ = Describe("Src/InfixGenerator/InfixGeneratorValidator", func() {

	testPostfixCalc := []struct {
		name     string
		input    string
		expected float64
	}{
		{
			name:     "Calculate 1 + 2",
			input:    "1+2",
			expected: 3,
		},
		{
			name:     "Calculate 1 + 2.2",
			input:    "1+2.2",
			expected: 3.2,
		},
		{
			name:     "Calculate 3*4*(1 + 2)",
			input:    "3*4*(1 + 2)",
			expected: 36,
		},
		{
			name:     "Calculate 3*4/(1 + 2)",
			input:    "3*4/(1 + 2)",
			expected: 4,
		},
		{
			name:     "Calculate 3*4/(3 + 1)",
			input:    "3*4/(3 + 1)",
			expected: 3,
		},
		{
			name:     "Calculate 3*4/(3 + (1+5*3))",
			input:    "3*4/(3 + (1+5*3))",
			expected: 0.631578947368421,
		},
		{
			name:     "Calculate 3*4/(3 + (1+5*3))",
			input:    "3*4/(3 / (1+5*3) + 6/3)",
			expected: 5.4857142857142857,
		},
	}

	for _, test := range testPostfixCalc {
		It(test.name, func() {
			tokens, _ := tManager.TokenManager(test.input)
			postfix := pGen.PostfixGenerator(tokens)

			result, _ := pGen.PostFixCalculator(postfix)

			Expect(result).To(Equal(test.expected))
		})
	}
})
