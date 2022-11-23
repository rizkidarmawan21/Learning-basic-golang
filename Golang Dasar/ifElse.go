package main

import "fmt"

func main(){
	var name = "Manadd"

	if name == "Rizki" {
		fmt.Println("Hello rizki")
	} else if name == "Mana"{
		fmt.Println("Hello Mana")
	} else {
		fmt.Println("Hello guest")
	}

	// if short statement
	if length := len(name); length > 5 {
		fmt.Println("Lebih dari 5")
	}
}