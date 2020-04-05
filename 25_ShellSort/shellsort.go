package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ShellSort(slice []int) []int {
	// 適当な間隔 h を開ける
	// 間隔 h をあけて取り出したデータ列に挿入ソートを適用する
	// 間隔 h を狭めて挿入ソートを繰り返す
	// h=1 になったら最後に挿入ソートを適用して終了

	// 間隔を決定する
	dist := 1

	// h_(i+1) = 3* h_i + 1 になるように間隔を選んでいくと計算量が最適になる
	for dist <= len(slice) {
		dist = 3*dist + 1
	}

	dist = (dist - 1) / 3

	for dist > 1 {
		for i := 0; i < dist; i++ {
			// 間隔ごとのグループを作成する
			fmt.Println("--------------------")
			fmt.Println("間隔：", dist)
			group := make([]int, len(slice)/dist+1)
			cnt := 0
			for j := 0; j < len(slice)/dist+1; j++ {
				if i+dist*j < len(slice) {
					fmt.Println("対象番号:", i+dist*j, "対象値:", slice[i+dist*j], "カウント数：", cnt)
					group[j] = slice[i+dist*j]
					cnt++
				}
			}

			// 間隔をあけて取り出したデータ列に挿入ソートを適用する
			fmt.Println("挿入ソート前：", group)

			group_sorted := InsertionSort(group[:cnt])

			fmt.Println("挿入ソート後：", group_sorted)

			// 挿入ソート後のグループをスライスに代入
			for j := 0; j < len(group_sorted); j++ {
				if i+dist*j < len(slice) {
					fmt.Println("対象番号:", i+dist*j, "変更前:", slice[i+dist*j], "変更後", group_sorted[j])
					slice[i+dist*j] = group_sorted[j]
				}
			}
		}
		fmt.Println("--------------------")
		// 間隔を狭めて挿入ソートを繰り返す
		dist = (dist - 1) / 3
	}
	// h=1 になったら最後に挿入ソートを適用して終了
	slice = InsertionSort(slice)
	return slice

}

func InsertionSort(slice []int) []int {

	// 2番目から最終番を終点対象とする
	// 取り出した値を「整列済み配列」の適切な位置に挿入する
	// 繰り返し全て「整列済み配列」に挿入し終われば完了
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

	//シェルソートを実行する
	slice_sorted := ShellSort(slice)

	//シェルソートされたスライスを再表示する
	ShowSlice(slice_sorted, 10)

	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())

}
