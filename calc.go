package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct {
}

func (calc) operate(input string, operation string) int {
	cleanInput := strings.Split(input, operation)
	operator1 := parser(cleanInput[0])
	operator2 := parser(cleanInput[1])
	switch operation {
	case "+":
		// fmt.Println(operator1 + operator2)
		return operator1 + operator2
	case "-":
		// fmt.Println(operator1 - operator2)
		return operator1 - operator2
	case "*":
		// fmt.Println(operator1 * operator2)
		return operator1 * operator2
	case "/":
		// fmt.Println(operator1 / operator2)
		return operator1 / operator2
	default:
		fmt.Println(operation, "no esta soportado!")
		return 0
	}
}

func parser(input string) int {
	operator, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
	}
	return operator
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	input := readInput()
	operator := readInput()
	c := calc{}
	fmt.Println(c.operate(input, operator))
}
