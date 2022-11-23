package main

import "fmt"

func getFullName()(string,string) {
	return "Eko","Khannedy"
}

func main() {
	// firstName,lastName := getFullName()
	firstName,_ := getFullName()
	fmt.Println(firstName)
}