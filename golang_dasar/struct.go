package main

import "fmt"

// Struct
// Struct mirip dengan type di typescript
// Struct biasa di deklarasi di luar func main, jika nama strct diawali dengan kapital maka struct dapat diakses
// di berbagai package.

type RumusBangunPersegiPanjang struct {
	Panjang int
	Lebar int
}

// Pada golang fungsi struct biasa disebut method
// penulisan methode mirip dengan fungsi hanya saja terdapat receiver sebelum nama func
func (rbpp RumusBangunPersegiPanjang) LuasPersegiPanjang() int {
	return rbpp.Lebar * rbpp.Panjang
}

func main() {
	var bangunPersegiPanjang1 = RumusBangunPersegiPanjang{Panjang: 20, Lebar: 40}
	fmt.Println(bangunPersegiPanjang1)
	
	bangunPersegiPanjang2 := RumusBangunPersegiPanjang{}
	bangunPersegiPanjang2.Panjang = 40
	bangunPersegiPanjang2.Lebar = 80
	fmt.Println(bangunPersegiPanjang2)

	type User struct {
		Nama string
		Umur int
	}

	var User1 User
	User1.Nama = "Rizal"
	User1.Umur = 30
	fmt.Printf("Nama saya %v, Umur %v.", User1.Nama,User1.Umur)

	// Implementasi method
	var CariLuasPersegiPanjang1 = RumusBangunPersegiPanjang{Panjang: 5, Lebar: 20}
	var CariLuasPersegiPanjang2 = RumusBangunPersegiPanjang{Panjang: 10, Lebar: 100}
	fmt.Printf("\nLuas dari persegi panjang 1 adalah %v", CariLuasPersegiPanjang1.LuasPersegiPanjang())
	fmt.Printf("\nLuas dari persegi panjang 2 adalah %v", CariLuasPersegiPanjang2.LuasPersegiPanjang())
}