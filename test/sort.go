package main

import (
	"fmt"
	"sort"
)

func mainSort() {
	s := []struct {
		id   int
		name string
	}{
		{5, "a"},
		{2, "b"},
		{6, "c"},
		{3, "d"},
		{1, "e"},
		{4, "f"},
	}

	sort.Slice(s, func(i, j int) bool {
		return s[i].id > s[j].id
	})

	fmt.Println(s)
	// [{6 c} {5 a} {4 f} {3 d} {2 b} {1 e}]

	q := Queue{
		{"d", 3},
		{"c", 2},
		{"e", 4},
		{"a", 0},
		{"b", 1},
	}

	fmt.Println(sort.IsSorted(q)) //false
	// sort.Slice(q, func(i, j int) bool {
	// 	return q[i].index > q[j].index
	// })
	sort.Sort(q)
	fmt.Println(q, sort.IsSorted(q)) // [{e 4} {d 3} {c 2} {b 1} {a 0}] true

	//reverse
	sort.Sort(sort.Reverse(q))
	fmt.Println(q) // [{a 0} {b 1} {c 2} {d 3} {e 4}]

	nums := []int{3, 1, 4, 1, 5}
	data := sort.IntSlice(nums) // data 是 IntSlice 类型
	ptr := &data                // ptr 是 *IntSlice

	var _ sort.Interface = data // ❌ 编译错误！IntSlice 没实现 Swap
	var _ sort.Interface = ptr  // ✅ 正确，*IntSlice 实现了全部方法
}

// 下面摘抄自源码
type Data struct {
	text  string
	index int
}

type Queue []Data

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) Less(i, j int) bool {
	return q[i].index > q[j].index
}

func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

type reverse struct { // embed sort的Interface，定义reverse
	sort.Interface
}

func (r reverse) Less(i, j int) bool {
	return r.Interface.Less(j, i)
	// return !r.Interface.Less(i,j ) //同上
}

// 将一个实现了 sort.Interface 的数据包装成一个“按反向顺序比较”的新对象，从而实现降序排序。
func Reverse(data sort.Interface) sort.Interface {
	return &reverse{data}
}
