package main

import "fmt"

func endApp()  {
	fmt.Println("Aplikasi Selesai")
}

func runApp(error bool)  {
	defer endApp()

	if error {
		panic("APLIKASI ERROR")
	}else {
		fmt.Println("Aplikasi Berjalan")
	}
}

func main()  {
	runApp(true)
}