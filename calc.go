package go_calc_package

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct {
}

func (calc) Operate(entry string, operator string) int {

	array := strings.Split(entry, operator)

	num1 := Parser(array[0])
	num2 := Parser(array[1])

	switch strings.ToLower(operator) {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	default:
		fmt.Println("Invalid operator")
		return 0
	}
}

func Parser(entry string) int {

	num, error := strconv.Atoi(entry)

	if error != nil {
		fmt.Println("Error while converting the value : ", entry)
		return 0
	}

	return num
}

func ReadEntry() string {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func ShowMenu() {
	for {
		fmt.Println("")
		fmt.Println("Enter Basic Math Symbol :")
		fmt.Println("")
		fmt.Println("+ Addition")
		fmt.Println("- Subtraction")
		fmt.Println("* Multiplication")
		fmt.Println("/ Division")
		fmt.Println("x Exit")
		fmt.Println("")
		fmt.Print(": ")

		option := ReadEntry()

		switch strings.ToLower(option) {
		case "x":
			return
		case "+", "-", "*", "/":

			fmt.Println("Enter a Math Equation :")
			operation := ReadEntry()

			c := calc{}
			result := c.Operate(operation, option)
			fmt.Println("Result  (", option, ")  : ", result)

		default:
			fmt.Println("Invalid operation")
		}
	}
}

func main() {

	ShowMenu()
}
