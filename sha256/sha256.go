package main

import ( 
    "crypto/sha256"
    "fmt"
)

func zero(ptr *[32]byte) {
    *ptr = [32]byte{}
}

func cmp(s1 *[32]uint8, s2 *[32]uint8) int{
    dif_count := 0
    for i := range s1 {
        if s1[i] != s2[i] {
            dif_count++
        }
    }
    return dif_count
}

func main() {
    c1 := sha256.Sum256([]byte("x"))
    c2 := sha256.Sum256([]byte("X"))
    fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
    fmt.Printf("different count = %d\n", cmp(&c1, &c2))
}
