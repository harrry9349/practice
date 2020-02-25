package main

import (
	"flag"
	"fmt"
	"math"
	"time"
)

func isPrime(num int) bool {

	// 指定した数が2より小さければ素数ではない
	if num < 2 {
		return false
	}

	// 指定した数の平方根を求める
	numsqrt := int(math.Sqrt(float64(num)))
	for n := 2; n <= numsqrt; n++ {
		// 指定した数の平方根の数だけ判定すればよい
		if num%n == 0 {
			// 指定した数を割り切ることができる2以上の整数があれば素数ではない
			return false
		}
	}

	// 上記の条件を満たさなければ素数である
	return true
}

func main() {

	// コマンド　go run primenumber.go -low 整数 -high 整数
	start := time.Now()

	var (
		// 下限
		low = flag.Int("low", 0, "int flag")
		// 上限
		high = flag.Int("high", 0, "int flag")
	)
	flag.Parse()

	//下限＞上限ならば入れ替え
	if *low > *high {
		low, high = high, low
	}

	if *low != 0 && *high != 0 {
		for n := *low; n <= *high; n++ {
			if isPrime(n) {
				fmt.Printf("%dは素数です。\n", n)
			} else {
				fmt.Printf("%dは素数ではありません。\n", n)
			}
		}
	} else {
		// 下限または上限が未入力の場合はメッセージを返す
		fmt.Println("引数が入力されていません")
	}

	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}
