package main

import (
	"fmt"
)

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}
func (a Customer) sayHuu() {
	fmt.Println("Huuuu from", a.Name)
}

func main() {
	// Template data untuk menggabungkan beberapa data / prorotype data
	var eko Customer

	eko.Name = "Eko Kurni"
	eko.Address = "Bandung"
	eko.Age = 30

	eko.sayHello("Joko")

	eko.sayHuu()

	fmt.Println(eko.Name)
	fmt.Println(eko.Address)
	fmt.Println(eko.Age)

	// cara membuat struct cara lain

	joko := Customer{
		Name:    "Joko",
		Address: "Indo",
		Age:     24,
	}

	fmt.Println(joko)

	budi := Customer{"Budi", "Jepang", 22}
	fmt.Println(budi)

}
