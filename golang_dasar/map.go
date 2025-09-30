package main

import "fmt"

func main()  {
	// Pada golang map ini adalah object di Javascript
	var userData = map[string]string{"nama":"rizal","status":"belum menikah"}

	var animalData = make(map[string]string)
	animalData["nama"] = "Singa"
	animalData["gender"] = "laki-laki"

	fmt.Println(userData)
	fmt.Println(animalData)

	// Map sama seperti slice dimana data pada map dapat di loop menggunakan range
	// Jika pada array/slice variable yang digunakan adalah index dan value, disini adalah key dan value
	for key,val:= range userData{
		fmt.Printf("\nkey %v , dengan value %v", key,val)
	}

	// Menghapus item pada map
	// untuk menghapus sebuah atau lebih item pada map, dapat menggunakan delete()
	delete(animalData, "gender")
	fmt.Println("\nnilai animal data baru:",animalData)

	// Mencari nilai value item
	// Pada map kita dapat mencari atau cek keberadaan sebuah item menggunakan 2 variable
	var value, check = userData["gender"] // var value adalah nilai data, check adalah bool true of false
	// 2 variable di atas bersifat bebas dapat dinamakan apa saja
	fmt.Println(check) // true
	fmt.Println(value) // rizal
}