package main

import "fmt"
	 
type RumusBangunPersegi struct {
	panjang int // jika awal nama key bukan kapital maka key ini tidak akan ikut di export
	lebar int
}

func (p RumusBangunPersegi) cariLuasPersegi() int {
	return p.panjang * p.lebar
}

// Interface
type BangunRuang interface {
	luas() float64
}

func main() {
	// Inisialisasi variabel
	var name string = "yoga"
	var age int = 29
	fmt.Println(name,age)
	
	// Penulisan di bawah ini bersifat scope artinya hanya bisa diakses dari dalam fungsi dimana dia
	// di inisialisasi
	status:="Lajang"
	fmt.Println(status)
	
	// Array
	var arrayOfString = [4]string{"denis","mojang","Tedy","Sonic"}
	var arrayOfNumber = [4]int{1,2,3,4}
	var arrayOfMix = [...]any{"yoga",29,"slepi",2400}
	fmt.Println(arrayOfString,arrayOfNumber, arrayOfMix)
	fmt.Printf("dartar array string %v \ndartar array number %v \ndartar array campuran %v",arrayOfString,arrayOfNumber, arrayOfMix)

	// Map
	// create map manual
	var dataUser = map[string]string{"nama":"yoga","umur":"29 tahun","alamat":"slepi"}
	fmt.Println(dataUser["nama"])
	fmt.Println(dataUser["umur"])
	fmt.Println(dataUser["alamat"])
	
	// create map with make
	var dataUser2 = make(map[string]int)
	dataUser2["satu"] = 1
	dataUser2["dua"] = 2
	dataUser2["tiga"] = 3

	fmt.Println(dataUser2["satu"])
	fmt.Println(dataUser2["dua"])
	fmt.Println(dataUser2["tiga"])

	// Slice 
	var sliceInt = []int{0,2,5,1,10}
	fmt.Printf("Halis %v",sliceInt[3]) // output :1

	// Looping
	for i:=0; i <= len(sliceInt); i++ {
		fmt.Printf("index %v \n",i)
	}

	for _,v := range sliceInt{
		fmt.Println(v)
	}

	// Struct
	var persegi = RumusBangunPersegi{panjang: 20,lebar: 20}
	fmt.Println(persegi)

	// Membuat variable dari struct
	type Mahasiswa struct {
		Nama string
		Jurusan string
		Angkatan int
		Aktif bool
	}

	var mahasiswaData Mahasiswa
	mahasiswaData.Nama = "Rizal"
	mahasiswaData.Angkatan = 2016
	mahasiswaData.Jurusan = "Informatika"
	mahasiswaData.Aktif = true
	fmt.Println(mahasiswaData)
	
	// Membuat variable secara langsung (struct literal)
	mhs1 := Mahasiswa {Nama: "Doni", Jurusan: "Mesin", Angkatan: 2016, Aktif: false}
	fmt.Println(mhs1)

	// Akses method struct
	var luasPersegi = RumusBangunPersegi{panjang:15, lebar:15}
	var hasil = luasPersegi.cariLuasPersegi()
	fmt.Println(hasil)
}

