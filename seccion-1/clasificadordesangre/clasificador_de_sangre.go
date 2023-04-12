package main

import (
	"fmt"
	"strings"
)

func clasificarSangre(tipoSangre string) string {
	var resultado string
	switch strings.ToUpper(tipoSangre) {
	case "O+":
		resultado = "Grupo sanguíneo O, factor Rh positivo"
	case "O-":
		resultado = "Grupo sanguíneo O, factor Rh negativo"
	case "A+":
		resultado = "Grupo sanguíneo A, factor Rh positivo"
	case "A-":
		resultado = "Grupo sanguíneo A, factor Rh negativo"
	case "B+":
		resultado = "Grupo sanguíneo B, factor Rh positivo"
	case "B-":
		resultado = "Grupo sanguíneo B, factor Rh negativo"
	case "AB+":
		resultado = "Grupo sanguíneo AB, factor Rh positivo"
	case "AB-":
		resultado = "Grupo sanguíneo AB, factor Rh negativo"
	default:
		resultado = "Tipo de sangre equivocado"
	}
	return resultado
}

func main() {
	var tipoSangre string
	fmt.Printf("Ingrese tipo de sangre: ")
	fmt.Scanln(&tipoSangre)

	resultado := clasificarSangre(tipoSangre)
	fmt.Println("Resultado: ", resultado)
}
