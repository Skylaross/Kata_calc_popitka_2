package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Введите пример:")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		pars(text)
	}
}

func pars(task string) {
	var rome []int
	var Kudanam int

	zadacha := strings.Split(task, " ")
	if len(zadacha) < 3 {
		panic(wrongelements)
	}
	if len(zadacha) > 3 {
		panic(wrongelements)
	}

	num1, err := strconv.Atoi(zadacha[0])
	if err != nil {
		Kudanam++
	}

	checkOperator(zadacha[1])

	num2, err := strconv.Atoi(zadacha[2])
	if err != nil {
		Kudanam++
	}

	switch Kudanam {
	case 0:
		errCheck := (num1 < 0 || num1 > 11) || (num2 < 0 || num2 > 11)
		if errCheck {
			panic(acceptable)
		} else {
			fmt.Println(calculate(num1, num2, zadacha[1]))
		}
	case 1:
		panic(romearabian)
	case 2:
		for i, elem := range zadacha {
			if i == 1 {
				continue
			}
			if val, ok := Rome1[elem]; ok && val > 0 && val < 11 {
				rome = append(rome, val)
			} else {
				panic(acceptable)
			}

		}
		romeorarabian(calculate(rome[0], rome[1], zadacha[1]))
	}
}

func checkOperator(op string) {
	if op == "+" || op == "-" || op == "*" || op == "/" {
		return
	}
	panic(operator)
}

func romeorarabian(result int) {
	var romanElem string
	if result == 0 {
		panic(nozero)
	} else if result < 0 {
		panic(negative)
	} else {
		for _, elem := range convinrome {
			for i := elem; i <= result; {
				for idx, value := range Rome1 {
					if value == elem {
						romanElem += idx
						result -= elem
					}
				}
			}
		}
	}
	fmt.Println(romanElem)
}

func calculate(a int, b int, oper string) int {
	switch oper {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		return 0
	}
}

var Rome1 = map[string]int{
	"C":    100,
	"XC":   90,
	"L":    50,
	"XL":   40,
	"X":    10,
	"IX":   9,
	"VIII": 8,
	"VII":  7,
	"VI":   6,
	"V":    5,
	"IV":   4,
	"III":  3,
	"II":   2,
	"I":    1,
}
var convinrome = []int{
	100,
	90,
	50,
	40,
	10,
	9,
	8,
	7,
	6,
	5,
	4,
	3,
	2,
	1,
}

const (
	operator      = "Допустимые операции + - / *"
	wrongelements = "Вводите элементы через пробел, в формате: 2 + 1"
	romearabian   = "Можно использовать одновременно только целые Арабские или Римские числа!"
	negative      = "В римской системе нет отрицательных чисел!"
	nozero        = "В римской системе нет числа 0!"
	acceptable    = "Калькулятор принимает на вход целые числа от 1 до 10"
)
