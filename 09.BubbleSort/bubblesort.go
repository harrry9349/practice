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
	// 100個の0~99のランダム値をスライスに格納する
	slice := []int{}
	for n := 0; n < 100; n++ {
		slice = append(slice, rand.Intn(100))
		fmt.Print(slice[n])
		if n%10 == 9 {
			fmt.Println("")
		} else {
			fmt.Print(",")
		}
	}

	//バブルソートを実行する
	// 100番目から2番目を終点対象とする
	for n := 0; n < len(slice); n++ {
		// 1番目から終点まで比較をとる
		for m := 0; m < len(slice)-(n+1); m++ {
			// 終点と現在の位置を記録する
			fmt.Printf("(n,m)=(%d番目,%d番目)    ", len(slice)-n, m+1)
			// m番目の値がm+1番目の値より大きい場合
			if slice[m] > slice[m+1] {
				// m番目の値がm+1番目の値を入れ替える
				fmt.Printf("入れ替え:%d番目（%d）,%d番目（%d）\n", m+1, slice[m], m+2, slice[m+1])
				tmp := slice[m]
				slice[m] = slice[m+1]
				slice[m+1] = tmp
			} else {
				fmt.Println("")
			}
		}
		// 終点まで比較を終えたら、終点の一つ前を新たな終点にする
	}

	//バブルソートされたスライスを再表示する
	for n := 0; n < 100; n++ {
		fmt.Print(slice[n])
		if n%10 == 9 {
			fmt.Println("")
		} else {
			fmt.Print(",")
		}
	}

	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())

	// 参考：sort関数で配列の要素をソートする
	start = time.Now()
	slice2 := []int{}
	for n := 0; n < 100; n++ {
		slice2 = append(slice2, rand.Intn(100))
	}
	fmt.Println("参考：sort関数で配列の要素をソートする")
	fmt.Print("sort_before:slice2=")
	fmt.Println(slice2)
	sort.IntSlice(slice2).Sort()
	fmt.Print("sort_after :slice2=")
	fmt.Println(slice2)

	end = time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}
