package main

import (
	f "fmt"

	"rahulchhabra.io/calc"
	"rahulchhabra.io/hello"
	"rahulchhabra.io/hello/greeter"
)

func greet() {
	f.Println("Hello World!")
	f.Println(url)
	f.Println(calc.Add(1, 2))
	f.Println(calc.Mul(2, 3))
	f.Println(hello.Hello())
	f.Println(mine.Doublegreeter());
	hello.Header();
}
