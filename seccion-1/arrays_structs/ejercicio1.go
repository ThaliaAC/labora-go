package main

import (
	"fmt"
)

type Persona struct {
	nombre   string
	edad     int
	ciudad   string
	telefono int64
}

func incrementarEdad(p *Persona) {
	p.edad++
}

func buscarEdad(a []Persona, edad int) Persona {
	if len(a) == 0 {
		return Persona{}
	}
	if a[0].edad == edad {
		return a[0]
	}
	return buscarEdad(a[1:], edad)
}

func crearPersona(nombre string, edad int, ciudad string, telefono int64) Persona {
	personaNueva := Persona{nombre: nombre, edad: edad, ciudad: ciudad, telefono: telefono}
	return personaNueva
}

func actualizarTelefono(p *Persona, nuevoTelefono int64) {
	p.telefono = nuevoTelefono
}

func main() {
	persona1 := Persona{"Catherine Ayala", 24, "Piura", 999000111}
	persona2 := Persona{"Claudia Martinez", 51, "Lima", 999888777}
	persona3 := Persona{"Alfredo Castillo", 32, "Lima", 999666555}
	persona4 := Persona{"Cristian Bustamante", 43, "Arequipa", 999444333}
	persona5 := Persona{"Juana Fernandez", 29, "Lima", 999111222}

	fmt.Println("Persona 1: ", persona1)
	incrementarEdad(&persona1)
	fmt.Println("Incremento de edad de Persona 1: ", persona1)

	personas := []Persona{persona1, persona2, persona3, persona4, persona5}

	fmt.Println("Buscar Edad 29: ", buscarEdad(personas, 29))
	fmt.Println("Buscar Edad 18: ", buscarEdad(personas, 18))

	nuevaPersona := Persona{"Barbara Estrada", 36, "Lima", 999888111}
	fmt.Println("Nueva Persona: ", nuevaPersona)

	actualizarTelefono(&nuevaPersona, 999111888)
	fmt.Println("Nueva Persona con telefono actualizado: ", nuevaPersona)
}
