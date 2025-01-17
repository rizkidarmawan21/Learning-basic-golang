package main

import "fmt"

func main()  {

	// Pada kasus ini tidak ada yg berbeda seperti bahasa pemograman lain , sama aja 


	var a = 10
	var b = 13
	var c = a + b // mirip bahasa lain

	fmt.Println(c)


	// augmented assignment
	var  d = 10

	d = d + 10
	d += d

	fmt.Println(d)
	
	
	// unari operator
	// ++ == a = a+1
	
	var i = 7
	i++
	fmt.Println(i)



}