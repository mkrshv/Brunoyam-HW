package main

import (
	"fmt"
	intdiv "hard/intdiv"
	mult "hard/mult"
	remdiv "hard/remdiv"
	sub "hard/sub"
	sum "hard/sum"
)

func main() {
	var a, b int
	fmt.Scanln(&a, &b)
	results := make(map[string]int)

	results["Сложение"] = sum.Sum(a, b)
	results["Вычитание"] = sub.Sub(a, b)
	results["Умножение"] = mult.Mult(a, b)
	results["Деление нацело"] = intdiv.Intdiv(a, b)
	results["Остаток от деления"] = remdiv.Remdiv(a, b)

	for k, v := range results {
		fmt.Printf("%s: %d\n", k, v)
	}
}
