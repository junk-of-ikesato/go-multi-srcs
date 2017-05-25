package main

import (
	"fmt"
	"./sub1"
)

func main() {
	fmt.Println("hello")
	fmt.Printf("1-1: %v\n", Func1_1())
	fmt.Printf("1-2: %v\n", func1_2())
	fmt.Printf("2-1: %v\n", Func2_1())
	var f2 File2  = Func2_1()
	fmt.Printf("2-2a: %v\n", f2)
	f2.Name = "hoge"
	fmt.Printf("2-2b: %v\n", f2)
	fmt.Printf("sub1: %v\n", submodule1.Sub1())
	//fmt.Println(submodule1.sub2()) => Compile Error!
}
