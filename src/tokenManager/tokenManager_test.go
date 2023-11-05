package token_manager_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	tManager "genarold/calculator/src/tokenManager"
	"testing"
)

func TestTokenManger(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tokenizer Suite")
}

var _ = Describe("New Units", func() {
	tokenToUnittest := []struct {
		name     string
		input    string
		expected tManager.Unit
	}{
		{
			name:  "Teste 1",
			input: "1",
			expected: tManager.Unit{
				Unit: "1",
				Ps: tManager.Ps{
					Pe: 0,
					Pp: 0,
				},
				Is: tManager.Number,
			},
		},
		{
			name:  "Teste 2",
			input: "+",
			expected: tManager.Unit{
				Unit: "+",
				Ps: tManager.Ps{
					Pe: 1,
					Pp: 1,
				},
				Is: tManager.Operator,
			},
		},
	}

	for _, test := range tokenToUnittest {
		test := test

		It(test.name, func() {
			token := tManager.NewUnit(test.input)

			Expect(token.Unit).To(Equal(test.expected.Unit))
			Expect(token.Is).To(Equal(test.expected.Is))
			Expect(token.Ps.Pe).To(Equal(test.expected.Ps.Pe))
			Expect(token.Ps.Pp).To(Equal(test.expected.Ps.Pp))
		})
	}

})

var _ = Describe("Token Manager", func() {
	testExpressionToUnits := []struct {
		name     string
		input    string
		expected []tManager.Unit
	}{
		{
			name:  "Teste 1",
			input: "1 + 2",
			expected: []tManager.Unit{
				{
					Unit: "1",
					Ps: tManager.Ps{
						Pe: 0,
						Pp: 0,
					},
					Is: tManager.Number,
				},
				{
					Unit: "+",
					Ps: tManager.Ps{
						Pe: 1,
						Pp: 1,
					},
					Is: tManager.Operator,
				},
				{
					Unit: "2",
					Ps: tManager.Ps{
						Pe: 0,
						Pp: 0,
					},
					Is: tManager.Number,
				},
			},
		},
		{
			name:  "Teste 2",
			input: "1 + (2 * 3)",
			expected: []tManager.Unit{
				{
					Unit: "1",
					Ps: tManager.Ps{
						Pe: 0,
						Pp: 0,
					},
					Is: tManager.Number,
				},
				{
					Unit: "+",
					Ps: tManager.Ps{
						Pe: 1,
						Pp: 1,
					},
					Is: tManager.Operator,
				},
				{
					Unit: "(",
					Ps: tManager.Ps{
						Pe: 5,
						Pp: 0,
					},
					Is: tManager.Operator,
				},
				{
					Unit: "2",
					Ps: tManager.Ps{
						Pe: 0,
						Pp: 0,
					},
					Is: tManager.Number,
				},
				{
					Unit: "*",
					Ps: tManager.Ps{
						Pe: 2,
						Pp: 2,
					},
					Is: tManager.Operator,
				},
				{
					Unit: "3",
					Ps: tManager.Ps{
						Pe: 0,
						Pp: 0,
					},
					Is: tManager.Number,
				},
				{
					Unit: ")",
					Ps: tManager.Ps{
						Pe: -1,
						Pp: -1,
					},
					Is: tManager.Operator,
				},
			},
		},
	}

	for _, test := range testExpressionToUnits {
		test := test

		It(test.name, func() {
			tokens, _ := tManager.TokenManager(test.input)

			for i, token := range tokens {
				Expect(token.Unit).To(Equal(test.expected[i].Unit))
				Expect(token.Is).To(Equal(test.expected[i].Is))
				Expect(token.Ps.Pe).To(Equal(test.expected[i].Ps.Pe))
				Expect(token.Ps.Pp).To(Equal(test.expected[i].Ps.Pp))
			}
		})

		testErrorCases := []struct {
			name  string
			input string
		}{
			{
				name:  "Teste 1",
				input: "1 + 2 * 3 )",
			},
			{
				name:  "Teste 2",
				input: "1 + 2 * 3 (",
			},
			{
				name:  "Teste 3",
				input: "1 + 2(6) + 3",
			},
			{
				name:  "Teste 4",
				input: "1 + 2 * (6 + 3",
			},
		}

		for _, test := range testErrorCases {
			test := test

			It(test.name, func() {
				_, err := tManager.TokenManager(test.input)

				Expect(err).To(HaveOccurred())
			})
		}
	}
})
