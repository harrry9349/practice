package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// 10*10の大きさの平面上で、5つの頂点を必ず一回だけ通るという制約のもとで、
// 距離が最小となる巡回ルートを見つける

// 平面上の点の位置を格納する
var plane [][]int

// 頂点の数
var Num = 4

// 距離の測定(辺の起点,辺の終点)
func measure(n int, m int) float64 {
	x := plane[n][0] - plane[m][0]
	y := plane[n][1] - plane[m][1]
	return math.Sqrt(float64(x*x) + float64(y*y))
}

func main() {

	start := time.Now()

	// 初期化
	rand.Seed(time.Now().Unix())
	// 領域確保
	plane = make([][]int, Num)
	for n := 0; n < Num; n++ {
		// 頂点の座標を格納(x,y)
		plane[n] = []int{rand.Intn(10), rand.Intn(10)}
	}
	fmt.Println("頂点の座標:", plane)

	// 組み合わせの設定
	iterate := [][]int{{0, 1, 2, 3}, {0, 2, 1, 3}, {0, 1, 3, 2}, {0, 3, 1, 2}, {0, 2, 3, 1}, {0, 3, 2, 1}}

	cnt := 0
	sum := 0.0

	// 測定
	for i, v := range iterate {
		current := 0.0
		for j := 0; j < Num; j++ {
			current += measure(v[j], v[(j+1)%Num])
		}

		fmt.Println("今回の組み合わせ:", v, "今回の距離", current, "現在の最短距離:", sum)
		if i == 0 {
			sum = current
			cnt = i
		} else {

			if current < sum {
				fmt.Println("更新:", i, v)
				sum = current
				cnt = i
			}
		}

	}

	//最短距離
	fmt.Println("最短距離:", sum)
	//最短組み合わせ
	fmt.Println("最短組み合わせ:", iterate[cnt])

	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}
