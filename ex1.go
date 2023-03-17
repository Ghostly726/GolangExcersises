package check3

import "fmt"

/*
 Create a program that takes a number as input and displays the multiplication table of that number up to 10.
*/

func Ex1(nbr int) {
	for i := 0; i <= 10; i++ {
		fmt.Println(nbr * i)
	}
}
