package main

import "fmt"

// 定义一个基础结构体（父类）
type Animal struct {
	Name string
	Age  int
}

func (a *Animal) Speak() {
	// a.Animal.Speak() // 调用“父类”方法
	fmt.Printf("I am %s, I am %d years old.\n", a.Name, a.Age)
}

// 重写父类方法：  dog.Speak() // 调用的是 Dog 的 Speak，而不是 Animal 的
func (d *Dog) Speak() {
	fmt.Printf("I am %s the dog, woof!\n", d.Name)
}

// Dog 继承 Animal 的属性和方法
type Dog struct {
	Animal // 匿名嵌入，相当于继承
	Breed  string
}

func mainCompositionInherit() {
	dog := Dog{
		Animal: Animal{Name: "Buddy", Age: 3},
		Breed:  "Golden Retriever",
	}

	// 可以直接调用 Animal 的方法
	dog.Speak() // 输出: I am Buddy, I am 3 years old.
	// dog.Animal.Speak()// 也可以

	// 也可以访问嵌入字段
	fmt.Println(dog.Name) // Buddy
}

// I am Buddy, I am 3 years old.
// Buddy

func mainCompositionOverwrite() {
	dog := Dog{
		Animal: Animal{Name: "Buddy", Age: 3},
		Breed:  "Golden Retriever",
	}

	dog.Speak() // 调用的是 Dog 的 Speak，而不是 Animal 的
}

// I am Buddy the dog, woof!
