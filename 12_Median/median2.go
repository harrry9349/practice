package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	start := time.Now()

	slice := []int{}

	for n := 0; n < rand.Intn(99)+2; n++ {
		slice = append(slice, rand.Intn(100))
	}

	sort.IntSlice(slice).Sort()

	if len(slice)%2 == 1 {
		fmt.Println(slice)
		fmt.Printf("配列の要素が奇数のため、（要素の長さ/2を切り捨てた値）番目にある値を取得します。(slice[%d]=%d)\n", len(slice)/2, slice[len(slice)/2])
		fmt.Printf("中央値は%dです。", slice[len(slice)/2])
	} else {
		fmt.Println(slice)
		fmt.Printf("配列の要素が偶数のため、中央に最も近い２つの値の平均値を取ります。(slice[%d]=%d,slice[%d]=%d)\n", len(slice)/2-1, slice[len(slice)/2-1], len(slice)/2, slice[len(slice)/2])

		fmt.Printf("中央値は%gです。", float64((slice[len(slice)/2-1]+slice[len(slice)/2])/2))
	}

	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}
