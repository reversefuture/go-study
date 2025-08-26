package main

// 用接口定义行为
type Speaker interface {
	Speak() // 无返回类型的方法
}

// 只要实现了 Speak() 方法，就自动满足 Speaker 接口，无需显式声明“继承”
func MakeSound(s Speaker) {
	s.Speak()
}

func mainInterfacePolymorphism() {
	animal := &Animal{Name: "Cat", Age: 2}
	dog := &Dog{
		Animal: Animal{Name: "Buddy", Age: 3},
		Breed:  "Husky",
	}

	MakeSound(animal) // I am Cat, I am 2 years old.
	MakeSound(dog)    // I am Buddy the dog, woof!
}
