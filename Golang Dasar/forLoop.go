package main

import "fmt"

func main()  {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("ke-",counter)
	// 	counter++
	// }
	
	for counter := 1; counter <= 10; counter++ {
			fmt.Println("ke-",counter)
	}

	// loop array.slice,map

	slice := []string{"Rizki","Darmawan","Arina","Mana"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	//for range

	for index, name := range slice {
		fmt.Println("index ",index, "=",name)
	}

	// kalo index gk mau di pake kasih ada _ (underscore) 
	for _, name := range slice {
		fmt.Println("index =",name)
	}

	person := make(map[string]string)
	person["name"] = "Eko"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key ,"=",value)
	}
}