package main

type File2 struct {
	Name string
	Value string
}

func Func2_1() File2 {
	var r File2
	r.Name = "hello"
	r.Value = "world"
	return r
}
