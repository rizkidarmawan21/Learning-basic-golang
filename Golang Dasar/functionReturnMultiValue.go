package main

import "fmt"

func getFullName()(string,string) {
	return "Eko","Khannedy"
}

func main() {
	// firstName,lastName := getFullName()
	firstName,_ := getFullName() // kalau cuman ingin ambil beberapa bisa ganti _ yg tidk dibutuhkan
	fmt.Println(firstName)
}