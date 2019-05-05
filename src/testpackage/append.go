package testpackage

func AppendInt(x []int, y int) []int {
    var z []int
    zlen := len(x) + 1
    if zlen <= cap(x) {
        //slice still has space to grow
        z = x[:zlen]
    } else {
        //slice has none space to grow
        zcap := zlen
        if zcap < 2*len(x) {
            zcap = 2 * len(x)
        }
        z = make([]int, zlen, zcap)
        copy(z, x)
    }    
    z[len(x)] = y
    return z
}
