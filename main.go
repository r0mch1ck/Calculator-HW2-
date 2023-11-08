package main

import (
	"deque/deq"
	"deque/deqchar"
	"fmt"
)

func isNumber(symbol uint8) bool {
	returnValue := false
	if symbol == '0' || symbol == '1' || symbol == '2' || symbol == '3' || symbol == '4' || symbol == '5' || symbol == '6' || symbol == '7' || symbol == '8' || symbol == '9' {
		returnValue = true
	}
	return returnValue
}

func isOperator(symbol uint8) bool {
	returnValue := false
	if symbol == '+' || symbol == '-' || symbol == '*' || symbol == '/' {
		returnValue = true
	}
	return returnValue
}

func operation(a int, b int, action uint8) int {
	result := 0
	if action == '+' {
		result = a + b
	} else if action == '-' {
		result = a - b
	} else if action == '*' {
		result = a * b
	} else if action == '/' {
		result = a / b
	}
	return result
}

func stringToNum(symbol uint8) int {
	returnValue := -1
	if isNumber(symbol) {
		if symbol == '0' {
			returnValue = 0
		} else if symbol == '1' {
			returnValue = 1
		} else if symbol == '2' {
			returnValue = 2
		} else if symbol == '3' {
			returnValue = 3
		} else if symbol == '4' {
			returnValue = 4
		} else if symbol == '5' {
			returnValue = 5
		} else if symbol == '6' {
			returnValue = 6
		} else if symbol == '7' {
			returnValue = 7
		} else if symbol == '8' {
			returnValue = 8
		} else if symbol == '9' {
			returnValue = 9
		}
	}
	return returnValue
}

func main() {
	inputValue := "2+2"
	stackNumbers := deq.ZeroDeque()
	stackOperations := deqchar.ZeroDequeChar()
	operations := make(map[uint8]int)
	operations[' '] = 0
	operations['('] = 0
	operations[')'] = 0
	operations['+'] = 1
	operations['-'] = 1
	operations['*'] = 2
	operations['/'] = 2
	//fmt.Scanf("%s", &inputValue)
	str := inputValue
	for i := 0; i < len(str); i++ {
		if isNumber(str[i]) {
			stackNumbers.AppendRight(stringToNum(str[i]))
			i++
			for i < len(str) && isNumber(str[i]) {
				stackNumbers.Tail.Value = stackNumbers.Tail.Value*10 + stringToNum(str[i])
				i++
			}
		}
		if str[i] == '(' {
			stackOperations.AppendRight('(')
		} else if isOperator(str[i]) {
			if deqchar.IsDequeCharEmpty(stackOperations) || operations[str[i]] > operations[stackOperations.Tail.Value] {
				stackOperations.AppendRight(str[i])
			} else {
				for !deqchar.IsDequeCharEmpty(stackOperations) || !(operations[str[i]] > operations[stackOperations.Tail.Value]) {
					tmp := stackNumbers.Tail.Value
					stackNumbers.PopRight()
					stackNumbers.Tail.Value = operation(stackNumbers.Tail.Value, tmp, stackOperations.Tail.Value)
					stackOperations.PopRight()
				}
			}
		} else if str[i] == ')' {
			for !('(' == stackOperations.Tail.Value) {
				tmp := stackNumbers.Tail.Value
				stackNumbers.PopRight()
				stackNumbers.Tail.Value = operation(stackNumbers.Tail.Value, tmp, stackOperations.Tail.Value)
				stackOperations.PopRight()
			}
			stackOperations.PopRight()
		}
	}
	if deq.IsDequeEmpty(stackNumbers) && deqchar.IsDequeCharEmpty(stackOperations) {
		for !deq.IsDequeEmpty(stackNumbers) {
			tmp := stackNumbers.Tail.Value
			stackNumbers.PopRight()
			stackNumbers.Tail.Value = operation(stackNumbers.Tail.Value, tmp, stackOperations.Tail.Value)
			stackOperations.PopRight()
		}
	}
	fmt.Print(stackNumbers.Tail.Value)
}
