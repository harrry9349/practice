package main

import (
	"fmt"
	"math/rand"
	"time"
)

func binary_search(box [100]int, target int, low int, high int) int {
	var mid int
	if high < low {
		return -1
	} else {
		mid = (low + high) / 2
		if box[mid] > target {
			fmt.Printf("%dは中央値%dより小さいので、%dから%dまでの範囲で再帰します\n", target, mid, low, mid-1)
			return binary_search(box, target, low, mid-1)
		} else if box[mid] < target {
			fmt.Printf("%dは中央値%dより大きいので、%dから%dまでの範囲で再帰します\n", target, mid, mid+1, high)
			return binary_search(box, target, mid+1, high)
		} else {
			fmt.Printf("%dは中央値%dと等しいので、バイナリサーチを終了します\n", target, mid)
			return mid
		}
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	start := time.Now()

	box := [100]int{}
	for n := 0; n < 100; n++ {
		box[n] = n
		fmt.Print(box[n])
		if n%10 == 9 {
			fmt.Println("")
		} else {
			fmt.Print(",")
		}
	}

	//target := rand.Intn(100)
	//fmt.Printf("検索する数値：%d\n", target)

	res := -1
	for n := 0; n < 100; n++ {
		target := n
		fmt.Printf("検索する数値：%d\n", target)
		res = binary_search(box, target, 0, 99)
		if res == -1 {
			fmt.Println("検索する数値は配列の中にありませんでした。")
		} else {
			fmt.Printf("%dを%d番目で発見しました。\n", target, res+1)
		}
	}

	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}
