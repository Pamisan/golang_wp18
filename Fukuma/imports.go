package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Printf("Now you have %g problems.", math.Sqrt(7)) //%gは%e(科学的記数法)、%f（指数なしの小数）のどちらか出力の短い方
}
