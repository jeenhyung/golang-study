package main

import . "fmt"

func main() {

	for bottles := 99; bottles >= 0; bottles-- {
		if bottles != 1 {
			Println(bottles, " bottles of beer on the wall, ", bottles, " bottles of beer.")
			Println("Take one down, pass it around, ", bottles-1, " bottles of beer on the wall.")
		} else if bottles == 1 {
			Println(bottles, " bottle of beer on the wall, ", bottles, " bottle of beer.")
			Println("Take one down, pass it around, No more bottles of beer on the wall.")
		} else {
			Println("No more bottles of beer on the wall, No more bottles of beer.")
			Println("Go to the store and buy some more, 99 bottles of beer on the wall.")
		}
	}

}
