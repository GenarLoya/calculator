package postfix_generator

import (
	"fmt"
	tManager "genarold/calculator/src/tokenManager"
	"genarold/calculator/utils"
)

func PostfixGenerator(units []*tManager.Unit) []string {

	stack := utils.Stack[*tManager.Unit]{}
	postfix := utils.Stack[string]{}

	for _, unit := range units {
		if unit.Is == tManager.Number {
			postfix.Push(unit.Unit)
		} else if unit.Is == tManager.Operator {
			if stack.IsEmpty() {
				stack.Push(unit)
			} else {

				if unit.Unit == "(" {
					stack.Push(unit)
					continue
				}

				fmt.Println()

				if unit.Unit == ")" {
					for {
						val, _ := stack.Pop()
						if val.Unit == "(" {
							break
						}
						postfix.Push(val.Unit)
					}
					continue
				}

				peek, ok := stack.Peek()

				if ok && (unit.Ps.Pe > peek.Ps.Pp) {
					stack.Push(unit)
				} else {
					stack.Pop()

					if ok {
						postfix.Push(peek.Unit)
					}
					stack.Push(unit)
				}
			}
		}
	}

	for !stack.IsEmpty() {
		val, _ := stack.Pop()

		fmt.Println(val.Unit, "VAL FINAL IN STACK")
		postfix.Push(val.Unit)
	}

	fmt.Println(postfix)
	return postfix
}
