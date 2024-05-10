package main

import "fmt"

func main() {

	var n int
	var k int

	fmt.Scan(&n)

	x := make([]int, n, 1000)

	for i := 0; i < n; i++{
		fmt.Scan(&x[i])
	}

	fmt.Scan(&k)

	var y []int

	for _, v := range x{
		if v >= k{
			y = append(y, v)
		}
	}
	fmt.Print(len(y))
}
