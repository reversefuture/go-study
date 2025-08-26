package main

import (
	"fmt"
	"moduleDemo/shopping" // 注意：这里使用模块路径（同go.mod里面) + 相对路径。 同一个module内导入包总是从module名开始
)

func main() {
	// fmt.Println(shopping.PriceCheck(4343))
	fmt.Println(shopping.PriceCheck2(4343))
}
