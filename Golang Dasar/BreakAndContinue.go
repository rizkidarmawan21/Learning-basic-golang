package main

import "fmt"

func main(){
	for i := 0; i < 10; i++ {
		if i == 5 {
			break	// kalau i itu sudah sampai 5 maka berhenti		
		}
		fmt.Println("perulangan ke-",i)
	}
	fmt.Println("\n")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue    // kalau i di modulo 2 sama dgn 0 berati next atau di skip
					    // kalau genap maka skip			
		}
		fmt.Println("perulangan ke-",i)
	}
}