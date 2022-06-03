package main

import (
	"first/calculate"
	"fmt"
)

func main() {
	fmt.Println("hallo dunia")
	result := calculate.Sum(5, 2)
	fmt.Println(result)

	takeHomePay := calculate.Salary(3, 5)
	fmt.Println("your take home pay salary is :", takeHomePay)
}
