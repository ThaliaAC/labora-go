package main

import (
	"fmt"
	"strings"
)

var (
	coins        = 50
	users        = []string{"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth"}
	distribution = make(map[string]int)
)

func main() {
	for _, user := range users {
		count := 0
		for _, letter := range strings.ToUpper(user) {
			switch letter {
			case 'A':
				count++
				coins--
			case 'E':
				count++
				coins--
			case 'I':
				count += 2
				coins -= 2
			case 'O':
				count += 3
				coins -= 3
			case 'U':
				count += 4
				coins -= 4
			}
		}
		if count > 10 {
			x := count - 10
			count = 10
			coins = coins + x
		}
		distribution[user] = count
	}

	fmt.Println("User counts:", distribution)
	fmt.Println("Coins left:", coins)
}
