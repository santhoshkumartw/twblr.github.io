package control_structures

import "fmt"

func fizzBuzz(i int32) string {
	word := ""
	if i%3 == 0 {
		word += "Fizz"
	}
	if i%5 == 0 {
	 	word += "Buzz"
	 }
	 if i%3 !=0 && i%5 !=0 {
	 	return fmt.Sprintf("%d", i)
	 }
	return word
}
