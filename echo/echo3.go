//echo3输出命令行参数
package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Print("Args: ")
    fmt.Println(os.Args[1:])
}
