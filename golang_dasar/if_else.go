package main

import "fmt"

func checkError(err bool) {
	if err {
		fmt.Println("Terjadi Error")
	} else {
		fmt.Println("Tidak ada error")
	}
}

func main() {
	var error bool = false
	checkError(error)

	var nilai int  = 65

	if nilai >= 70 {
		fmt.Println("Selamat anda lulus!")
	} else if nilai >= 60 && nilai < 70 {
		fmt.Println("Maaf anda belum lulus dan harus mengulang!")
	} else {
		fmt.Println("Maaf anda tidak lulus!")
	}
}