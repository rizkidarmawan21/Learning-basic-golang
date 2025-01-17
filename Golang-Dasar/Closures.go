package main

import "fmt"

// Merupakan kemampuan sebuah function berinteraksi dengan data-data disekitar dalam scope yang sama
func main(){
	counter :=0
	increment := func ()  {
		// counter := 1
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
}