//echo2输出命令行参数
package main

import (
    "fmt"
    "os"
)

func main() {
    s, sep := "", ""
    fmt.Print("Args: ")
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)
}
