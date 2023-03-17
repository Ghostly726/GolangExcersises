package check3

//random password generator with a specified charset and length

import (
	"math/rand"

	"github.com/01-edu/z01"
)

func Ex3(length int, charset []rune) {
	for i := 0; i < length; i++ {
		z01.PrintRune(charset[rand.Intn(len(charset))])
	}
}
