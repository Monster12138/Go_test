package basename

import "strings"

func Basename2(s string) string {
    slash := strings.LastIndex(s, "/") //if find "/" slash = -1
    s = s[slash+1:]
    if dot := strings.LastIndex(s, "."); dot >= 0 {
        s = s[:dot]
    }
    return s
}
