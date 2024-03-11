package main

import "fmt"

type Thread struct {
	count int
	sum   int
}

func ProAdder(values ...int) Thread {
	count := len(values);
	sum := 0;
	for _ , value := range values {
		sum += value
	}
	ans := Thread { count: count, sum : sum };
	return ans;
}
func main() {

	fmt.Printf("%v \n",ProAdder(1,2,3,4,5,6,7));
	fmt.Printf("%v \n",ProAdder(1,2,3,4,5,6,7,8));
	fmt.Printf("%v \n",ProAdder(1,2,3,4,5,6,7,8,9));
	fmt.Printf("%v \n",ProAdder(1,2,3,4,5,6,7,8,9,10));
}
