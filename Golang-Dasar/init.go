package main

import (
	"Golang-Dasar/database"
	_"Golang-Dasar/internal"
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}