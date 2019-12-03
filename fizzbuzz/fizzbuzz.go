package main

import(
	"fmt"
	"time"
)

func fizzbuzzer(max int) {
	for n := 1; n <= max; n++ {
		if n % 3 == 0 {
			if n % 5 == 0 {			
				fmt.Println(n ,":FizzBuzz")
			} else {
				fmt.Println(n ,":Fizz")
			}
		} else {			
			if n % 5 == 0 {			
				fmt.Println(n ,":Buzz")
			} else {
				fmt.Println(n)
			}
		}
	}
}

func main(){
	start := time.Now()
	max := 100
	fizzbuzzer(max)
	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}	