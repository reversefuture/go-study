package main

import (
	"encoding/json"
	"fmt"
)

/*
结构体是将零个或多个任意类型的变量，组合在一起的聚合数据类型，也可以看做是数据的集合。

尽管缺少构造器，Go 语言却有一个内置的 new 函数，使用它来分配类型所需要的内存。 new(X) 的结果与 &X{} 相同。
goku := new(Saiyan)
// same as
goku := &Saiyan{}
*/

type Person struct {
	Name string
	Age  int
}

func mainStruct1() {
	var p1 Person
	p1.Name = "Tom"
	p1.Age = 30
	fmt.Println("p1 =", p1)

	var p2 = Person{Name: "Burke", Age: 31}
	fmt.Println("p2 =", p2)

	p3 := Person{Name: "Aaron", Age: 32}
	fmt.Println("p2 =", p3)

	//匿名结构体
	p4 := struct {
		Name string
		Age  int
	}{Name: "匿名", Age: 33}
	fmt.Println("p4 =", p4)
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

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

// 生成 JSON
func mainJson1() {
	var res Result
	res.Code = 200
	res.Message = "success"

	//序列化
	jsons, errs := json.Marshal(res)
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}
	fmt.Println("json data :", string(jsons))

	//反序列化
	var res2 Result
	errs = json.Unmarshal(jsons, &res2)
	if errs != nil {
		fmt.Println("json unmarshal error:", errs)
	}
	fmt.Println("res2 :", res2)
}

// json data : {"code":200,"msg":"success"}
// res2 : {200 success}

// 改变数据
func mainJosn2() {
	var res Result
	res.Code = 200
	res.Message = "success"
	toJson(&res)

	setData(&res)
	toJson(&res)
}

func setData(res *Result) { // 修改json必须要指针！
	res.Code = 500
	res.Message = "fail"
}

func toJson(res *Result) {
	jsons, errs := json.Marshal(res) // 序列化
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}
	fmt.Println("json data :", string(jsons))
}

// json data : {"code":200,"msg":"success"}
// json data : {"code":500,"msg":"fail"}
