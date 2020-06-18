package main

import (
	"fmt"
	"testlib/liba"
)

func main() {
	fmt.Println("main.go")

	s := liba.Show{}
	s.Foo()
	s.Foo2()
	liba.Fshow()
}
