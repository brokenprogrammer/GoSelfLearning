package main

import "fmt"

func main() {
    sum, t1, t2, temp := 0, 1, 1, 0

    for t2 < 4e6 {
        temp = t2
        t2 = t1 + t2
        t1 = temp

        if isEven(t2) {
            sum += t2
        }
    }

    fmt.Println(sum)
}

func isEven(n int) bool{
    return n % 2 == 0
}
