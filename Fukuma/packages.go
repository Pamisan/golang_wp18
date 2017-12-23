package main

import (
    "fmt"
    "math/rand" //mathパッケージが持っているrandパッケージをimport
)

func main() {
    fmt.Println("My favorite number is", rand.Intn(10))
}
