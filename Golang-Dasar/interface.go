package main

import "fmt"

type HasName interface {
	GetName() string // di dalam ini tentukan method2 apa aja yg dimiliki HasBane
	GetAlamat() string
}

func sayHello(hasName HasName) {
	fmt.Println("hello :", hasName.GetName() , "\nAlamat :",hasName.GetAlamat())
}

type Person struct {
	Name string
	Alamat string
}

func (person Person) GetName() string {
	return person.Name
}
func (person Person) GetAlamat() string {
	return person.Alamat
}

type Animal struct {
	Name string
	Alamat string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func (animal Animal) GetAlamat() string {
	return animal.Alamat
}

func main() {
	// Tipe data abstract mirip constract
	var eko Person
	eko.Name = "Eko"
	eko.Alamat = "Semarang"

	sayHello(eko)
	// cat := Animal{
	// 	Name:"Kucing orange",
	// }

	// sayHello(cat)

	cat := Animal{
		Name: "Cat",
		Alamat: "Bandung",
	}

	sayHello(cat)
}
