package main

import "fmt"

// Defer adalah sebuah fungsi yang dapat digunakan untuk memanggail fungsi lain
// yang ingin di eksekusi di akhir.

func logging() {
	fmt.Println("Logging: APP selesai dijalankan.")
}

func runApplication() {
	defer logging() // fungsi logging ini akan dijalankan di akhir, setelah runApplication dijalankan semua
	fmt.Println("APP untuk defer dijalankan...")

	// kenapa tidak panggil fungsi logging di baris akhir fungsi?
	// hal ini karena jika dipanggil di baris akhir dan ternyata terjadi error di baris tengah fungsi,
	// maka app akan panic / berhenti, jadi fungsi logging tidak akan sempat dijalankan.
	// fungsi defer ini membuat fungsi yang dipanggil (logging) akan dieksekusi di akhir meskipun 
	// nantinya terdapat error di tengah operasi.

	// logging()
}

// Panic adalah sebuah fungsi yang digunakan untuk menghentikan program jika nanti terjadi 
// error di tengah program dijalankan.

func endApp(){
	fmt.Println("End Program...")
}

func runApp(err bool){
	defer endApp() // fungsi endApp akan tetap dijalankan meskipun terjadi panic, 
	// hal ini karena fungsi endApp dipanggil menggunakan defer

	fmt.Println("\nProgram untuk panic dijalankan...")
	if(err){
		panic("Ups terjadi error...")
	}

	fmt.Println("Program sukses tanpa ada error")
}

// Recover adalah sebauh fugnsi yang digunakan untuk menangkap error di Panic.
// Di golang jika terjadi panic maka semua program akan langsung dihentikan, agar program tetap berjalan 
// perlu menggunakan revocer. Recover harus dijalankan fungsi yang dipanggil dever.

func programLogging(){
	fmt.Println("Program berhenti...")
	message:=recover()
	fmt.Println("Program mengalami error dengan pesan: ",message)
}

func startProgram(err bool){
	defer programLogging()

	fmt.Println("\nProgram untuk recover berjalan...")

	if err{
		panic("Upps, terjadi error di line 55.")
	}
}

func main()  {
	// defert
	runApplication()

	// panic
	runApp(false)

	// recover
	startProgram(true)
	fmt.Println("Recover berjalan...")
}