//It's a temp program
package main

import (
    "basename"
    "fmt"
)

func main() {
    var s string = "a/b/c.go"
    s = basename.Basename(s)
    fmt.Println(s)
}
