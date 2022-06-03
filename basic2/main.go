package main

import (
	"fmt"
)

type User struct {
	Id   int
	Name string
	Age  int
}

type Group struct {
	Id    int
	Name  string
	Users []User
}

type Laptop struct {
	Name   string
	Status bool
}

// ini method
func (user User) display() string {
	return fmt.Sprintf(user.Name)
}

func main() {

	// * FUNCTION
	// function standard
	myP2p := Taxcalc(210)
	fmt.Println("Passive Income:", myP2p)

	// function 2 balikan
	luas, keliling := fieldCalc(100, 85)
	fmt.Println(luas)
	fmt.Println(keliling)

	// * Struct
	user := User{923, "ziyad", 27}
	// user2 := User{782, "omar", 26}
	// groupUser := []User{user, user2}

	fmt.Println(user)
	// fmt.Println("=============")
	// group := Group{29, "Engineer", groupUser}
	// fmt.Println(group)

	// * METHOD
	username := user.display()
	fmt.Println(username)

	// * POINTER
	// numberA := 5
	// numberB := &numberA // utk simpan pointer pake tanda & (reference artinya)

	// fmt.Println(numberA)
	// fmt.Println(*numberB) // di kasih * utk convert alamat pointer ke nilai asli
	// contoh real ubah nilai variabel
	nilai := 10
	ChangeValue(&nilai, 30)
	fmt.Println(nilai)

	// contoh lain , ubah struct pake
	dataLaptop := Laptop{"macbook pro", true}
	fmt.Println("dataLaptop lama:", dataLaptop)
	ChangeLaptop(&dataLaptop)
	fmt.Println("status laptop now:", dataLaptop.Status)

}

func ChangeValue(old *int, new int) {
	*old = new
}

func ChangeLaptop(laptop *Laptop) {
	laptop.Status = false
}

func Taxcalc(p2p int) float64 {
	taxP2p := (float64(p2p) * 15) / float64(100)
	incomeP2p := float64(p2p) - float64(taxP2p)
	return incomeP2p
}

func fieldCalc(p int, l int) (int, int) {
	luas := p * l
	keliling := 2 * (p + l)
	return luas, keliling
}
