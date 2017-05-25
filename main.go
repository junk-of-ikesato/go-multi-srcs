package main

import (
	"fmt"
	"./sub1"
)

func main() {
	fmt.Println("hello")
	fmt.Println(Func1())
	fmt.Println(func2())
	fmt.Println(submodule1.Sub1())
	//fmt.Println(submodule1.sub2()) => Compile Error!
}
