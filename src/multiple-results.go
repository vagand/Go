package main

import "fmt"

func swap2(x, y int) (int, int) {
	return y, x
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	c, d := swap2(3, 4)
	fmt.Println(a, b)
	fmt.Println(c, d)
}
