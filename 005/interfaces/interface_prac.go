package main

import "fmt"

/*
learn about interfaces and reveivers
under the hood interfaces are like tuples (value,type)
*/

type I interface {
	M()
}

type T struct {
	S string
}

// this means that T implements interface I
func (t T) M() {
	fmt.Println(t.S)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	var i I = T{"Hello"}
	i.M()

	var r interface{} = "hello"

	s, ok := r.(float64)
	fmt.Println(s, ok) //output will be 0,false

	do(21)
	do("hello")
	do(true)

}
