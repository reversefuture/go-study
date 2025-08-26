package main

import (
	"fmt"
)

/*
go without OPP.
While arrays are used to store multiple values of the same data type into a single variable, structs are used to store multiple values of different data types into a single variable.
字段可以是任何类型：
type Saiyan struct {   Name string   Power int   Father *Saiyan }
然后我们通过下面的方式初始化：
gohan := &Saiyan{   Name: "Gohan",   Power: 1000,   Father: &Saiyan {     Name: "Goku",     Power: 9001,     Father: nil,   }, }

尽管缺少构造器，Go 语言却有一个内置的 new 函数，使用它来分配类型所需要的内存。 new(X) 的结果与 &X{} 相同。
goku := new(Saiyan)
// same as
goku := &Saiyan{}
*/

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func mainStruct() {
	var pers1 Person

	// Pers1 specification
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000

	// Pers2 specification
	pers2 := Person{
		name:   "Cecilie",
		age:    24,
		job:    "Marketing",
		salary: 4500,
	}

	pers3 := Person{} // properties will have default values of each type
	pers4 := Person{"bob", 22, "Tester", 9000}

	// Access and print Pers1 info
	fmt.Println("Name: ", pers1.name)
	fmt.Println("Age: ", pers1.age)
	fmt.Println("Job: ", pers1.job)
	fmt.Println("Salary: ", pers1.salary)

	// Access and print Pers2 info
	fmt.Println("Name: ", pers2.name)
	fmt.Println("Age: ", pers2.age)
	fmt.Println("Job: ", pers2.job)
	fmt.Println("Salary: ", pers2.salary)
	// Print Pers1 info by calling a function
	printPerson(pers1)
	printPerson(pers2)
	printPerson(pers3)
	printPerson(pers4)
}

func printPerson(pers Person) {
	fmt.Println("Name: ", pers.name)
	fmt.Println("Age: ", pers.age)
	fmt.Println("Job: ", pers.job)
	fmt.Println("Salary: ", pers.salary)
}

// 结构体没有构造器。但是，你可以创建一个返回所期望类型的实例的函数
func NewSaiyan(name string, power int) *Saiyan {
	return &Saiyan{
		Name:  name,
		Power: power,
	}
}

// 我们的工厂不必返回一个指针；下面的形式是完全有效的：
func NewSaiyan2(name string, power int) Saiyan {
	return Saiyan{
		Name:  name,
		Power: power,
	}
}
