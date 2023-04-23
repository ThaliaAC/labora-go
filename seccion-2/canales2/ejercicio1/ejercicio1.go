/*
Ejercicio 1: Suma de números en un rango utilizando gorutinas y canales
Escribe un programa en Go que sume todos los números en un rango dado
(por ejemplo, de 1 a 100) utilizando gorutinas y canales para dividir
el trabajo entre varias gorutinas.
*/
package main

import "fmt"

func sumaNumeros(inicio int, fin int, canal chan int) {
	suma := 0
	for numero := inicio; numero <= fin; numero++ {
		suma += numero
	}
	canal <- suma
}

func main() {
	inicio := 1
	fin := 100

	canal := make(chan int)
	mitad := (inicio + fin) / 2
	go sumaNumeros(inicio, mitad, canal)
	go sumaNumeros(mitad+1, fin, canal)

	suma1 := <-canal
	suma2 := <-canal
	sumaTotal := suma1 + suma2

	fmt.Printf("El resultado de la suma de los números del %v al %v es %v.", inicio, fin, sumaTotal)
}
