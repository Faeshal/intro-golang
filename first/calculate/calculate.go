package calculate

// * public function (cirinya huruf besar)
func Sum(num1 int, num2 int) int {
	return num1 + num2
}

func Salary(num1 int, num2 int) int {
	return calSalary(num1, num2)
}

// * function huruf kecil artinya private, hanya bisa dipanggil di file yg sama
func calSalary(num1 int, num2 int) int {
	rumus := (num1 + num2) * 2
	return rumus
}
