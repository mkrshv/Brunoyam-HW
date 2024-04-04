package main

import "fmt"

func main() {
	testSlice := []int{1, 2, 3}
	testSlice = Append(testSlice, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(testSlice, len(testSlice), cap(testSlice))
}

func Append(slice []int, elems ...int) []int {
	var result []int
	rLen := len(slice) + len(elems)

	if rLen <= cap(slice) {
		result = slice[:rLen]
	} else {
		rCap := len(slice)
		fmt.Println(rCap, len(slice))
		for capacity := rCap; capacity <= rLen; {
			rCap = capacity
			capacity *= 2
			rCap = capacity
		}
		result = make([]int, rLen, rCap)
		copy(result, slice)
	}

	for i := len(slice); i < rLen; i++ {
		result[i] = elems[i-len(slice)]
	}

	return result
}
