package main

import "fmt"

type Address struct {
	Kota string
	Provinsi string
	Negara string
}

func main(){
	// Pointer adalah reference atau alamat dari sebuah data
	// Secara default di Go-Lang semua variable itu di passing by value, bukan by reference
	// Artinya, jika kita mengirim sebuah variable ke dalam function, method atau variable lain,
	// sebenarnya yang dikirim adalah duplikasi value nya

	var alamat Address = Address{"Mojokerto","Jatim", "Indonesia"}
	var alamat1 Address =  alamat

	// Nilai alamat1 akan mengikuti variable alamat, dan ketika kita modifikasi salah satu nilai pada alamat1
	// modifikasi ini tidak akan berdampak pada variable alamat 

	fmt.Println(alamat) // {"Mojokerto","Jatim", "Indonesia"}
	fmt.Println(alamat1) // {"Mojokerto","Jatim", "Indonesia"}


	alamat1.Negara = "Belanda"
	fmt.Println(alamat) // {"Mojokerto","Jatim", "Indonesia"}
	fmt.Println(alamat1) // {"Mojokerto","Jatim", "Belanda"}

	// Perubahan nilai negara hanya berlaku pada variable alamat1. Jika ingin  nilai negara pada variable alamat
	// ikut berubah maka kita bisa menggunakan simbol & diawal nama variable 

	var alamat2  = &alamat
	alamat2.Negara = "Cina"
	fmt.Println(alamat) // {"Mojokerto","Jatim", "Cina"}
	fmt.Println(alamat2) // {"Mojokerto","Jatim", "Cina"}

	// Kenapa nilai alamat ikut berubah ketika item Negara di alamat2 dirubah?
	// Hal ini karena simbol (&) membuat alamat2 memiliki reference atau alamat memori yang sama dengan alamat.
	// Jika tanpa (&) maka alamat2 hanya akan mengkopy data saja dan menggunakan alamat memory yang baru.

	// Lalu bagaimana jika variable alamat2 merubah alamat memorinya?
	// Maka alamat2 akan memiliki alamat memory yang baru dan ketika modfikasi item/data dilakukan 
	// tidak akan berdampak pada variable alamat.

	alamat2 = &Address{"Jombang","Jatim","Indonesia"}
	fmt.Println(alamat) // {"Mojokerto","Jatim", "Cina"}
	fmt.Println(alamat2) // {"Jombang","Jatim","Indonesia"}

	// Pada pointer juga terdapat simbol (*) yang berfungsi untuk mengambil nilai asli / value data dari sebuah pointer.
	// Jika simbol (&) digunakan untuk mendapatkan alamat memory maka simbol (*) untuk mendapat valuenya
	// Penggunakan (&) disebut referencing dan (*) disebut dereferencing.
	
	// variabel biasa
    a := 10
    fmt.Println("Nilai a:", a)
    // buat pointer ke a
    p := &a
    fmt.Println("Alamat a:", p) //0xc000010178 ( alamat memory a )
    // mengambil nilai lewat pointer
    fmt.Println("Nilai lewat *p:", *p) // 10
    // mengubah nilai lewat pointer
    *p = 20
    fmt.Println("Setelah *p = 20, nilai a menjadi:", a) // 20



}