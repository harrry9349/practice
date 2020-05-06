//野球シミュレーション

//プログラムはイニングのイベントを読み取り、そのイニングのスコアを出力します。次の3つのイベントのみがあります。

//シングルヒット
//一塁にランナーを置きます。
//1塁のランナーは2塁に進み、2塁のランナーは3塁に進みます。
//3塁のランナーはホームベースに進み（そしてベースから出る）、得点にポイントが加算されます。

//ホームラン
//ベースのすべてのランナーがホームベースに進みます。
//スコアは、ランナーの数に1を足した数に加算されます。

//アウト
//アウトの数が1増えます。
//ランナーとスコアは変化しません。
//イニングはスリーアウトで終了します。

//これらのイベントをそれぞれ「HIT」、「HOMERUN」、「OUT」で表現する。
//9回までイニングを行い、最終的な勝者を表示します。9回までの得点が同じ場合は引き分けとします。

//イニングのイベントを読み取り、そのイニングのスコアと、ゲームの勝者を出力するプログラムを作成します。

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	start := time.Now()

	// イベントの定義
	input := make([]string, 0)

	// HIT:25%,HOMERUN:5%,OUT:70%
	for i := 0; i < 25; i++ {
		input = append(input, "HIT")
	}
	for i := 0; i < 5; i++ {
		input = append(input, "HOMERUN")
	}
	for i := 0; i < 70; i++ {
		input = append(input, "OUT")
	}

	// 各イニングごとの得点
	point := make([]int, 18)

	//イベント格納用
	events := make([]string, 0)

	//現在のイニング
	inning := 0
	// アウトの数
	outs := 0
	// ベース上の人数
	onbase := 0

	// 各プレイヤーの総得点数
	player1 := 0
	player2 := 0

	for {
		// イベントの決定
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Millisecond * 1)
		event := input[rand.Intn(100)]
		events = append(events, event)

		if event == "HIT" {
			// HITの場合
			if onbase < 3 {
				//ベース上の人数が3人より少ない場合は人数を増やす
				onbase++
			} else {
				// ベース上の人数が3人の場合は得点を1点入れる
				point[inning]++
			}
		} else if event == "HOMERUN" {
			// HOMERUNの場合
			// ベース上の人数+1点だけ得点を入れて、ベース上の人数を0に戻す
			point[inning] += onbase + 1
			onbase = 0
		} else {
			// OUTの場合
			// アウトの数を一つ増やす
			outs++
			if outs == 3 {
				// アウトの数が3つになった場合
				// アウトの数を0に戻す
				outs = 0
				if inning == 17 {
					//18イニング目（9回裏）終了時
					//ゲームを終了する
					break
				} else {
					// 17,18イニング(9回)でない場合
					//イニングを1つ進める
					inning++
				}
			}
		}
	}

	points1 := make([]int, 0)
	points2 := make([]int, 0)

	for i, v := range point {
		if i%2 == 0 {
			points1 = append(points1, v)
		} else {
			points2 = append(points2, v)
		}
	}

	fmt.Println("イベント:", events)

	fmt.Println("得点：", point)
	fmt.Println("先攻：", points1)
	fmt.Println("後攻：", points2)

	for _, v := range points1 {
		player1 += v
	}

	for _, v := range points2 {
		player2 += v
	}

	if player2 < player1 {
		fmt.Println(player1, "-", player2, "で先攻の勝利です")
	} else if player1 < player2 {
		fmt.Println(player1, "-", player2, "で後攻の勝利です")
	} else {
		fmt.Println(player1, "-", player2, "で引き分けです")
	}

	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())

}
