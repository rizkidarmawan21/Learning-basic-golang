package main

import "fmt"

type Filter func(string) string

// kalo ini kepanjangan tipe funcnya , maka buat type Filter sendiri
// func sayHelloWithFilter(name string,filter func(string) string) {
func sayHelloWithFilter(name string,filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello",nameFiltered)
}

func spamFilter(name string)string{
	if name == "Anjing" {
		return "..saru.."
	}else {
		return name
	}
}


func main() {
	sayHelloWithFilter("Rizki",spamFilter)
	sayHelloWithFilter("Anjing",spamFilter)
}