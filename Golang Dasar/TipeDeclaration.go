package main

import "fmt"

func main()  {
	// Merupakan kemampuan membuat ulang tipe data baru dari yg sudah ada
	type NoKTP string
	type Married bool

	var ktpRizki NoKTP = "24253452345324"
	var marriedStatus Married = true
	fmt.Println(ktpRizki)
	fmt.Println(marriedStatus)


}