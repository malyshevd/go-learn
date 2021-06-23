package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var ints []int
	fmt.Print("Введите массив числе через пробел: ")
	strs := strings.Split(lineScan(), " ")

	for _, s := range strs {
		num, err := strconv.Atoi(s)
		if err == nil {
			ints = append(ints, num)
		}
	}

	insertionSort(ints)

	fmt.Println(ints)
}

func lineScan() string {
	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')

	if err != nil {
		fmt.Println("Ошибка ввода: ", err)
	}
	return str
}

func insertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		x := array[i]
		j := i
		for ; j >= 1 && array[j-1] > x; j-- {
			array[j] = array[j-1]
		}
		array[j] = x
	}
}
