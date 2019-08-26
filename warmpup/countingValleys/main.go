package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int32 = 8
	var s = "UDDDUDUU"

	fmt.Println(countingValleys(n, s))
}

// Complete the countingValleys function below.
func countingValleys(n int32, s string) int32 {
	var result int32
	sArr := strings.Split(s, "")

	var seaLevel int32

	for i, v := range sArr {
		if v == "U" {
			seaLevel++
		}
		if v == "D" {
			seaLevel--
		}

		if i < len(sArr)-1 {
			if seaLevel == -1 && sArr[i+1] == "U" {
				result++
			}
		}
	}

	return result
}
