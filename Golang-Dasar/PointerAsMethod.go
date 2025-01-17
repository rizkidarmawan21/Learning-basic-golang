package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married(){
	man.Name = "Mr. " + man.Name
}

func main() {
	darms := Man{"Darms"}
	darms.Married()
	
	fmt.Println(darms.Name) // output: Darms
}