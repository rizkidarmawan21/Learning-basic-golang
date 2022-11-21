package main

import "fmt"

func main(){
	// di golang satu variabel hanya untuk 1 tipe data saja

	var name,lang string

	name = "Rizki Darmawan"
	fmt.Println(name)

	name = "Rizki Mana"
	lang = "indo"
	fmt.Println(lang)
	fmt.Println(name)

	// contoh 2

	// JIka ingin langsung menginisialisasi datanya maka type data tidak lah wajib
	var friendName  = "Budy"

	fmt.Println(friendName)

	// pada dasarnya ketika kita membuat variabel yang langsung diinisialisasi / diisi maka tidak wajib di beri "var"
	// Maka dibasa diganti dengan := seperti contoh di bawah
	myGirlFriend  := "Mana"

	fmt.Println(myGirlFriend)

	// multiple variabel
	var (
		firstName = "Rizki first"
		lastname = "darma last"
	)

	fmt.Println(firstName)
	fmt.Println(lastname)
}