package main

import "fmt"

func main() {
	a := []int32{1, 2, 3, 4, 5}
	var d int32 = 4

	fmt.Println(rotLeft(a, d))
}

// Complete the rotLeft function below.
func rotLeft(a []int32, d int32) []int32 {
	mod := d % int32(len(a))
	if mod == 0 {
		return a
	}
	newArr := make([]int32, int32(len(a)))

	for i, v := range a {
		subtract := int32(i) - mod
		if subtract < 0 {
			subtract = subtract + int32(len(a))
		}
		newArr[subtract] = v
	}

	return newArr
}
