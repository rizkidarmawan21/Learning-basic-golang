package main

import "fmt"

func main()  {
	var names [3]string

	names[0] ="Rizki"
	names[1] = "Dar"
	names[2] = "mawan"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,80,95,
	}

	fmt.Println(values)
	// len untuk menghitung panjang si arraynya
	fmt.Println(len(values))
}