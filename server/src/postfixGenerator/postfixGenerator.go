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
				fmt.Println("NUMBER")
				stack.Push(unit)
			} else {
				fmt.Println("OPERATOR")

				if unit.Unit == "(" {
					stack.Push(unit)
					continue
				}

				if unit.Unit == ")" {
					fmt.Println("PARENT")
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
					fmt.Println("PSPE")
					stack.Pop()

					if ok {
						postfix.Push(peek)
					}
					stack.Push(unit)
				}
			}
		}
	}

	fmt.Println("HOLA", postfix, stack)

	for !stack.IsEmpty() {
		val, _ := stack.Pop()

		postfix.Push(val)
	}

	fmt.Println("HOLA2")
	return postfix
}

func PostFixCalculator(units []*tManager.Unit) (float64, error) {
	stack := utils.Stack[*tManager.Unit]{}
	result := 0.0

	for _, unit := range units {
		fmt.Println(unit)
	}

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

	fmt.Println(text)

	postfix := PostfixGenerator(tokens)

	for _, unit := range postfix {
		fmt.Println(unit)
	}

	result, err := PostFixCalculator(postfix)

	println("HOLA 3")

	if err != nil {
		return 0.0, err
	}

	return result, nil

}
