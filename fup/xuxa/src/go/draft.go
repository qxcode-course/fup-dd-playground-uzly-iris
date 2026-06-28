package main 
import (
    "fmt"
    "bufio"
    "os"
)
func main() { 
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    str := scanner.Text()

    for i := len(str) - 1 ; i >= 0 ; i-- {
        fmt.Print(string(str[i]))
    }

    fmt.Println()
}