//package submodule2 => Compile Error!
package submodule1

type Sub2 struct {
	Name string
	Value int
}

func Sub2_1() Sub2 {
	var r Sub2
	r.Name = "hello"
	r.Value = 123
	return r
}
