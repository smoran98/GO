package main
import (
    "fmt"
    "time"
)
func main() {
    s := 0
    for {
        if s >= 6 { break }
        fmt.Println("s is:", s)
        s++
        time.Sleep(1000 * time.Millisecond)
    }
    fmt.Println("Break happened and 6ixth number guessed")
}