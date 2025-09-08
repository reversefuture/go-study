package main

import (
	"fmt"
	"sync"
)

var _ Friend = (*friend)(nil) //预检查friend是否实现Friend

// 创建一个对象池（Object Pool），用于高效复用 *option 结构体实例，避免频繁的内存分配和垃圾回收（GC）压力。
var (
	cache = &sync.Pool{ // Go 标准库提供的一个并发安全的对象池，用于缓存和复用临时对象。
		New: func() interface{} { // 在每次 GC 之前，会清空池中对象，所以准备好 New 函数兜底
			return &option{sex: 0} //  初始化 sex 为 0，其他字段为零值（int 是 0，string 是 ""）
		},
	}
)

type Option func(*option)

type option struct {
	sex    int
	age    int
	height int
	weight int
	hobby  string
}

func (o *option) reset() {
	o.sex = 0
	o.age = 0
	o.height = 0
	o.weight = 0
	o.hobby = ""
}

// 借 → 用 → 重置 → 还
func getOption() *option { // 从池中取，没有就 New 一个
	return cache.Get().(*option)
}

func releaseOption(opt *option) {
	opt.reset()    // 重置字段，避免残留数据影响下次使用
	cache.Put(opt) // 归还对象到池中，供下次复用
}

type Friend interface {
	Find(keywrods string, option ...Option) (string, error)
}

type friend struct {
	sex    int
	age    int
	height int
	weight int
	hobby  string
}

func (f *friend) Find(where string, options ...Option) (string, error) {
	friend := fmt.Sprintf("从 %s 找朋友\n", where)

	opt := getOption()
	defer func() {
		releaseOption(opt)
	}()

	for _, f := range options {
		f(opt)
	}

	if opt.sex == 1 {
		sex := "性别：女性"
		friend += fmt.Sprintf("%s\n", sex)
	}
	if opt.sex == 2 {
		sex := "性别：男性"
		friend += fmt.Sprintf("%s\n", sex)
	}

	if opt.age != 0 {
		age := fmt.Sprintf("年龄：%d岁", opt.age)
		friend += fmt.Sprintf("%s\n", age)
	}

	if opt.height != 0 {
		height := fmt.Sprintf("身高：%dcm", opt.height)
		friend += fmt.Sprintf("%s\n", height)
	}

	if opt.weight != 0 {
		weight := fmt.Sprintf("体重：%dkg", opt.weight)
		friend += fmt.Sprintf("%s\n", weight)
	}

	if opt.hobby != "" {
		hobby := fmt.Sprintf("爱好：%s", opt.hobby)
		friend += fmt.Sprintf("%s\n", hobby)
	}

	return friend, nil
}

// WithSex setup sex, 1=female 2=male
func (f *friend) WithSex(sex int) Option {
	return func(opt *option) {
		opt.sex = sex
	}
}

// WithAge setup age
func (f *friend) WithAge(age int) Option {
	return func(opt *option) {
		opt.age = age
	}
}

// WithHeight set up height
func (f *friend) WithHeight(height int) Option {
	return func(opt *option) {
		opt.height = height
	}
}

// WithWeight set up weight
func (f *friend) WithWeight(weight int) Option {
	return func(opt *option) {
		opt.weight = weight
	}
}

// WithHobby set up Hobby
func (f *friend) WithHobby(hobby string) Option {
	return func(opt *option) {
		opt.hobby = hobby
	}
}
