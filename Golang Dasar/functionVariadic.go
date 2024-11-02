package main

import "fmt"

// function yang menerima jumlah argumen yang tak terbatas
// Function sumAll didefinisikan sebagai variadik function karena menggunakan parameter numbers ...int. 
// Ini artinya function ini bisa menerima sejumlah argumen bertipe integer yang tidak ditentukan jumlahnya.
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
	total2 := sumAll(10,20,2)
	fmt.Println(total2)
	
	slice := []int{10,20,12}
	
	total = sumAll(slice...) // ditambah ... auto diconvert jadi argument
	fmt.Println(total)
}