package main

import (
	"fmt"
	"math/rand"
	"time"
)

func InsertionSort(slice []int) []int {

	// 2番目から100番目を終点対象とする
	for n := 1; n < len(slice); n++ {
		m := n
		for m > 0 && slice[m-1] > slice[m] {
			fmt.Printf("入れ替え:%d番目（%d）,%d番目（%d）\n", m, slice[m-1], m+1, slice[m])
			slice[m-1], slice[m] = slice[m], slice[m-1]
			m--
		}
	}
	return slice
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
	// 100個の0~99のランダム値をスライスに格納する
	slice := []int{}
	for n := 0; n < 100; n++ {
		slice = append(slice, rand.Intn(100))
	}

	ShowSlice(slice, 10)

	//挿入ソートを実行する
	slice2 := InsertionSort(slice)

	//バブルソートされたスライスを再表示する
	ShowSlice(slice2, 10)

	// さらに100個の100~199のランダム値をスライスに格納する
	fmt.Println("---さらに100個追加---")
	for n := 0; n < 100; n++ {
		slice2 = append(slice2, rand.Intn(100)+100)
	}

	// スライスを表示する
	ShowSlice(slice2, 20)

	//挿入ソートを実行する
	slice3 := InsertionSort(slice2)

	//バブルソートされたスライスを再表示する
	ShowSlice(slice3, 20)

	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())

}
