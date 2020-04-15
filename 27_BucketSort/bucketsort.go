package main

import (
	"fmt"
	"math/rand"
	"time"
)

func BucketSort(slice []int) []int {
	// 出現回数
	var count []int
	count = make([]int, len(slice))
	// 各値の開始位置
	var offset []int
	offset = make([]int, len(count))
	// ソート後のスライス
	var slice2 []int
	slice2 = make([]int, len(slice))

	// スライスの値の出現回数を計算する
	for i := 0; i < len(slice); i++ {
		count[slice[i]]++
	}

	// 各値の開始位置を計算する
	offset[0] = 0
	for i := 1; i < len(count); i++ {
		offset[i] = offset[i-1] + count[i-1]
	}
	fmt.Println("count=", count)
	fmt.Println("offset=", offset)

	// ソート開始
	for i := 0; i < len(slice); i++ {
		// スライスの値をtmpに格納する
		tmp := slice[i]
		fmt.Println("i=", i, "slice[i]の値=", tmp, "値の出現回数=", count[tmp], "値の格納位置=", offset[tmp])
		// tmpの値をその値の開始位置に格納する
		slice2[offset[tmp]] = tmp
		// 次に同じ値が来た時のために格納位置をずらす
		offset[tmp]++
	}

	return slice2

}

func ShowSlice(slice []int, length int) {

	for n := 0; n < len(slice); n++ {
		fmt.Print(slice[n])
		if n%length == length-1 {
			fmt.Println("")
		} else {
			fmt.Print(",")
		}
	}

}

func main() {
	rand.Seed(time.Now().Unix())
	start := time.Now()
	// 100個の0~49のランダム値をスライスに格納する
	slice := []int{}
	for n := 0; n < 100; n++ {
		slice = append(slice, rand.Intn(50))
	}

	ShowSlice(slice, 10)

	//バケットソートを実行する
	slice_sorted := BucketSort(slice)

	//バケットソートされたスライスを再表示する
	ShowSlice(slice_sorted, 10)

	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())

}
