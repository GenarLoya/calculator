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
	RunSpecs(t, "Tokenizer Suite")
}

var _ = Describe("Src/InfixGenerator", func() {
	It("Return postfix expression '1+2' => '12+'", func() {
		tokens, _ := tManager.TokenManager("1 + 2")
		postfix := pGen.PostfixGenerator(tokens)

		expectTokens := []string{"1", "2", "+"}

		Expect(postfix).To(Equal(expectTokens))
	})

	It("Return postfix expression '3+4*(2/2/2)' => '34222//*+'", func() {
		tokens, _ := tManager.TokenManager("3+4*(2/2/2)")
		postfix := pGen.PostfixGenerator(tokens)

		expectTokens := []string{"3", "4", "2", "2", "/", "2", "/", "*", "+"}

		Expect(postfix).To(Equal(expectTokens))
	})

	It("Return postfix expression '34*(3+1)' => '34*31+'", func() {
		tokens, _ := tManager.TokenManager("3*4/ (3+1)")
		postfix := pGen.PostfixGenerator(tokens)

		expectTokens := []string{"3", "4", "*", "3", "1", "+", "/"}

		Expect(postfix).To(Equal(expectTokens))
	})
})
