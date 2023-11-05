package postfix_generator

import (
	"fmt"
	tManager "genarold/calculator/src/tokenManager"
	"genarold/calculator/utils"
	"math"
	"strconv"
)

func PostfixGenerator(units []*tManager.Unit) []*tManager.Unit {

	stack := utils.Stack[*tManager.Unit]{}
	postfix := utils.Stack[*tManager.Unit]{}

	for _, unit := range units {
		if unit.Is == tManager.Number {
			postfix.Push(unit)
		} else if unit.Is == tManager.Operator {
			if stack.IsEmpty() {
				stack.Push(unit)
			} else {

				if unit.Unit == "(" {
					stack.Push(unit)
					continue
				}

				if unit.Unit == ")" {
					for {
						val, _ := stack.Pop()
						if val.Unit == "(" {
							break
						}
						postfix.Push(val)
					}
					continue
				}

				peek, ok := stack.Peek()

				if ok && (unit.Ps.Pe > peek.Ps.Pp) {
					stack.Push(unit)
				} else {
					stack.Pop()

					if ok {
						postfix.Push(peek)
					}
					stack.Push(unit)
				}
			}
		}
	}

	for !stack.IsEmpty() {
		val, _ := stack.Pop()

		postfix.Push(val)
	}

	return postfix
}

func PostFixCalculator(units []*tManager.Unit) (float64, error) {
	stack := utils.Stack[*tManager.Unit]{}
	result := 0.0

	for _, unit := range units {
		if unit.Is == tManager.Number {
			stack.Push(unit)
		} else if unit.Is == tManager.Operator {
			second, _ := stack.Pop()
			first, _ := stack.Pop()

			secondInt, errSecond := strconv.ParseFloat(second.Unit, 64)

			if errSecond != nil {
				return 0.0, errSecond
			}

			firstInt, errFirst := strconv.ParseFloat(first.Unit, 64)

			if errFirst != nil {
				return 0.0, errFirst
			}

			switch unit.Unit {
			case "+":
				result = firstInt + secondInt
			case "-":
				result = firstInt - secondInt
			case "*":
				result = firstInt * secondInt
			case "/":
				result = firstInt / secondInt
			case "^":
				result = math.Pow(firstInt, secondInt)
			}

			stack.Push(&tManager.Unit{
				Is:   tManager.Number,
				Unit: fmt.Sprintf("%f", result),
			})
		}
	}

	return result, nil
}

func PostfixManager(text string) (float64, error) {

	tokens, errFormat := tManager.TokenManager(text)

	if errFormat != nil {
		return 0.0, errFormat
	}

	postfix := PostfixGenerator(tokens)

	result, err := PostFixCalculator(postfix)

	if err != nil {
		return 0.0, err
	}

	return result, nil

}