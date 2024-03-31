package main

import "fmt"

func main() {
	products := make(map[string]float32)
	products["Яблоки"] = 238.6
	products["Бананы"] = 126.3
	products["Апельсины"] = 289.78
	products["Груши"] = 198.34
	fmt.Println(products["Апельсины"])
	delete(products, "Бананы")
	fmt.Println(products)
}
