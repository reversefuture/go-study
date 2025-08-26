package main

import (
	"fmt"
)

/*
Go 支持组合， 这是将一个结构包含进另一个结构的行为。在某些语言中，这种行为叫做 特质 或者 混合
*/

type Person2 struct {
	Name string
}

func (p *Person2) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

// Saiyan 结构体有一个 Person2 类型的字段。由于我们没有显式地给它一个字段名，所以我们可以隐式地访问组合类型的字段和函数。同时也可以用.Person2访问到
type Saiyan2 struct {
	*Person2 // composition Person2 with &Person2.
	Power    int
}

func maincomposition() {
	goku5 := &Saiyan2{
		Person2: &Person2{"Goku5"},
		Power:   9001,
	}
	goku5.Introduce() //Hi, I'm Goku5

	goku6 := &Saiyan2{
		Person2: &Person2{"Goku6"},
	}
	fmt.Println(goku6.Name)         // Goku6
	fmt.Println(goku6.Person2.Name) // Goku6

}
