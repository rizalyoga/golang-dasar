package main

import "fmt"

func main() {
	// Slice adalah bentuk tipe reference daru array, bisa di bilang slice adalah referensi dari nilai array
	// hanya saja slice lebih dinamis karena panjang slice bisa berubah - ubah

	// Pembuatan slice sama dengan array hanya saja kita kosongkan nilai panjangnya
	var nomorAntrian = []int{1,2,3,4,5,6,7,8,9,10}
	for _,v := range nomorAntrian{
		fmt.Println(v)
	}

	// Pembuatan slice menggunakan make
	alpabet := make([]string,8) 
	// 8 adalah panjang array, jika nilai yang diisi hanya 4 maka sisanya akan bernilai zero value
	alpabet[0] = "a"
	alpabet[1] = "b"
	alpabet[2] = "c"
	alpabet[3] = "d"
	fmt.Println("Ini alpabet:",alpabet)
	fmt.Println("lebar alpabet:",len(alpabet))

	// Membuat slice dari array yang sudah ada
	var users = []string{"Doni","Maja","Ilham","Senja"}
	users2 :=users[0:3] // copy data dimulai dari indeks ke-0, hingga elemen sebelum indeks ke-3. 
	users3 :=users[1:3] // copy data dimulai dari indeks ke-1, hingga elemen sebelum indeks ke-3. 
	// users[:]		// ["Doni","Maja","Ilham","Senja"]	semua elemen
	// users[2:]	// ["Ilham","Senja"]	semua elemen mulai indeks ke-2
	// users[:2]	// ["Doni","Maja"]	semua elemen hingga sebelum indeks ke-2

	fmt.Println("nilai users2",users2) //"Doni", "Maja", "Ilham"
	fmt.Println("nilai users3",users3) //"Maja", "Ilham"

	// Pada golang terdapat yang namanya len dan cap
	// len adalah fungsi yang digunakan untuk mendapatkan panjang array
	// cap adalah fungsi yang digunakan untuk mendapatkan kapasitas array

	// len(users)  adalah 4
	// cap(users)  adalah 4

	// Pada users2, panjang slice adalah 3 dan kapasitas adalah 4
	fmt.Println("Panjang:",len(users2)) //3 
	fmt.Println("Kapasitas:",cap(users2)) //4

	// Contoh lain
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println("fruits",len(fruits))  // len: 4
	fmt.Println("fruits",cap(fruits))  // cap: 4

	var aFruits = fruits[0:3]
	fmt.Println("aFruits",len(aFruits)) // len: 3
	fmt.Println("aFruits",cap(aFruits)) // cap: 4

	var bFruits = fruits[1:4]
	fmt.Println("bFruits",len(bFruits)) // len: 3
	fmt.Println("bFruits",cap(bFruits)) // cap: 3, bernilai 3 karena slice dimulai dari index 1
	
	var animals = []string{"Singa","Macan","Kucing","Badak","Buaya","Kera"}
	var animals1 = animals[2:]
	fmt.Println("animals1:",len(animals1)) // len: 4
	fmt.Println("animals1:",cap(animals1)) // cap: 4

	var animals2 = animals[2:4]
	fmt.Println("animals2:",len(animals2)) // len: 2
	fmt.Println("animals2:",cap(animals2)) // cap: 4
	
	//Menambah nilai pada slice di akhir
	nomorAntrian = append(nomorAntrian, 11,12)
	for _,v := range nomorAntrian{
		fmt.Println(v)
	}

	// Append juga bisa digunakan untuk membuat array/slice baru
	var campuranData = append(fruits,animals...)
	fmt.Println("Campuran data:",campuranData)

	// Fungsi copy
	// copy(dst,src) digunakan untuk men-copy elements slice pada src (parameter ke-2), ke dst (parameter pertama).
	dst := make([]string, 3) // 3 adalah panjang dst
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple apple
	fmt.Println(src) // watermelon pinnaple apple orange
	fmt.Println(n)   // 3, 3 adalah jumlah data yang berhasil dicopy

	// Contoh lain jika dst sudah memiliki data
	dst1 := []string{"potato", "potato", "potato"}
	src1 := []string{"watermelon", "pinnaple"}
	n1 := copy(dst1, src1)

	fmt.Println(dst1) // watermelon pinnaple potato, nilai dst1 akan direplace dengan data pada src1
	fmt.Println(src1) // watermelon pinnaple
	fmt.Println(n1)   // 2

	// Slice data array dengan 3 index
	// 3 index adalah teknik slicing untuk pengaksesan elemen yang sekaligus menentukan kapasitasnya. Cara penggunaannya yaitu dengan menyisipkan angka kapasitas di belakang, seperti fruits[0:1:1]. Angka kapasitas yang diisikan tidak boleh melebihi kapasitas slice yang akan di slicing.

	var fruitss = []string{"apple", "grape", "banana"}
	var aFruitss = fruits[0:2]
	var bFruitss = fruits[0:2:2]

	fmt.Println(fruitss)      // ["apple", "grape", "banana"]
	fmt.Println(len(fruitss)) // len: 3
	fmt.Println(cap(fruitss)) // cap: 3

	fmt.Println(aFruitss)      // ["apple", "grape"]
	fmt.Println(len(aFruitss)) // len: 2
	fmt.Println(cap(aFruitss)) // cap: 3

	fmt.Println(bFruitss)      // ["apple", "grape"]
	fmt.Println(len(bFruitss)) // len: 2
	fmt.Println(cap(bFruitss)) // cap: 2
}