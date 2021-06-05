package main

import (
	"fmt"
	"math"
)

func main() {
	var a int
start:
	fmt.Print("1. Calculate area of a rectangle. " +
		"\n2. Calculate length and diameter of a circle by area." +
		"\n3. Expand three-digit number into hundreds, tens, units." +
		"\n9. For exit." +
		"\nChoose program:")
	fmt.Scanln(&a)

	switch a {
	case 1:
		sRectangle()
	case 2:
		circleParametersByArea()
	case 3:
		numericParameters()
	case 9:
		break
	default:
		fmt.Println("Program not found. Repeat.")
		goto start
	}
}

func sRectangle() {
	var h, w int

height:
	fmt.Print("Enter height:")
	fmt.Scanln(&h)
	if h == 0 {
		fmt.Println("You entered empty value ")
		goto height
	}

width:
	fmt.Print("Enter width:")
	fmt.Scanln(&w)
	if w == 0 {
		fmt.Println("You entered empty value ")
		goto width
	}

	fmt.Println("Area of rectangle:", h*w)
}

func circleParametersByArea() {
	var s float64

start:
	fmt.Print("Enter area:")
	fmt.Scanln(&s)
	if s == 0 {
		fmt.Println("You entered empty value")
		goto start
	}

	fmt.Println("Circle diameter:", 2*math.Sqrt(s/math.Pi),
		"\nCircle length:", math.Sqrt(s*4*math.Pi))
}

func numericParameters() {
	var a int

start:
	fmt.Print("Enter value:")
	fmt.Scanln(&a)
	if a == 0 || a > 999 {
		fmt.Println("You entered non-three-digit value or empty value")
		goto start
	}

	fmt.Println("Count of hundreds in value:", a/100,
		"\nCount of tens in value:", a/10%10,
		"\nCount of units in value:", a%10)
}
