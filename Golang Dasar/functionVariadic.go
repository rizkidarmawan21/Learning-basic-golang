package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}

	return total
}

func main()  {
	total := sumAll()
	fmt.Println(total)
	
	slice := []int{10,20,12}
	
	total = sumAll(slice...)
	fmt.Println(total)
}