package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"math"
	"os"
)

func main() {
	var a, b, res float64
	var op string

aEnter:
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)
	if a == 0 {
		fmt.Println("Введено пустое значение или 0. Повторите ввод.")
		goto aEnter
	}

bEnter:
	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)
	if b == 0 {
		fmt.Println("Введено пустое значение или 0. Повторите ввод.")
		goto bEnter
	}

opEnter:
	fmt.Print("Введите арифметическую операцию (+, -, *, /, %, ^): ")
	fmt.Scanln(&op)

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	case "%":
		res = math.Mod(a, b)
	case "^":
		res = math.Pow(a, b)
	default:
		fmt.Println("Операция выбрана неверно. Повторите ввод операции.")
		goto opEnter
	}

	fmt.Printf("Результат выполнения операции: %f\nДля продолжения расчетов нажмите Enter. \nДля выхода из программы нажмите Ctrl+Q.\n", res)

	_, key, err := keyboard.GetSingleKey()
	if err != nil {
		panic(err)
	}
	if keyboard.KeyEnter == key {
		a = res
		goto bEnter
	}
	if keyboard.KeyCtrlQ == key {
		os.Exit(1)
	}
}
