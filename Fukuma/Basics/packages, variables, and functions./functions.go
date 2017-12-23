package main

import "fmt"

func add(x int, y int) int { //引数リストの後ろに方を指定することで、関数から戻り値を返すことができる
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
