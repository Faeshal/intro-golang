package main

import "fmt"

func main() {

	// * 1. VARIABLE
	// typenya ada banyak, yg sering dipake: string,int,float64,bool
	// var name string = "faeshal"
	// var age int

	// type varible disingkat
	// kalau mau disingkat, ga perlu di declare typenya apa, golang udh tau sendiri
	// isSingle := true

	// contoh reassign
	// age = 20

	// fmt.Println("my name is:", name)
	// fmt.Println("age:", age)
	// fmt.Println("is single:", isSingle)

	// * CONDITIONAL (IF-ELSE etc)
	// if age > 20 {
	// 	log.Printf("umur lebih dari 20")
	// } else if age == 20 {
	// 	log.Printf("umur spesial")
	// } else {
	// 	log.Printf("masih dibawah umur")
	// }

	// looping
	// for i := 1; i < 100; i++ {
	// 	log.Printf("still learning golang")
	// }

	// text := "Golang the best language"
	// for index, letter := range text {

	// 	// * cek angka genap
	// 	if letter%2 != 0 {
	// 		fmt.Println(index)
	// 		fmt.Println("Boom, Angka Ganjil..")
	// 	}
	// 	// * cek huruf vocal
	// 	fmtString := string(letter)
	// 	if fmtString == "a" || fmtString == "i" || fmtString == "u" || fmtString == "e" || fmtString == "o" {
	// 		fmt.Println(fmtString)
	// 		fmt.Println("huruf vocal")
	// 	}

	// }

	// * ARRAY
	var income []int
	fmt.Println(income)

	var stack [3]string
	stack[0] = "JavaScript"
	stack[1] = "PHP"
	stack[2] = "Golang"
	// fmt.Println(stack)

	// cara lain , init arraynya langsung di isi, hasilnya sama kaya diatas
	stack2 := [3]string{"golang", "nodejs", "C#"}
	// fmt.Println(stack2)

	// stackLength := len(stack)
	// fmt.Println(stackLength)

	// looping array
	for _, v := range stack2 {
		fmt.Println(v)
	}

	// * SLICE
	// mirip kaya array tapi lebih dynamic
	// cocok dipake kalau ga tau pasti mau simpen berapa banyak data di dalam array
	console := []string{"ps1", "ps2"}
	for _, v := range console {
		fmt.Println(v)
	}

	// * MAP
	// dipake untuk bikin array yg keynya bisa ditentuin sendiri , ga wajib angka (0,1,2,3) bisa string
	// car := map[string]string{"echo": "golang", "express": "nodejs", "rails": "ruby"}
	// for index, v := range car {
	// 	fmt.Println("index:", index, "data:", v)
	// }

	// * SLICE of MAP
	// gabungan slice dan map , sangat dynamic. cocok buat array of object
	// students := []map[string]string{
	// 	{"name": "faeshal", "score": "A"},
	// 	{"name": "fahad", "score": "B"},
	// 	{"name": "khalid", "score": "C"},
	// }

	// for _, v := range students {
	// 	fmt.Println(v)
	// }

	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	total := 0
	for _, v := range scores {
		total = total + v
	}
	finalScore := float64(total) / float64(len(scores))
	fmt.Println("final score:", finalScore)

}
