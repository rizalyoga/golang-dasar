package main

import "fmt"

func main() {
	// Tipe data array	
	// Tipe data array pada golang terdapat 2 bentuk, yaitu yang panjang nya diketahui dan tidak
	// Jika kita sudah yakin dengan panjang array maka kita bisa menuliskan panjangnya secara langsung saat inisialisasi
	// Jika tidak maka kita gunakan titik 3 [...]
	// Catatan array pada golang panjangnya tidak bisa dirubah

	var angka = [10]int{1,2,3,4,5,6,7,8,9,0}
	for i:=0; i < len(angka); i++ {
		fmt.Println("index:",i)
		fmt.Println("nilai:",angka[i])
	}

	var huruf = [...]string{"a","b","c","d","e"}
	for i,n := range huruf {
		fmt.Printf("\nIndex ke %v adalah huruf %v", i,n)
	}
	
	// Jika kita ingin looping dan tidak ingin menggunakan index atau value maka cukup gunakan _ untuk
	// mengabaikan nilai tersebut. 
	// Pada contoh di bawah kita tidak ingin menggunakan index, jadi kita gunakan _ pada bagian inisialisasi index
	for _,n := range angka {
		fmt.Printf("\nAngka = %v", n)
	}
}