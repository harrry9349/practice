package main

import(
	"fmt"
	"time"
	"strconv"
)

func fizzbuzzer(n int) string {
	var res string
	res = strconv.Itoa(n) + ":"
		if n % 3 == 0 {
			res = res + "Fizz"
			if n % 5 == 0 {			
				res = res + "Buzz"
			}
		} else if n % 5 == 0 {				
			res = res + "Buzz"
		}

		return res
}

func main(){
	start := time.Now()
	max := 100
	for n := 1; n <= max; n++ {
		fmt.Println(fizzbuzzer(n))
	}
	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}	