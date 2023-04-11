package main

import "fmt"

type Planeta struct {
	nombre string
	lunas  int
}

func main() {
	planeta1 := Planeta{"Mercurio", 0}
	planeta2 := Planeta{"Venus", 0}
	planeta3 := Planeta{"Tierra", 1}
	planeta4 := Planeta{"Marte", 2}
	planeta5 := Planeta{"Júpiter", 79}
	planeta6 := Planeta{"Saturno", 83}
	planeta7 := Planeta{"Urano", 27}
	planeta8 := Planeta{"Neptuno", 14}
	planeta9 := Planeta{"Plutón", 5}

	planetas := []Planeta{planeta1, planeta2, planeta3, planeta4, planeta5, planeta6, planeta7, planeta8, planeta9}

	for i := 0; i < len(planetas); i++ {
		if planetas[i].lunas == 0 {
			fmt.Printf("%d. %s - No tiene lunas\n", i+1, planetas[i].nombre)
		} else if planetas[i].lunas == 1 {
			fmt.Printf("%d. %s - 1 luna\n", i+1, planetas[i].nombre)
		} else {
			fmt.Printf("%d. %s - %d lunas\n", i+1, planetas[i].nombre, planetas[i].lunas)
		}
	}
}
