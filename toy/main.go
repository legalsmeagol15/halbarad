package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

type Foo struct {
	the_vec *mat.VecDense
}

func main() {
	s := Foo{
		the_vec: &mat.VecDense{},
	}
	if s.the_vec != nil {
		fmt.Println("Hello, Gonum!")
	} else {
		fmt.Println("booo!")
	}
}
