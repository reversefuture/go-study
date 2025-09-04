package main

import "fmt"

func mainAssert() {
	// 示例1: a == b
	a := []byte("hello")
	b := []byte("hello")
	fmt.Printf("Compare(%q, %q) = %d\n", a, b, Compare(a, b)) // 输出: 0

	//配合 `switch` 使用,switch 独有
	var t interface{}
	t = Compare(a, b)
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t) // %T 打印任何类型的 t
	case bool:
		fmt.Printf("boolean %t\n", t) // t 是 bool 类型
	case int:
		fmt.Printf("integer %d\n", t) // t 是 int 类型
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t 是 *bool 类型
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t 是 *int 类型
	}

	// 从 `interface{}` 中取出具体值
	// var x interface{} = "hello"
	// v := x.(int) // 崩溃，panic: interface is string, not int
	// fmt.Println("v: ", v)

	// printValue("adf") // 不会崩溃

	data := map[string]interface{}{
		"name":   "Alice",
		"age":    30,
		"active": true,
	}

	// 与 `map[string]interface{}` 配合（如解析 JSON）
	name := data["name"].(string) // 已知是 string
	age, ok := data["age"].(int)  // 安全断言
	if ok {
		fmt.Println(name, "年龄:", age)
	}

}

// 比较两个字节型切片，返回一个整数
// 按字典顺序.
// 如果a == b，结果为0；如果a < b，结果为-1；如果a > b，结果为+1
func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	switch {
	case len(a) > len(b):
		return 1
	case len(a) < len(b):
		return -1
	}
	return 0
}

func printValue(v interface{}) {
	if str, ok := v.(string); ok { // 不会crash的类型断言
		fmt.Println("字符串:", str)
	} else if n, ok := v.(int); ok {
		fmt.Println("整数:", n)
	} else {
		fmt.Println("未知类型")
	}
}
