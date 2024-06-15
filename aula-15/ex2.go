package main

import (
	"fmt"
)

func main() {
    var n int
    
    fmt.Scan(&n)
    
	var x []float64
	
	for i := 0; i < n; i++ {
	    var num float64
	    fmt.Scan(&num)
	    x = append(x, num)
	}
	 
	 fmt.Print(soma(x))
}

func soma(x []float64) float64 {
    if len(x) == 0 {
       return 0
    }
    return x[0] + soma(x[1:])
}
