// 平方採中法
// 古典的な乱数生成方法の一つである平方採中法のプログラムを作成します。平方採中法は、フォンノイマンによって 1940 年代半ばに提案された方法です。

// 平方採中法は、生成する乱数の桁数を n としたとき、初期値 s の2乗を計算し、その数値を 2n 桁の数値とみて、（下の例のように 2 乗した桁数が足りないときは、0 を補います。）
// その中央にある n 個の数字を最初の乱数とします。次にこの乱数を 2 乗して、同じ様に、中央にある n 個の数字をとって、次の乱数とします。例えば、123 を初期値とすると

// 1232 = 00015129 → 0151
// 1512  = 00022801 → 0228
// 2282  = 00051984 → 0519
// 5192  = 00269361 → 2693
// 26932  = 07252249 → 2522
// の様になります。この方法を用いて、初期値 s（10000未満の正の整数）を入力とし、n = 4 の場合の乱数を 10 個生成し出力するプログラムを作成してください。

// Input
// 複数のデータセットが与えられます。１行目にデータセットの数 d (d ≤ 10) が与えられます。各データセットとして、１行に初期値 s（整数、1 ≤ s < 10000）が与えられます。

// Output
// 各データセットに対して、

// Case x: (x は 1 から始まるデータセット番号)
// １個目の生成した乱数（整数）
// ２個目の生成した乱数（整数）
// ：
// ：
// １０個目の生成した乱数（整数）
// を出力してください。

// Sample Input
// 2
// 123
// 567

// Output for the Sample Input
// Case 1:
// 151
// 228
// 519
// 2693
// 2522
// 3604
// 9888
// 7725
// 6756
// 6435
// Case 2:
// 3214
// 3297
// 8702
// 7248
// 5335
// 4622
// 3628
// 1623
// 6341
// 2082

// 元問題
//http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=0137

package main

import (
	"fmt"
	"time"
)

func Pow(x int, y int) int {
	res := 1
	for i := 0; i < y; i++ {
		res = res * x
	}
	return res
}

func MiddleSquare(value int) int {

	n := 4
	// 初期値を2乗する
	sqr := value * value

	// 8桁分取り出す
	res := sqr % Pow(10, n*2)

	// 下2桁を除外する
	res = res / Pow(10, n/2)
	// 上2桁を除外する
	res = res % Pow(10, n)

	fmt.Println(res)
	return res
}

func main() {

	starttime := time.Now()

	cases := 2
	values := make([]int, cases)
	values = []int{123, 567}

	for i := 0; i < cases; i++ {
		fmt.Println("case ", i+1, ":")
		value := values[i]
		for j := 0; j < 10; j++ {
			value = MiddleSquare(value)
		}
	}

	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(starttime)).Seconds())

}
