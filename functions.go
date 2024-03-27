package main

import (
	f "fmt"
	add "rahulchhabra.io/calc/adder";
	mul "rahulchhabra.io/calc/multiplier"
)

func greet() {
	f.Println("Hello World!")
	f.Println(url)
	f.Println(add.Add(1, 2))
	f.Println(mul.Mul(2, 3))
}
