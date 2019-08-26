package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "aab"
	var n int64 = 882787

	fmt.Println(repeatedString(s, n))
}

func repeatedString(s string, n int64) int64 {
	var result int64
	sArr := strings.Split(s, "")

	//count total a in string
	var countAinArr int64
	for _, v := range sArr {
		if v == "a" {
			countAinArr++
		}
	}
	repeated := n / int64(len(sArr))
	result = countAinArr * repeated

	sis := n - (repeated * int64(len(sArr)))
	if sis != 0 {
		for i := 0; int64(i) < sis; i++ {
			if sArr[i] == "a" {
				result++
			}
		}
	}

	// fmt.Println(countAinArr)

	// if len(sArr) == 1 && sArr[0] == "a" {
	// 	return n
	// }

	// var count int64
	// for {
	// 	for _, v := range sArr {
	// 		if count == n {
	// 			break
	// 		}

	// 		if v == "a" {
	// 			result++
	// 		}

	// 		count++
	// 	}
	// 	if count == n {
	// 		break
	// 	}
	// }

	return result
}
