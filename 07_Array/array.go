package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func delete(slice []int, target int) []int {
	// targetの値を除いたスライスを再作成する
	res := []int{}
	for _, n := range slice {
		if n != target {
			res = append(res, n)
		}
	}
	return res
}

func main() {
	rand.Seed(time.Now().Unix())
	start := time.Now()

	//定義
	// array:要素と型を定義した配列
	// printlnで0が100個表示される
	array := [100]int{}
	fmt.Print("array=")
	fmt.Println(array)
	// slice:型を定義した配列、要素数を書かない
	// printlnで空のsliceが表示される
	slice := []int{}
	fmt.Print("slice=")
	fmt.Println(slice)

	// 要素の追加
	// array:要素の数が固定されているので追加できない
	// slice:append関数で実行する
	for n := 0; n < 10; n++ {
		slice = append(slice, n)
		fmt.Print("slice=")
		fmt.Println(slice)
	}

	fmt.Print("slice[2]=")
	fmt.Println(slice[2]) //配列の3番目を表示
	fmt.Print("slice[0:5]=")
	fmt.Println(slice[0:5]) //配列の1番目から6番目までを表示
	fmt.Print("slice[3:]=")
	fmt.Println(slice[3:]) //配列の4番目以降を表示
	fmt.Print("slice[:7]=")
	fmt.Println(slice[:7]) //配列の8番目までを表示
	fmt.Print("slice[:]=")
	fmt.Println(slice[:]) //配列の最初から最後までを表示

	var num []int
	num = slice[0:5]
	fmt.Print("len(num)=")
	fmt.Println(len(num)) // len関数でスライスの要素数を返す
	fmt.Print("cap(num)=")
	fmt.Println(cap(num)) // cap関数でスライスの容量を返す

	// sliceと同じ型であるslice2にsliceを代入する
	//
	var slice2 []int
	slice2 = slice
	fmt.Print("slice2=")
	fmt.Println(slice2)

	// slice2の値を変更すると、sliceの値も同様に変更される
	slice2[0] = 100
	fmt.Print("slice=")
	fmt.Println(slice)
	fmt.Print("slice2=")
	fmt.Println(slice2)

	// make関数でもsliceが作成できる
	var slice3 []int
	slice3 = make([]int, 10, 10)
	fmt.Print("slice3=")
	fmt.Println(slice3)

	// スライスの要素を削除する関数を作成して実行する
	slice[0] = 0
	fmt.Print("delete_before:slice=")
	fmt.Println(slice)
	// 3の値を持つ要素を削除する
	slice = delete(slice, 3)
	fmt.Print("delete_after :slice=")
	fmt.Println(slice)

	// sort関数で配列の要素をソートする
	var slice4 []int
	for n := 0; n < 100; n++ {
		slice4 = append(slice4, rand.Intn(100))
	}
	fmt.Print("sort_before:slice4=")
	fmt.Println(slice4)
	sort.IntSlice(slice4).Sort()
	fmt.Print("sort_after :slice4=")
	fmt.Println(slice4)

	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}
