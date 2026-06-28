package main
import "fmt"
func main() {
    var v, t, c int
    h := t/60
    d := v*h
    df := float64(d)
    cf := float64(c)
    dp := int(df / cf)
    fmt.Println(dp)

}
