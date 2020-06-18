package liba

import (
	"fmt"
	_ "testlib/libc"
)

type Show struct{}
type hide struct{}

func init() {
	fmt.Println("liba.go init")
}

func (f Show) Foo() {
	fmt.Println("Show.Foo")
}

func (f Show) foo() {
	fmt.Println("Show.foo")
}

func (f Show) Foo2() {
	f.foo()
}

func fhide() {
	fmt.Println("fhide")
}

func Fshow() {
	fmt.Println("Fshow")
}
