package main

import (
	"fmt"
)

func main() {
	var a int = 10
	var b int = 20
	fmt.Printf("Valores iniciales: a = %d , b = %d\n", a, b)
	var ptrA *int = &a
	*ptrA, b = b, *ptrA
	fmt.Printf("Valores intercambiados: a = %d, b = %d\n", a, b)
}