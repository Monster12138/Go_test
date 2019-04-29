//basename remove arg's path and . suffix
//e.g., a.go => a, a/b/c.go => c, a/b.c.go => b.c
package basename

type INT_t int

func Basename(s string) string {
    //remove before the last '/' 
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == '/' {
            s = s[i+1:]
            break
        }
    }

    //retain before the last '.'
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == '.' {
            s = s[:i]
            break
        }
    }
    return s
}
