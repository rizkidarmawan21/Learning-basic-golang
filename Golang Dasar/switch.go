package main

import (
	"fmt"
)

func main(){
	name := "Jokoddsd"


	switch name {
	case "Rizki":
		fmt.Println("Rizki")
	
	case "Joko":
		fmt.Println("Joko")
		
	default:
		fmt.Println("Kenalan dong")
	}

	// short statement

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Nama terlalu panjang")
	// case false:
	// 	fmt.Println("Nama sudah benar")
	// }

	// swicth di golang bisa tanpa kondisi
	length := len(name)
	switch {
	case length > 5:
		fmt.Println("Nama terlalu panjang")
	case length > 10:
		fmt.Println("Nama panjang banget")
	default:
		fmt.Println("Nama sudah benar")
	}
}