//It's a temp program
package main

import (
    "fmt"
    "os"
    "bufio"
    "io"
    "unicode"
    "unicode/utf8"
    "encoding/json"
    "log"
)

func nonempty(strings []string)[]string {
    i := 0
    for _, s := range strings {
        if s != "" {
            strings[i] = s
            i++
        }
    }
    return strings[:i]
}

func nonempty2(strings []string) []string {
    out := strings[:0]
    for _, s := range strings {
        if s != "" {
            out = append(out, s)
        }
    }
    return out
}

func remove(slice []int, i int) []int {
    slice[i] = slice[len(slice) - 1]
    return slice[:len(slice) - 1]
}

func routate(x []int, n int) []int {
    temp := x[:n]
    x = x[n:]
    for _, a := range temp {
        x = append(x, a)
    } 
    return x
}

func maptest() {
    ages := make(map[string]int)
    ages["alice"] = 31
    ages["charlie"] = 34

    if age, ok := ages["bob"]; !ok {
        fmt.Print("bob is't in ages map\n")
        _ = age
    }
}

func equal_map(x, y map[string]int) bool {
    if len(x) != len(y) {
        return false
    } 
    for k, xv := range x {
        if yv, ok := y[k];  !ok || yv != xv {
            return false
        }
    }
    return true
}

func dup2() {
    seen := make(map[string]bool)
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        line := input.Text()
        if !seen[line] {
            seen[line] = true
            fmt.Println(line)
        }
    }

    if err := input.Err(); err != nil {
        fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
        os.Exit(1)
    }
}
/*
var m = make(map[string]int)
func k(list []string) string { return fmt.Sprintf("%q", list) }
func add(list []string) { m[k(list)++] }
func Count(list []string)int { return m[k(list)] }
*/

func unicode_test() {
    counts := make(map[rune]int)
    var utflen [utf8.UTFMax + 1]int
    invalid := 0
    in := bufio.NewReader(os.Stdin)
    for {
        r, n, err := in.ReadRune() 
        if err == io.EOF {
            break
        }
        if err != nil {
            fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
            os.Exit(1)
        }
        if r == unicode.ReplacementChar && n == 1 {
            invalid++
            continue
        }
        counts[r]++
        utflen[n]++
    }
    fmt.Printf("rune\tcount\n")
    for c, n := range counts {
        fmt.Printf("%q\t%d\n",c ,n)
    }
    fmt.Print("\nlen\tcount\n") 
    for i, n := range utflen {
        if i > 0 {
            fmt.Printf("%d\t%d\n", i, n)
        }
    }
    if invalid > 0 {
        fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
    }
}

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
    edges := graph[from]
    if edges == nil {
        edges = make(map[string]bool)
        graph[from] = edges
    }
    edges[to] = true
}
func hasEdge(from, to string)bool {
    return graph[from][to]
}

func struct_test() {
    type Point struct {
        X, Y int
    }

    type Circle struct {
        Point
        Radius int
    }

    type Wheel struct {
        Circle
        Spokes int
    }

    var w Wheel
    w.X = 8
    w.Y = 8
    w. Radius = 5
    w.Spokes = 20

    w = Wheel{Circle{Point{8, 8}, 5}, 20}

    w = Wheel{
        Circle: Circle{
            Point: Point{X: 8, Y: 8},
            Radius: 5,
        },
        Spokes: 20,
    }

    fmt.Printf("%#v\n", w)

    w.X = 42

    fmt.Printf("%#v\n", w)
}

func json_test() {
    type Movie struct {
        Title string
        Year int `json:"released"r`
        Color bool `json:"color,omitempty"`
        Actors []string
    }

    var movies = []Movie {
        { Title:"Casablanca", Year:1942, Color: false,
            Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
        { Title:"Cool Hand Luke", Year: 1967, Color: true,
            Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
    }
    data, err := json.Marshal(movies)
    if err != nil {
        log.Fatalf("JSON marshaling failed: %s", err)
    }
    fmt.Printf("%s\n", data)
}

func main() {
    json_test()
}
