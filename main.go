package main

import "fmt"

func Sum(x int64, y int64) (z int64) {
	return (int64(x + y))
}

func main() {
	fmt.Println(Sum(10, 11))
}
