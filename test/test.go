//It's a temp program
package main

import "fmt"

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x%y
        fmt.Printf("%d %d\n", x, y)
    }
    return x
}

func main() {
    gcd(6, 8)
}
