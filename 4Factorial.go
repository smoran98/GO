package main

import (
    "fmt"
    "math/big"
)

func main() {

    x := new(big.Int)
    x.MulRange(1, 10)
    fmt.Println(x)
}

