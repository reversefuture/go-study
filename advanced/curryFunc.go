package main

import (
	"fmt"
)

func mainCurry() {
	me := &friend{
		sex:    1,
		age:    1,
		height: 1,
		weight: 1,
		hobby:  "",
	}
	friends, err := me.Find("附近的人",
		me.WithSex(1),
		me.WithAge(30),
		me.WithHeight(160),
		me.WithWeight(55),
		me.WithHobby("爬山"))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(friends)
}

// 从 附近的人 找朋友
// 性别：女性
// 年龄：30岁
// 身高：160cm
// 体重：55kg
// 爱好：爬山
