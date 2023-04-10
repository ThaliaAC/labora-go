package main

import "fmt"

func main() {
	var day int
	fmt.Print("Ingresa un número del 1 al 7: ")
	fmt.Scanln(&day)
	switch day {
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miercoles")
	case 4:
		fmt.Println("Jueves")
	case 5:
		fmt.Println("Viernes")
	case 6:
		fmt.Println("Sábado")
	case 7:
		fmt.Println("Domingo")
	default:
		fmt.Println("Número inválido.")
	}
}
