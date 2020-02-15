package main

import (
	"fmt"
	"math/rand"
	"time"
)

//「子要素は親要素より常に大きいか等しい（または常に小さいか等しい）」という制約を持つ木構造をヒープと呼ぶ
// ここでは二分ヒープ（バイナリヒープ）を配列で作成する

// ノードnの親ノードはn/2,子ノードは2n,2n+1でアクセスする

func createheap(slice []int) []int {
	// 最後のインデックス
	last := len(slice) - 1
	// 親ノードのインデックス
	parent := (last - 1) / 2

	// 親ノードのインデックスが負の値になるまで実行
	for parent >= 0 {
		// 子ノードのインデックス
		child := parent*2 + 1

		// 子ノードのインデックスが最後のインデックスより大きい場合（子ノードを持たない場合）
		// 親ノードのインデックスを一つ下げて再実行
		if child > last {
			parent = parent - 1
			continue
		}

		// 右子ノードと左子ノードで比較をとり、大きい方を子ノードに指定する
		if child != last && slice[child+1] > slice[child] {
			child = child + 1
		}

		// 子ノードのインデックスが親ノードのインデックスより大きい場合
		// 子ノードと親ノードの値を入れ替え、子ノードを新しい親とする
		// そうでなければ親ノードのインデックスを一つ下げて再実行
		if slice[parent] >= slice[child] {
			parent = parent - 1
			continue
		}
		slice[parent], slice[child] = slice[child], slice[parent]
		parent = child
	}

	return slice
}

func heapsort(slice []int) []int {
	// スライスの要素数
	length := len(slice)

	// スライスの要素が一つ以下ならスライス自身を返して終了
	if length <= 1 {
		return slice
	}

	// ヒープの作成
	slice = createheap(slice)
	if length >= 31 {
		fmt.Println("ヒープ生成：")
		for n := 0; n < length; n++ {
			fmt.Print(slice[n])
			if n == 0 || n == 2 || n == 6 || n == 14 || n == 30 {
				fmt.Println("")
			} else {
				fmt.Print(",")
			}
		}
	}

	// スライスの最初と最後を入れ替え
	slice[0], slice[length-1] = slice[length-1], slice[0]

	return append(heapsort(slice[:length-1]), slice[length-1])
}

func main() {

	rand.Seed(time.Now().Unix())
	start := time.Now()

	// 31個の0~50のランダム値をスライスに格納する
	slice := []int{}
	for n := 0; n < 31; n++ {
		slice = append(slice, rand.Intn(50))
	}

	fmt.Println("ヒープソート前：", slice)

	heapedslice := heapsort(slice)

	fmt.Println("ヒープソート後：", heapedslice)

	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}
