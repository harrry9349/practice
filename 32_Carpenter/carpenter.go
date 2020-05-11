//大工の褒美（改題）

//ある日、殿様は一人の大工に、「台風や地震が来たときに町人が避難できる、頑丈で大きな建物を造りなさい。」と命じました。し
//かし、その頑丈で大きな建物を完成させるには、大きな太い柱が必要です。町（町１）にそんな大きな柱はありません。
//そこで、大工は遠くの山里（町６）まで大きな柱を調達しに行くことになりました（大工は町から里山に行って、町に戻ってくる必要はありません）。

//大工の褒美は、殿様から受け取ったお金から交通費を差し引いた余りです。
//下の地図のように、山里に行くには、いろいろな町を通るたくさんの街道があり、2つの町をつなぐ街道はそれぞれ交通費が違います。
//大工の褒美を最大にするにはどのように街道をたどり調達すればよいでしょうか。
//最大となる大工の褒美を出力するプログラムを作成してください。
//ただし、町の数を n とすれば、各町は 1 から n までの整数で識別されます。2 つの町を直接つなぐ街道は 2 本以上ありません。

// 元問題
//http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=0117

package main

import (
	"fmt"
	"time"
)

type Road struct {
	town1 int
	town2 int
	cost  int
}
type List []Road

func BellmanFord(road List, start int, goal int) int {

	// 各町の距離（町０は存在しない）
	dist := []int{100, 100, 100, 100, 100, 100, 100}

	// 起点の町の距離を0にする
	dist[start] = 0

	// 最短距離の調査を実行する
	fmt.Println(dist)

	for {
		isUpdated := false
		for i := 0; i < len(road); i++ {
			r := road[i]
			fmt.Println(i, " road[i]=", r)

			fmt.Println(dist[r.town1], r.cost, dist[r.town2])
			if dist[r.town1] != 100 {
				if dist[r.town1]+r.cost < dist[r.town2] {
					fmt.Println("更新：", dist[r.town1], r.cost, dist[r.town2])
					isUpdated = true
					dist[r.town2] = dist[r.town1] + r.cost
				}
			}
			fmt.Println("dist", dist)
		}

		fmt.Println(isUpdated, dist)
		if isUpdated == false {
			break
		}
	}

	fmt.Println(dist[goal])
	return dist[goal]
}

func main() {

	starttime := time.Now()

	// サイトの例に沿ってアルゴリズムを実行する
	// 町の数：６　街道の数：８とする
	// 起点：１　終点：６とする
	//towns := 6
	roads := 8

	// 街道情報の定義
	// 1,2列目：街道がつないでいる町の情報 3列目：交通費
	input := make(List, roads)

	input = List{Road{1, 2, 2}, Road{1, 3, 4}, Road{1, 4, 4}, Road{2, 5, 3}, Road{3, 4, 4}, Road{3, 6, 1}, Road{4, 6, 1}, Road{5, 6, 1},
		Road{2, 1, 2}, Road{3, 1, 3}, Road{4, 1, 2}, Road{5, 2, 2}, Road{4, 3, 2}, Road{6, 3, 2}, Road{6, 4, 1}, Road{6, 5, 2}}

	// 出発する町の番号
	start := 2
	// 目的地の町の番号
	goal := 4

	// 殿様から受け取った報酬
	prize := 50

	// 行き1→6についてベルマンフォード法を実行する
	result := prize - BellmanFord(input, start, goal)

	fmt.Println("大工の褒章は", result, "です。")

	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(starttime)).Seconds())

}
