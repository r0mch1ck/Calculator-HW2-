package main

import (
	"deque/deq"
	"deque/deqchar"
	"fmt"
	"math"
)

func inputStr(str *string) (int, error) {
	print("You can use these operators: +, -, *, /, ^.\nDivision performed integer.\nEnter your expression:\n")
	return fmt.Scanf("%s", str)
}

func inputValidation(scanResult int, ScanError error) {
	if scanResult != 1 || ScanError != nil {
		panic("Input Error")
	}
}

func mapOperationMaker() map[uint8]int {
	operations := make(map[uint8]int)
	operations['('] = 0
	operations[')'] = 0
	operations['+'] = 1
	operations['-'] = 1
	operations['*'] = 2
	operations['/'] = 2
	operations['^'] = 3
	return operations
}

func isRightValue(symbol uint8) bool {
	if symbol == '^' ||
		symbol == '(' ||
		symbol == ')' ||
		symbol == '1' ||
		symbol == '2' ||
		symbol == '3' ||
		symbol == '4' ||
		symbol == '5' ||
		symbol == '6' ||
		symbol == '7' ||
		symbol == '8' ||
		symbol == '9' ||
		symbol == '+' ||
		symbol == '-' ||
		symbol == '*' ||
		symbol == '/' {
		return true
	} else {
		return false
	}
}

func isNumber(symbol uint8) bool {
	returnValue := false
	if symbol == '0' || symbol == '1' || symbol == '2' || symbol == '3' || symbol == '4' || symbol == '5' || symbol == '6' || symbol == '7' || symbol == '8' || symbol == '9' {
		returnValue = true
	}
	return returnValue
}

func isOperator(symbol uint8) bool {
	returnValue := false
	if symbol == '^' || symbol == '+' || symbol == '-' || symbol == '*' || symbol == '/' {
		returnValue = true
	}
	return returnValue
}

func isBracket(symbol uint8) bool {
	return (symbol == '(') || (symbol == ')')
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
	} else if action == '^' {
		result = int(math.Pow(float64(a), float64(b)))
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

func clearStack(stackNumbers *deq.Deque, stackOperations *deqchar.DequeChar) {
	tmp := stackNumbers.Tail.Value
	stackNumbers.PopRight()
	stackNumbers.Tail.Value = operation(stackNumbers.Tail.Value, tmp, stackOperations.Tail.Value)
	stackOperations.PopRight()
}

func finalClearStack(stackNumbers *deq.Deque, stackOperations *deqchar.DequeChar) {
	for !deqchar.IsDequeCharEmpty(stackOperations) {
		clearStack(stackNumbers, stackOperations)
	}
}

func checker(stackOperations *deqchar.DequeChar, operations map[uint8]int, str uint8) bool {
	if deqchar.IsDequeCharEmpty(stackOperations) {
		return false
	} else if operations[str] > operations[stackOperations.Tail.Value] {
		return false
	} else {
		return true
	}
}

func workWithNum(stackNumbers *deq.Deque, i *int, str string) {
	stackNumbers.AppendRight(stringToNum(str[*i]))
	*i++
	if *i < len(str) {
		for *i < len(str) && isNumber(str[*i]) {
			stackNumbers.Tail.Value = stackNumbers.Tail.Value*10 + stringToNum(str[*i])
			*i++
		}
	}
}

func workWithBracket(str string, i *int, stackNumbers *deq.Deque, stackOperations *deqchar.DequeChar) {
	if str[*i] == '(' {
		stackOperations.AppendRight('(')
	} else if str[*i] == ')' {
		for '(' != stackOperations.Tail.Value {
			clearStack(stackNumbers, stackOperations)
		}
		stackOperations.PopRight()
	}
}

func workWithOperator(operations map[uint8]int, str string, i *int, stackNumbers *deq.Deque, stackOperations *deqchar.DequeChar) {
	if deqchar.IsDequeCharEmpty(stackOperations) {
		stackOperations.AppendRight(str[*i])
	} else if operations[str[*i]] > operations[stackOperations.Tail.Value] {
		stackOperations.AppendRight(str[*i])
	} else {
		for checker(stackOperations, operations, str[*i]) {
			clearStack(stackNumbers, stackOperations)
		}
		stackOperations.AppendRight(str[*i])
	}
}

func runStrTreatment(operations map[uint8]int, str string, stackNumbers *deq.Deque, stackOperations *deqchar.DequeChar, i *int) {
	if isNumber(str[*i]) {
		workWithNum(stackNumbers, i, str)
		if *i >= len(str) {
			return
		}
	}
	if isBracket(str[*i]) {
		workWithBracket(str, i, stackNumbers, stackOperations)
	} else if isOperator(str[*i]) {
		workWithOperator(operations, str, i, stackNumbers, stackOperations)
	}
}

func main() {
	var scanResult int
	var ScanError error
	var str string
	stackNumbers := deq.ZeroDeque()
	stackOperations := deqchar.ZeroDequeChar()
	operations := mapOperationMaker()
	scanResult, ScanError = inputStr(&str)
	inputValidation(scanResult, ScanError)
	for i := 0; i < len(str); i++ {
		if isRightValue(str[i]) {
			runStrTreatment(operations, str, stackNumbers, stackOperations, &i)
		} else {
			panic("You write wrong value!")
		}
	}
	finalClearStack(stackNumbers, stackOperations)
	fmt.Print(stackNumbers.Tail.Value)
}
