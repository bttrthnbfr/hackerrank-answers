package mb

import (
	"fmt"
	"math"
)

func minimumBribes(q []int32) error {
	var bribe int32
	for i, v := range q {
		if int(v)-(i+1) > 3 {
			fmt.Println("Too chaotic")
		}

		for j := math.Max(0, float64(v-2)); int(j) < i; j++ {
			if q[int(j)] > v {
				fmt.Println(q[int(j)] > v, q[int(j)], v)
				bribe++
			}
		}
	}

	fmt.Println(bribe)

	return nil
}
