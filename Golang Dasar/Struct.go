package main

import (
	"fmt"
)

type Customer struct {
	Name, Address string
	Age           int
}

// ini adalah method bukan func / function yg menempel pada struct
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

// ini adalah method bukan func
func (a Customer) sayHuu() {
	fmt.Println("Huuuu from", a.Name)
}

func main() {
	// Template data untuk menggabungkan beberapa data / prototype data
	var eko Customer
	fmt.Println("kosongan",eko)

	// ini kalay mau set valuenya
	eko.Name = "Eko Kurni"
	eko.Address = "Bandung"
	eko.Age = 30

	eko.sayHello("Joko")

	eko.sayHuu()

	// kalau mau ngambil valuenya
	fmt.Println(eko)
	fmt.Println(eko.Name)
	fmt.Println(eko.Address)
	fmt.Println(eko.Age)

	// cara membuat struct cara lain atau struck literals
	joko := Customer{
		Name:    "Joko",
		Address: "Indo",
		Age:     24,
	}

	fmt.Println(joko)

	budi := Customer{"Budi", "Jepang", 22}
	fmt.Println(budi)

}
