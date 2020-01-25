package main

import (
	"fmt"
	"math/rand"
	"time"
)

// n人の生徒がある教科のテストに挑んだ。
// n人の生徒の中から、テストの点数の最大値を獲得した生徒、最小値を獲得した生徒を検索したい。

// 構造体の定義
type Student struct {
	Num   int
	Name  string
	Point int
}

//構造体メソッドの定義
func (s Student) ShowMax() {
	fmt.Printf("点数が最も高かった生徒は番号%d、名前%sで、点数は%dです。\n", s.Num, s.Name, s.Point)
}

func (s Student) ShowMin() {
	fmt.Printf("点数が最も低かった生徒は番号%d、名前%sで、点数は%dです。\n", s.Num, s.Name, s.Point)
}

//Student構造体スライスの定義
type Test []Student

// ランダム文字列を生成する関数
func CreateName(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {

	rand.Seed(time.Now().Unix())
	start := time.Now()

	// Testスライスの宣言
	res := make(Test, 0)

	// 100人の名前と0~100のランダム値をTestスライスに格納する
	for n := 0; n < 100; n++ {
		res = append(res, Student{n + 1, CreateName(20), rand.Intn(101)})
		fmt.Printf("番号： %d 名前：%s 点数：%d \n", res[n].Num, res[n].Name, res[n].Point)
	}

	//最大値を取得する
	fmt.Println("最大値")
	//最大値のためのStudent構造体を作成
	max := res[0]
	// 1番目から終点まで比較をとる
	for n := 0; n < len(res); n++ {
		// n番目の値が最大値より大きい場合
		if max.Point < res[n].Point {
			//最大値を更新する
			fmt.Printf("更新： 前番号：%d 前名前：%s 前点数：%d  後番号：%d 後名前：%s 後点数：%d \n", max.Num, max.Name, max.Point, res[n].Num, res[n].Name, res[n].Point)
			max = Student{res[n].Num, res[n].Name, res[n].Point}
		}
	}

	//最小値を取得する
	fmt.Println("最小値")
	min := res[0]
	// 1番目から終点まで比較をとる
	for n := 0; n < len(res); n++ {
		// n番目の値が最小値より小さい場合
		if min.Point > res[n].Point {
			//最小値を更新する
			fmt.Printf("更新： 前番号：%d 前名前：%s 前点数：%d  後番号：%d 後名前：%s 後点数：%d \n", min.Num, min.Name, min.Point, res[n].Num, res[n].Name, res[n].Point)
			min = Student{res[n].Num, res[n].Name, res[n].Point}
		}
	}

	// Pointが最小、最大のStudentのNumとNameとPointを表示する
	// なお、Pointが同値の場合、Numが小さい方を優先する
	max.ShowMax()
	min.ShowMin()

	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}
