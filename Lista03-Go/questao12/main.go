package main

import (
    "fmt"
    "sort"
)

func main() {
       
        var n int
        fmt.Scan(&n)

        x := make([]int, n, 1000)

        for i := 0; i < n; i++{
        fmt.Scan(&x[i])
        }
       
        fmt.Println("SAIDA")
        sort.Ints(x)
        for _, v := range x{
            fmt.Printf("%d\n", v)
        }
       
}