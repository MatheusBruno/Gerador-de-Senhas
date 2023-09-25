package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {

	var presenha string
	var newsenha string

	for i := 0; i <= 6; i++ {
		presenha += senha(rand.Intn(20))
	}
	rest := strings.Split(presenha, "")

	for i := 0; i < len(rest); i++ {

		if i == 0 {
			newsenha += strings.ToTitle(rest[0])
		}

		if i != 0 {
			newsenha += rest[i]
		}
	}

	fmt.Println(newsenha)
}

func senha(number int) string {
	numberlocal := [24]string{"a", "b", "c", "d", "e", "f", "g", "h", "j", "k", "l", "m", "n", "o", "p", "q", "r", "t", "u", "v", "w", "x", "y", "z"}
	numberlocaltwo := [24]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "!", "@", "#", "$", "%", "&", "*", "A", "B", "C", "D", "M", "N", "T"}

	return numberlocal[number] + numberlocaltwo[number]

}
