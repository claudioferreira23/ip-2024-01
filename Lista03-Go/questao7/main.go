package main

import "fmt"

func main(){

	for {
		var n int
		fmt.Scan(&n)

		if n == 0{
			break
		}

		x := make([]int, n)

		for i := 0; i < n; i++{
			fmt.Scan(&x[i])
		}
		
		var y []int
		for i, v := range x{
			if v <= i{
				y = append(y, v)
			}
			fmt.Printf("(%d) %d\n", i, len(y))
		}
	}
}