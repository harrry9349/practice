package main

import (
	"fmt"
	"math/rand"
	"time"
)

func calcEuclidean(num1 int, num2 int) int {
	fmt.Println("--------------------")
	fmt.Printf("%dと%dで計算\n", num1, num2)

	// num1をnum2で割り切ることが出来たときのnum2の値が最大公約数
	if num1%num2 == 0 {
		// num2を返す
		fmt.Printf("%dと%dで割り切れたので互除法を終了する\n", num1, num2)
		fmt.Println("--------------------")
		return num2
	}

	// num1をnum2で割った剰余を代入する
	euc := num1 % num2
	fmt.Printf("%dと%dの剰余は%d\n", num1, num2, euc)

	// return文に自身の関数を設定し再帰する
	// 引数はnum2とeucに設定する
	return calcEuclidean(num2, euc)
}

func main() {

	rand.Seed(time.Now().Unix())
	start := time.Now()

	// 2つの1000以上2000以下の整数を作成する
	// ただし、number_1 >= number_2であること
	number_1 := rand.Intn(1000) + 1000
	number_2 := rand.Intn(1000) + 1000
	if number_2 >= number_1 {
		number_1, number_2 = number_2, number_1
	}

	fmt.Printf("%dと%dの最大公約数を求める：\n", number_1, number_2)

	// ユークリッドの互除法を実行する関数
	euclidean := calcEuclidean(number_1, number_2)

	fmt.Printf("%dと%dの最大公約数は%d\n", number_1, number_2, euclidean)

	// 2つの整数の最大公約数が1のとき、2つの整数は「互いに素」と呼ぶ
	if euclidean == 1 {
		fmt.Printf("すなわち、%dと%dは互いに素である。\n", number_1, number_2)
	}

	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}
