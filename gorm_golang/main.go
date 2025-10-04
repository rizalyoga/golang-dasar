package main

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Menu struct {
	ID uint `gorm:"primaryKey;autoincrement;type:serial;column:id"`
	Nama string `gorm:"type:varchar(255);column:nama"`
	Kategori string `gorm:"type:varchar(50);column:kategori"`
	Harga int `gorm:"type:int;column:harga"`
}

// Fungsi untuk mengatur nama tabel menu, jika fungsi ini tidak mengembalikan return maka 
// secara otomatis nama table di database akan mejadi jamak yaitu menus
func (Menu) TableName() string {
	return "menu"
}

func main(){
	connURI := "postgresql://postgres:mysecretpassword@localhost:5432/toko?sslmode=disable"
	db, err := gorm.Open(postgres.Open(connURI),&gorm.Config{})
	if err != nil {
		fmt.Printf("Gagal menghubungkan database: %v\n", err)
		os.Exit(1)
	}

	sqlDB,_  := db.DB()
	// db.DB() di Gorm memberikan nilai sql.DB yang bisa kita gunakan untuk memanggil fungsi milik interface sql.DB.
	defer sqlDB.Close()

	fmt.Println("Database berhasil terhubung...")
	
	// AutoMigrate berguna untuk membuat table di database
	db.AutoMigrate(&Menu{})
	fmt.Println("Tabel berhasil dibuat.")

	// Create new data
	newMenu := Menu{Nama: "Kopi hitam", Kategori: "Minuman", Harga: 2500}

	result := db.Create(&newMenu)
	if result.Error != nil {
		fmt.Printf("Gagal menambahkan menu %v\n", result.Error)
		os.Exit(1)	
	}

	fmt.Println("Menu baru berhasil ditambahkan...")

	// Get data by ID
	var menu Menu

	getDataByIdResult := db.First(&menu, 1)
	// db.First menampung 2 parameter, pointer dari variable penampung dan primary key data yg ingin dicari
	if getDataByIdResult.Error != nil {
		fmt.Printf("Gagal mengambil data %v \n",getDataByIdResult.Error)
		os.Exit(1)
	}

	fmt.Println(menu)

	// Get all data
	var menus []Menu
	getAllDataResult := db.Find(&menus)
	if getAllDataResult.Error != nil {
		fmt.Printf("Gagal mengambil data %v \n",getAllDataResult.Error)
		os.Exit(1)
	}

	fmt.Println(menus)

	// Get some data
	var someMenus []Menu
	// parameter ke 2 dari Find adalah id dari data mana saja yang ingin diambil
	getSomeMusResult := db.Find(&someMenus, []uint{1,5})
	if getSomeMusResult.Error != nil {
		fmt.Printf("Gagal mengambil data %v \n",getSomeMusResult.Error)
		os.Exit(1)
	}

	fmt.Println(someMenus)

	// Get data dengan kondisi WHERE dan NOT
	var listMenuByCategory []Menu
	// get data menggunakan where, maka yang tampil adalah data yang sesaui dengan kondisi where.
	// res = db.Where(map[string]interface{}{"kategori":"Minuman"}).Find(&listMenuByCategory)
	// get data menggunakan not, maka yang tampil adalah data yang tidak sesaui dengan kondisi.
	// contoh disini jika kondisi "kategori":"Minuman" maka yang diambil adalah data yang bukan minuman.
	res := db.Not(map[string]interface{}{"kategori":"Minuman"}).Find(&listMenuByCategory)
	if res.Error != nil {
		fmt.Printf("Gagal mengambil data %v \n",res.Error)
		os.Exit(1)
	}

	fmt.Println(listMenuByCategory)

	// Update data
	resultUpdate := db.Model(Menu{ID: 3}).Updates(Menu{Nama: "Bakmi goreng", Kategori: "Makanan", Harga: 6000})
	if resultUpdate.Error != nil {
		fmt.Printf("Gagal memperbarui data %v \n",resultUpdate.Error)
		os.Exit(1)
	}

	fmt.Println("Data berhasil di perbarui...")

	// Detele data
	resultDelete := db.Delete(&Menu{ID: 1})
	if resultDelete.Error != nil {
		fmt.Printf("Gagal Menghapus data %v \n",res.Error)
		os.Exit(1)
	}

	fmt.Println("Data berhasil di hapus...")
}
