package main

import (
	"calculator/calc"
	
	"fmt"
)


func main() {
	fmt.Println(calc.Mult(2, 3))
	fmt.Println(calc.Div(8, 4))
	fmt.Println(calc.Add(8, 4))
	fmt.Println(calc.Sub(8, 4))
}