package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("run application")

	result := 10 / value
	fmt.Println("result",result)
}

func main() {
	runApplication(0)
}
