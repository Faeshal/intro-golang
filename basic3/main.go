package main

import "fmt"

// * interface
type BangunDatar interface {
	HitungLuas() int
}

// * struct
type Persegi struct {
	Sisi int
}

type PersegiPanjang struct {
	Panjang int
	Lebar   int
}

// * method
func (obj Persegi) HitungLuas() int {
	return obj.Sisi * obj.Sisi
}

func (obj PersegiPanjang) HitungLuas() int {
	return obj.Panjang * obj.Lebar
}

func main() {
	persegi := Persegi{5}
	luasPersegi := seberapaLuas(persegi)

	persegiPanjang := PersegiPanjang{2, 3}
	luasPersegiPanjang := seberapaLuas(persegiPanjang)

	// * INTERFACE
	fmt.Println("luas persegi:", luasPersegi)
	fmt.Println("luas persegi panjang:", luasPersegiPanjang)
}

// * function
func seberapaLuas(val BangunDatar) int {
	return val.HitungLuas()
}
