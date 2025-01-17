package main

import (
	"Golang-Dasar/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Darmawan")
	fmt.Println(result)
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // error karena version adalah private
	// fmt.Println(helper.sayGoodBye("Darmawan")) // error karena sayGoodBye adalah private
}