package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// 100*100の大きさの平面上で、5つの頂点を必ず一回だけ通るという制約のもとで、
// 距離が最小となる巡回ルートを見つける

// easy版は純粋に組み合わせを設定して全探索

// difficult版はシミュレーテッドアニーリング（焼きなまし法）という方法で
// 巡回セールスマン問題を「近似」する
// (巡回セールスマン問題は全探索より大幅に効率の良い方法を見つけられていないため近似という表現を用いる)

// 平面上の辺を格納する
var graph [][]int

// 平面上の点の位置を格納する
var plane [][]int

// 頂点の数
var Num = 30

// 温度
var t = 1000000

// 初期化
func init() {
	rand.Seed(time.Now().Unix())
	// 領域確保
	graph = make([][]int, Num)
	plane = make([][]int, Num)
	for n := 0; n < Num; n++ {
		// 辺の格納(起点,終点)
		graph[n] = []int{n, (n + 1) % Num}
		// 頂点の座標を格納(x,y)
		plane[n] = []int{rand.Intn(100), rand.Intn(100)}
	}
	fmt.Println("辺:", graph)
	fmt.Println("座標:", plane)
}

// 距離の測定
func measure(n int, m int) float64 {
	x := plane[n][0] - plane[m][0]
	y := plane[n][1] - plane[m][1]
	return math.Sqrt(float64(x*x) + float64(y*y))
}

// 測定ラッパー
func cost(n int, m int) float64 {
	return measure(n, m)
}

// 遷移確率
func prob(before float64, after float64) float64 {
	if before >= after {
		return 1.0
	}
	// シミュレーテッドアニーリングでは値が悪くなる場合でもある確率で更新を許す。
	// 確率P = e ^ (Δb / T) (e:自然対数の底、Δb:値の変化量、T:温度)
	// 最初に温度が大きい状態では高い確率で遷移を許すが、Tが徐々に小さくなるにつれて徐々に遷移し辛くなる
	return math.Exp((before - after) / float64(t))
}

// 変更
func change() {
	a := rand.Intn(Num)
	b := 0
	for {
		b = rand.Intn(Num)
		if b != a {
			break
		}
	}

	current := cost(graph[a][0], graph[a][1]) + cost(graph[b][0], graph[b][1])
	next := cost(graph[a][0], graph[b][1]) + cost(graph[a][1], graph[b][0])

	if prob(current, next) >= math.Abs(rand.Float64()) && graph[a][0] != graph[b][1] && graph[a][1] != graph[b][0] {
		graph[a][0], graph[b][0] = graph[b][0], graph[a][0]
	}
}

func main() {

	start := time.Now()

	sum := 0.0
	// 最初のcost
	for i := 0; i < Num; i++ {
		sum += cost(graph[i][0], graph[i][1])
	}

	fmt.Println("最初のcost:", sum)

	//温度を下げながらループ
	for ; t > 1; t-- {
		change()
	}

	sum = 0.0
	for i := 0; i < Num; i++ {
		sum += cost(graph[i][0], graph[i][1])
	}

	//終了後のcost
	fmt.Println("終了後のcost:", sum)
	//終了後の組み合わせ
	fmt.Println("組み合わせ", graph)

	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}
