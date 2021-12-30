package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calculator struct {
}

func (Calculator) operate(data string, operator string) int {
	entradaLimpia := strings.Split(data, operator)
	operador1 := parsear(entradaLimpia[0])
	operador2 := parsear(entradaLimpia[1])
	switch operator {
	case "+":
		return operador1 + operador2
	case "-":
		return operador1 - operador2
	case "*":
		return operador1 * operador2
	case "/":
		return operador1 / operador2
	default:
		fmt.Println("Invalid operator")
		return 0
	}
}

func parsear(data string) int {
	operador1, _ := strconv.Atoi(data)
	return operador1
}

func readEntry() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
