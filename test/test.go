//It's a temp program
package main

import (
    "basename"
    "fmt"
)

func main() {
    var s string = "a/b/cccc.go"
    s = basename.Basename2(s)
    fmt.Println(s)
}
