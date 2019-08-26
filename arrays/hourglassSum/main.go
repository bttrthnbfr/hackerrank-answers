package main

import "fmt"

func main() {
	arr := [][]int32{
		[]int32{1, 1, 1, 0, 0, 0},
		[]int32{0, 1, 0, 0, 0, 0},
		[]int32{1, 1, 1, 0, 0, 0},
		[]int32{0, 0, 2, 4, 4, 0},
		[]int32{0, 0, 0, 2, 0, 0},
		[]int32{0, 0, 1, 2, 4, 0},
	}

	fmt.Println(hourglassSum(arr))
}

// Complete the hourglassSum function below.
func hourglassSum(arr [][]int32) int32 {
	var max int32
	for y := 0; y < 4; y++ {
		for z := 0; z < 4; z++ {
			var temp int32
			for i := y + 0; i < y+3; i++ {
				for j := z + 0; j < z+3; j++ {
					if (j == z+0 && i == y+1) || (j == z+2 && i == y+1) {
						continue
					}
					temp += arr[i][j]
				}
			}
			if max < temp {
				max = temp
			}
			if y == 0 && z == 0 {
				max = temp
			}
		}
	}

	return max
}
