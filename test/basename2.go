func basename(s string) string {
    slash := string.LastIndex(s, "/") //if find "/" slash = -1
    s = s[slash+1:]
    if dot := strings.LastIndex(s, "/"); dot >= 0 {
        s = s[:dot]
        return s
    }
}
