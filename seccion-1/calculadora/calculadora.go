package main
import "fmt"

func calcular(a *int, b *int) (int, int, int, float64) {
	sum := *a + *b
	rest := *a - *b
	mult := *a * *b
	div := float64(*a / *b)
	fmt.Printf("Suma: %d\n", sum) 
	fmt.Printf("Resta: %d\n", rest) 
	fmt.Printf("Multiplicación: %d\n", mult) 
	fmt.Printf("División: %.2f\n", div) 
	return sum, rest, mult, div
}

func main() {
	var primerNum int
	var p1 *int
	var segundoNum int
	var p2 *int
    fmt.Print("Ingrese el primer número: ")
    fmt.Scanln(&primerNum)
    fmt.Print("Ingrese el segundo número: ")
	fmt.Scanln(&segundoNum)
	p1 = &primerNum
	p2 = &segundoNum
	calcular(p1, p2)
}