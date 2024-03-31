package main

import (
	"fmt"
	intdiv "hw/intdiv"
	mult "hw/mult"
	remdiv "hw/remdiv"
	sub "hw/sub"
	sum "hw/sum"
)

func main() {
	var a, b int
	fmt.Scanln(&a, &b)
	fmt.Println(sum.Sum(a, b))
	fmt.Println(sub.Sub(a, b))
	fmt.Println(mult.Mult(a, b))
	fmt.Println(intdiv.Intdiv(a, b))
	fmt.Println(remdiv.Remdiv(a, b))
}
