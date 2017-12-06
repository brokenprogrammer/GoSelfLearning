package main

import "fmt"

func main() {
    fmt.Println(multiples(1000))
}

func multiples(n int) int {
    sum := 0

    for i:= 0; i < n; i++ {
        if i % 3 == 0 || i % 5 == 0 {
            sum += i
        }
    }

    return sum
}