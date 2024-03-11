package main

import "fmt"

type Thread struct {
	count int
	sum   int
}

func main() {

	result := func(values ...int) Thread {
		count := len(values)
		sum := 0

		for _, value := range values {
			sum += value
		}
		ans := Thread{count: count, sum: sum}
		return ans
	}(1, 2, 3, 4, 5, 6,7)
	fmt.Printf("Result is %v ", result)
}
