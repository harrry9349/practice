package main

import (
	"fmt"
	"math"
	"time"
)

// 4x4マスの迷路を開始地点(A)から終点(H)まで探索して移動する
// H F G #
// # E # #
// # D B A
// # # C #
//

func CheckContains(graph [][]string, i int, j int, check []string) bool {
	for _, v := range check {
		if graph[i][j] == v {
			return true
		}
	}
	return false
}

func main() {

	start := time.Now()

	// 平面
	graph := [][]string{{"H", "F", "G", "#"}, {"#", "E", "#", "#"}, {"#", "D", "B", "A"}, {"#", "#", "C", "#"}}

	// キュー
	queue := make([]string, 0)
	queue = append(queue, "A")

	// チェックリスト
	check := make([]string, 0)
	check = append(check, "A")

	count := 0 //カウント
	rNow := 0  // 現在位置（行）
	cNow := 0  // 現在位置（列）

	fmt.Println("H F G #")
	fmt.Println("# E # #")
	fmt.Println("# D B A")
	fmt.Println("# # C #")

	// 開始地点Aの探索
	for i, v := range graph {
		for j := 0; j < len(v); j++ {
			if v[j] == "A" {
				rNow, cNow = i, j
			}
		}
	}

	rStart, cStart := rNow, cNow
	fmt.Printf("開始位置:(%d,%d)\n", rStart, cStart)

	for {
		// 処理回数のカウント
		count++
		fmt.Print("現在のキュー:", queue)
		fmt.Println(",現在のチェック", check)

		// 現在位置の取得
		for i, v := range graph {
			for j := 0; j < len(v); j++ {
				if queue[0] == v[j] {
					rNow, cNow = i, j

				}
			}
		}

		fmt.Printf("現在の位置:(%d,%d)\n", rNow, cNow)

		// ゴールの判定
		if len(queue) != 0 {
			if queue[0] == "H" {
				fmt.Println("ゴール:")
				break
			}
		}

		// queueのpop　Golangの場合は単純に最初の要素を除くスライスを代入する
		queue = queue[1:]

		// ひとつ下の探索
		if rNow < 3 {
			// 対象が壁(#)ではない
			if graph[rNow+1][cNow] != "#" {
				// 対象の文字がチェックに含まれていない
				if !CheckContains(graph, rNow+1, cNow, check) {
					// キューに対象情報を追加
					queue = append(queue, graph[rNow+1][cNow])
					check = append(check, graph[rNow+1][cNow])
					fmt.Print("キュー追加（下）:", queue)
					fmt.Println(",チェック追加（下）:", check)
				}
			}
		}

		// ひとつ右の探索
		if cNow < 3 {
			// 対象が壁(#)ではない
			if graph[rNow][cNow+1] != "#" {
				// 対象の文字がチェックに含まれていない
				if !CheckContains(graph, rNow, cNow+1, check) {
					// キューに対象情報を追加
					queue = append(queue, graph[rNow][cNow+1])
					check = append(check, graph[rNow][cNow+1])
					fmt.Print("キュー追加（右）:", queue)
					fmt.Println(",チェック追加（右）:", check)
				}
			}

		}

		// ひとつ上の探索
		if rNow > 0 {
			// 対象が壁(#)ではない
			if graph[rNow-1][cNow] != "#" {
				// 対象の文字がチェックに含まれていない
				if !CheckContains(graph, rNow-1, cNow, check) {
					// キューに対象情報を追加
					queue = append(queue, graph[rNow-1][cNow])
					check = append(check, graph[rNow-1][cNow])
					fmt.Print("キュー追加（上）:", queue)
					fmt.Println(",チェック追加（上）:", check)
				}
			}

		}

		// ひとつ左の探索
		if cNow > 0 {
			// 対象が壁(#)ではない
			if graph[rNow][cNow-1] != "#" {
				// 対象の文字がチェックに含まれていない
				if !CheckContains(graph, rNow, cNow-1, check) {
					// キューに対象情報を追加
					queue = append(queue, graph[rNow][cNow-1])
					check = append(check, graph[rNow][cNow-1])
					fmt.Print("キュー追加（左）:", queue)
					fmt.Println(",チェック追加（左）:", check)
				}
			}

		}

	}

	//出力
	fmt.Println("終点までの距離:", math.Abs(float64(rNow-rStart))+math.Abs(float64(cNow-cStart)))
	fmt.Println("終点までの処理回数:", count)

	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}
