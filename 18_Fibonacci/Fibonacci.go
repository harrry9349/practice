package main

import (
	"fmt"
	"math/rand"
	"time"
)

func calcFibonacci(num int) int {
	// n番目(n >= 2)のフィナボッチ数をF_nとすると
	// F_0 = 0, F_1 = 1 ,F_n = F_(n-1) + F_(n-2)で表される
	fmt.Printf("現在の値：%d\n", num)
	if num == 0 {
		fmt.Println("F_0 = 0 のため0を返す")
		return 0
	} else if num == 1 {
		fmt.Println("F_1 = 1 のため1を返す")
		return 1
	} else {
		fmt.Println("F_n = F_(n-1) + F_(n-2)のため再帰する")
		return calcFibonacci(num-1) + calcFibonacci(num-2)
	}

}

func main() {

	rand.Seed(time.Now().Unix())
	start := time.Now()

	fib := []int{}
	for num := 0; num <= 20; num++ {
		fmt.Printf("---%d番目のフィボナッチ数を求める：---\n", num)
		fib = append(fib, calcFibonacci(num))
		fmt.Println("-----------------------------------")
	}

	fmt.Println("フィボナッチ数列：", fib)

	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())

}
