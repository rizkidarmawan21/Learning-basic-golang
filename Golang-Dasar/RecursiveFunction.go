package main

import "fmt"

// Function yang memanggil dirinya sendiri

// tanpa recursive
func factorialLoop(value int) int {
	result := 1
	for i :=value; i > 0; i-- {
		result *=i
	}
	return result
}


// dengan recursive 
func factorialRecursive(value int) int  {
	if value == 1 {
		return 1
	}else {
		fmt.Println("asas",value)
		fmt.Println("asas -> ",value -1)
		fmt.Println("asas ->-> ",value * factorialRecursive(value-1))
		return value * factorialRecursive(value-1)
	}
}

func main(){
	loop := factorialLoop(5)
	fmt.Println(loop)
	fmt.Println(5*4*3*2*1)

	recursive := factorialRecursive(5)
	fmt.Println(recursive)

}