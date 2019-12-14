package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	start := time.Now()

	box := [100]int{}
	for n := 0; n < 100; n++ {
		box[n] = rand.Intn(100)
		fmt.Print(box[n])
		if n%10 == 9 {
			fmt.Println("")
		} else {
			fmt.Print(",")
		}

	}

	target := rand.Intn(100)
	fmt.Printf("検索する数値：%d\n", target)
	res := 0
	for n := 0; n < 100; n++ {
		if box[n] == target {
			res = n + 1
			fmt.Printf("%dを%d番目で発見しました。\n", target, res)
		}
	}

	if res == 0 {
		fmt.Println("検索する数値は配列の中にありませんでした。")
	}

	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}
