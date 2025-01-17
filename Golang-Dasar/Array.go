package main

import "fmt"

func main()  {
	var names [3]string


	// karna kita set arraynya cuman 3, maka cuman bisa 3 aja
	names[0] ="Rizki"
	names[1] = "Dar"
	names[2] = "mawan"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names) // semua

	var values = [3]int{
		90,80,95, // kita bisa kasih default valuenya
	}

	var values2 = [...]int{
		90,80,95,100,38, // kalau kasih ... maka akan menyesuaikan dari panjang arraynya, dan ini cuman bisa ketika langsung di deklarasikan
	}

	fmt.Println(values)
	fmt.Println(values2)
	// len untuk menghitung panjang si arraynya
	fmt.Println(len(values))
	fmt.Println(len(values2)) 
}