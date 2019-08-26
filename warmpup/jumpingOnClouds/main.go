package main

import "fmt"

func main() {
	c := []int32{0, 0, 1, 0, 0, 1, 0}
	fmt.Println(jumpingOnClouds(c))
}

// Complete the jumpingOnClouds function below.
func jumpingOnClouds(c []int32) int32 {
	var count int32 = -1

	for i := 0; i < len(c); i++ {
		if i < len(c)-2 && c[i+2] == 0 {
			i++
		}
		count++
	}

	return count
}
