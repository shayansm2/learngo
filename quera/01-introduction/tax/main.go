package main

import "fmt"

func main() {
	var salary int
	fmt.Scanln(&salary)
	fmt.Println(calcTax(salary))
}

func calcTax(salary int) int {
	if salary <= 10000 {
		return int(float32(salary) * 0.05)
	}

	if salary <= 40000+10000 {
		return int(float32(salary-10000)*0.1) + int(10000*0.05)
	}

	if salary <= 50000+40000+10000 {
		return int(float32(salary-10000-40000)*0.15) + int(40000*0.1) + int(10000*0.05)
	}

	return int(float32(salary-10000-40000-50000)*0.2) + int(50000*0.15) + int(40000*0.1) + int(10000*0.05)
}
