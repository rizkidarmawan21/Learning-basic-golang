package main

import "fmt"

func main(){
	// potongan dari array (mirip kek array) cuman bedanya ukurannya bisa dinamis
	// slice dan array itu selalu terkoneksi

	// ada 3 tipe data pointer,lenght,capacity

	var mounts = [...]string{
		"january",
		"february",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"oktober",
		"september",
		"november",
		"desember",
	}

	var slice1 = mounts[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1)) // panjang si slice
	fmt.Println(cap(slice1)) // kapasitas dari slice dari pointer sampai akhir
	
	var slice2 = mounts[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2,"Eko") // membuat slice baru dgn menambahkan data baru di akhir array
	fmt.Println(slice3)
	slice3[1] = "Bukan desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(mounts)
	
	newSlice := make([]string ,2,5) // untuk membuat slice
	newSlice[0] = " Kurniawan" 
	newSlice[1] = " Eko kur" 
	fmt.Println(newSlice)
	
	copySlice := make([]string,len(newSlice), cap(newSlice)) // untuk copy slice
	copy(copySlice,newSlice)
	fmt.Println(copySlice)

	iniArray := [5]int{1,2,3}
	iniSlice := []int{1,2,3}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}