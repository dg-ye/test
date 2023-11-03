package main

import (
	"fmt"
	"github.com/dg-ye/test/lib"
	"github.com/patrickmn/go-cache"
	"time"
)

// The global cache for register
var _register *cache.Cache

func init() {
	_register = cache.New(30*time.Minute, 3*time.Minute)
}

type A1 struct {
	name string
	age  int8
}

func (a1 *A1) print() {
	println(a1.name)
}

type printer interface {
	print()
}

func printName(p printer) {
	p.print()

	var p2 printer
	p2 = p
	if a3, ok := p2.(*A1); ok {
		a3.name = "jack"
	}
	p.print()
}

func main() {
	a2 := A1{"armmy", 13}
	printName(&a2)
	fmt.Println(a2)
	sum := lib.Add(1, 2)
	fmt.Print(sum)
}
