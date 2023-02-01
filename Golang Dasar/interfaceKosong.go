package main

import "fmt"

// dengan interface kosong maka return value bisa integer,string,boolean,dan sebagainya
func Ups(i int) interface{} {
	if i == 1 {
		return 1
	}else if i == 2 {
		return 2
	}else {
		return "Ups"
	}
}

func main(){
	var data interface{} = Ups(4)
	fmt.Println(data)
}