package main

import "fmt"

func sayHello(firstName string,lastname string){
	fmt.Println("Hello,",firstName,lastname)
}

func main(){
	sayHello("Rizki","Darmawan")
}