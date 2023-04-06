package main

import "fmt"

func main() {
	peliculas := [10]string{"The Shawshank Redemption", "The Godfather", "The Dark Knight", "The Godfather Part II", "12 Angry Men", "Schindler's List", "The Lord of the Rings: The Return of the King", "Pulp Fiction", "The Lord of the Rings: The Fellowship of the Ring", "The Good, the Bad and the Ugly"}
	fmt.Println(peliculas)
	fmt.Println(peliculas[1])
	fmt.Println(len(peliculas))
}
