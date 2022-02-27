// https://codeforces.com/problemset/problem/1131/C

package main

import (
    "fmt"
    "sort"
    "os"
    "bufio"
)

func main() {
    in := bufio.NewReader(os.Stdin)
    out := bufio.NewWriter(os.Stdout)
    defer out.Flush()
    var n int
    fmt.Fscan(in, &n)
    ar := make([]int, n)
    for i := range ar {
        fmt.Fscan(in, &ar[i])
    }
    sort.Ints(ar)
    
    t := make([]int, n)
    if n <= 3 {
        for i := range ar {
            t[i] = ar[i]
        }
    } else {
        l := (n + 1) / 2 - 1
        r := l + 1
        isLeft := true
        i := 0
        for i < n {
            if isLeft {
                t[l] = ar[i]
                l--
            } else {
                t[r] = ar[i]
                r++
            }
            i++
            isLeft = !isLeft
        }
    }
    
    for i := range t {
        if i != 0 {
            fmt.Fprint(out, " ")
        }
        fmt.Fprint(out, t[i])
    }
    fmt.Fprintln(out)
}
