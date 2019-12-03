package main

import(
	"fmt"
)

func main(){
	max := 100
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