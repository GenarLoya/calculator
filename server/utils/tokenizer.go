package utils

import (
	"errors"
	"regexp"
	"strings"
)

// Tokenizer receives a string and returns a slice of strings representing the tokens found in the input.
// Tokens are defined as numbers, operators (+, -, *, /, ^), and parentheses.
// The function uses a regular expression to find all the tokens in the input string.
func Tokenizer(text string) []string {

	regex := regexp.MustCompile(`\d+|\+|\-|\*|\/|\(|\)|\^`)

	tokens := regex.FindAllString(text, -1)

	return tokens
}

func TokenizerValidator(text string) bool {
	// Remove spaces
	text = strings.ReplaceAll(text, " ", "")

	// Check for invalid characters
	invalidChars := regexp.MustCompile(`[^0-9+\-*/^()]`)
	if invalidChars.MatchString(text) {
		return false
	}

	invalidOpenParenthesesNotClose := regexp.MustCompile(`\(\)`)

	if invalidOpenParenthesesNotClose.MatchString(text) {
		return false
	}

	// Check for consecutive invalid signs
	invalidSigns := regexp.MustCompile(`[+\-*\/^]{2,}`)
	if invalidSigns.MatchString(text) {
		return false
	}

	// Check for valid operators
	validOperators := regexp.MustCompile(`[+*/^-]`)
	if !validOperators.MatchString(text) {
		return false
	}

	// Check for invalid placement of parentheses
	invalidOpenParentheses := regexp.MustCompile(`\d+\(`)
	if invalidOpenParentheses.MatchString(text) {
		return false
	}

	invalidCloseParentheses := regexp.MustCompile(`\)\d+`)
	if invalidCloseParentheses.MatchString(text) {
		return false
	}

	// Check for balanced parentheses
	parenthesesCount := 0
	for _, char := range text {
		if char == '(' {
			parenthesesCount++
		} else if char == ')' {
			parenthesesCount--
			if parenthesesCount < 0 {
				return false
			}
		}
	}

	return parenthesesCount == 0
}

func TokenManager(text string) ([]string, error) {

	tokens := Tokenizer(text)

	if !TokenizerValidator(text) {
		return nil, errors.New("invalid expression")
	}

	return tokens, nil
}
