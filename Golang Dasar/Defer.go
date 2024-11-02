package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil")
}

func runApplication(value int) {
	defer logging() // ini akan di eksekusi ketika fungsi selesai di eksekusi/ diakhir
				    // defer ini akan selalu dieksekusi walaupun function error
	fmt.Println("run application")

	result := 10 / value
	fmt.Println("result",result)
}

func main() {
	runApplication(1)
	// runApplication(0)
}
