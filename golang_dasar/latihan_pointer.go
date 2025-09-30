package main

import "fmt"

func decrement (p *int) {
	 *p--
}

type Person struct {
	name string
	age int
}

func birthday(p *Person) {
	p.age = p.age + 1
}

func prepend(s []int) []int {
	return append([]int{100}, s... )
}

func findFirstPositive(nums []int) *int {
	var positif int 
	for _,v := range nums{
		if v > 0  {
			positif = v
			break
		}
	} 
	
	if positif == 0 {
		return  nil
		} else {
		return  &positif
	}
}

func main()  {
	// Latihan
	// Level pemula
 	// Tugas: Buat program yang memiliki variabel b dengan nilai 7. Buat pointer ke b, 
	// lalu ubah nilai melalui pointer menjadi 14. Tampilkan nilai b sebelum dan sesudah.
	var b int = 7
	var pointer = &b
	fmt.Println("Nilai b sebelum dirubah:", b)
	*pointer = 14
	fmt.Println("Nilai b sesudah dirubah:", b)

	// Tugas: Fungsi decrement menerima *int dan mengurangi nilainya satu.
	// Tampilkan hasil sebelum dan sesudah pemanggilan.
	var parameterDecrement int = 8
	fmt.Println("nilai paramter sebelum:",parameterDecrement)
	decrement(&parameterDecrement)
	fmt.Println("nilai paramter sesudah:",parameterDecrement)

	// Level menengah
	// Tugas: Buat struct Person dengan field Name string dan Age int.
	// Buat fungsi Birthday(p *Person) yang menambah usia satu tahun. Demonstrasikan dengan objek.

	var person1 Person = Person{"yoga",29}
	fmt.Println("sebelum birthday:", person1)
	birthday(&person1)
	fmt.Println("setelah birthday:", person1)

	// Tugas: Fungsi prepend menerima slice dari int dan sebuah int baru,
	// mengembalikan slice baru dengan elemen baru di depan. Jelaskan mengapa kita perlu mengembalikan slice.
	var sliceInt = []int{10,20,30,40,50}
	fmt.Println("slice sebelum append:", sliceInt)
	newSlice := prepend(sliceInt)
	fmt.Println("slice sesudah append:", newSlice)

	// Tugas: Implementasikan fungsi findFirstPositive(nums []int) *int yang mengembalikan 
	// pointer ke elemen pertama yang bernilai positif. Jika tidak ada, kembalikan nil.
	var initialNums = []int{-1,-2,0,-3,-8,-9}
	fmt.Println("Nilai asli slice:", initialNums)
	var positif *int = findFirstPositive(initialNums)
	if positif != nil {
		fmt.Println("Nilai alamat positif slice:", positif)
		}else{
		fmt.Println("Tidak ada nilai positif")
	}
}