package main

import "fmt"

type Persona struct {
	nombre   string
	edad     int
	ciudad   string
	telefono string
}

func cambiarCiudad(p *Persona, nuevaCiudad string) {
	p.ciudad = nuevaCiudad
}

func main() {
	persona1 := Persona{"Pedro Gonzales", 35, "Lima", "+51987654321"}
	persona2 := Persona{"Andrea Fernandez", 32, "Arequipa", "+51912345678"}
	fmt.Println("Persona 1: ", persona1)
	fmt.Println("Persona 2: ", persona2)

	nuevaCiudad := "Piura"
	cambiarCiudad(&persona1, nuevaCiudad)
	fmt.Println("Persona 1 con ciudad actualizada: ", persona1)
	fmt.Println("Persona 2: ", persona2)
}
