package main

import (
	"fmt"
	"math/rand"
	"time"
)

func merge(left []int, right []int) []int {
	res := []int{}
	for len(left) > 0 && len(right) > 0 {
		var x int
		fmt.Printf("入れ替え:right[0]=%d,left[0]=%d\n", right[0], left[0])
		if right[0] > left[0] {
			x, left = left[0], left[1:]
			fmt.Println("入れ替え:x=", x, "left=", left)
		} else {
			x, right = right[0], right[1:]
			fmt.Println("入れ替え:x=", x, "right=", right)
		}

		res = append(res, x)
		fmt.Println("追加:x=", x, "res=", res)
	}

	res = append(res, left...)
	fmt.Println("左側追加:res=", res)
	res = append(res, right...)
	fmt.Println("右側追加:res=", res)
	return res
}

func sort(left []int, right []int) []int {
	if len(left) > 1 {
		lef, rig := split(left)
		left = sort(lef, rig)

	}

	if len(right) > 1 {
		lef, rig := split(right)
		right = sort(lef, rig)
	}

	fmt.Println("---マージ開始---")
	res := merge(left, right)
	fmt.Println("---マージ終了---")
	return res

}

func split(slice []int) (left, right []int) {

	left = slice[:len(slice)/2]
	right = slice[len(slice)/2:]
	fmt.Printf("slice():左側の長さ=%d,右側の長さ=%d\n", len(left), len(right))
	return

}

func MergeSort(slice []int) (res []int) {
	left, right := split(slice)
	res = sort(left, right)
	return
}

func main() {
	rand.Seed(time.Now().Unix())
	start := time.Now()
	// 100個の0~99の整数をランダムに格納する
	slice := []int{}
	for n := 0; n < 100; n++ {
		slice = append(slice, rand.Intn(100))
	}

	//スライスを表示する
	fmt.Println(slice)

	//マージソートを実行する
	slice2 := MergeSort(slice)

	//マージソートされたスライスを表示する
	fmt.Println(slice2)

	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())

}
