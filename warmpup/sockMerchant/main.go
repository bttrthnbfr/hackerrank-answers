package main

func main() {
	var n int32 = 9
	ar := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}

	sockMerchant(n, ar)
}

func sockMerchant(n int32, ar []int32) int32 {
	groupColor := make(map[int32]int32)

	for _, v := range ar {
		if _, ok := groupColor[v]; ok {
			groupColor[v]++
		} else {
			groupColor[v] = 1
		}
	}

	var result int32

	for _, v := range groupColor {
		result += v / 2
	}

	return result
}
