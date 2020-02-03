package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	start := time.Now()

	slice := []float64{}
	var sum float64 //合計
	var ave float64 //平均

	for n := 0; n < 100; n++ {
		slice = append(slice, math.Trunc(rand.Float64()*100))
		fmt.Print(slice[n])
		if n%10 == 9 {
			fmt.Println("")
		} else {
			fmt.Print(",")
		}
		sum = sum + slice[n]
	}

	ave = sum / float64(len(slice))

	fmt.Printf("合計は%gです。\n", sum)
	fmt.Printf("平均は%gです\n", ave)

	var vari float64     //分散
	var sum_vari float64 // 分散導出用の合計
	var std_vari float64 // 標準偏差

	for n := 0; n < len(slice); n++ {
		sum_vari = sum_vari + (slice[n]-ave)*(slice[n]-ave)
	}

	vari = sum_vari / float64(len(slice))
	std_vari = math.Sqrt(vari)

	fmt.Printf("分散は%gです。\n", vari)
	fmt.Printf("標準偏差は%gです。\n", std_vari)

	// 標準偏差から偏差値を求める
	var std float64 //偏差値
	for n := 0; n < len(slice); n++ {
		std = 10*(slice[n]-ave)/std_vari + 50
		fmt.Printf("平均点%gで%g点の偏差値は%gです。\n", ave, slice[n], std)
	}

	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}
