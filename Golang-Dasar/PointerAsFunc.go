package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// ini tidak akan berubah karena address di dalam fungsi ini adalah pass by value
func ChangeCountryToIndo(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{}
	ChangeCountryToIndo(&address)
	fmt.Println(address) // output: {  }
}
