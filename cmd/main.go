package main

import (
	"fmt"

	"goshift.com/src/goshift.com/shifter"
)

func main() {
	arr := []int{5, 4, 3, 2, 1}
	target := []int{2, 1, 5, 4, 3}

	result := shifter.NumberOfShift(arr, target)
	fmt.Println(result)
}
