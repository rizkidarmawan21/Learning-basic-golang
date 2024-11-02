package main

import "fmt"

// maksudnya function bisa dijadikan sebagai value
func getGoodBye(name string) string {
	return "Good bye " + name
}

func main() {
	goodbye := getGoodBye
	fmt.Println(goodbye("RIzki"))
}
