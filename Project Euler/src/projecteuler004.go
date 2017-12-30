package main

import (
    "fmt"
    "strconv"
)

func main() {
    largest := 0

    for a := 100; a < 1000; a++ {
        for b := 100; b < 1000; b++ {
            i := a * b

            if isPalindrome(i) {
                if i > largest {
                    largest = i
                }
            }
        }
    }

    fmt.Println(largest)
}

func isPalindrome(x int) bool {
    target := strconv.Itoa(x)
    mid := len(target) / 2
    last := len(target) - 1

    for i := 0; i < mid; i++ {
        if target[i] != target[last-i] {
            return false
        }
    }

    return true
}
