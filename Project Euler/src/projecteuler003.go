package main

import "fmt"

func main() {
    number := 600851475143
    largest := 0

    for i := 1; i <= number; i++ {
        if number % i == 0 {
            number = number / i

            if i > largest {
                largest = i
            }
        }
    }

    fmt.Println(largest)
}
