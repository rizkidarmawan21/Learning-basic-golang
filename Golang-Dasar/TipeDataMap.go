package main

import "fmt"

func main(){
		person := map[string]string{
			"name" : "Rizki",
			"address" : "Semarang",
		}

		fmt.Println(person["name"])
		fmt.Println(person["address"])
		
		var book map[string]string = make(map[string]string)
		book["title"] = "Belajar Golang"
		book["author"] = "Eko"
		book["ups"] = "delete"
		fmt.Println(book)
		
		delete(book,"ups")
		fmt.Println(book)
		fmt.Println(len(book))

}