package main

import (
	"encoding/json"
	"fmt"
)

//	使用json 字符串结构定义struct结构体，快捷方法可使用在线工具：https://mholt.github.io/json-to-go/。
//
// 对于固定字段的键值对，不要使用 map[string]interface{}
// go中只有首字母大写的标识符（变量、函数、结构体字段等）才是“导出的（Exported）”，才能被其他包访问。
type Response struct {
	Resultcode int `json:"resultcode"` //结构体字段名必须大写开头（即“导出”）才能被 encoding/json 包正确访问和解析
	Number     int `json:"number"`
}

func mainJson1() {
	jsonStr := `
		{
			"resultcode": 200,
			"number":1234567
		}
	`

	var result Response //用Response struct接收unmarshal自动赋值
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result.Resultcode) // 200
	fmt.Println(result.Number)     // 1234567
}

func mainJson2() {
	jsonStr := `
        {
            "resultcode": 200,
			"number":1234567
        }
    `

	var result map[string]interface{} //用map接收unmarshal后json子目录未知要用["resultcode"]
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(result["resultcode"])
	fmt.Println(result["number"])                // 1.234567e+06 map[string]interface{} 来接收反序列化结果时，大于6为变成科学计数法
	fmt.Println(int(result["number"].(float64))) //1234567 强制类型转换

	//1. 使用WeakDecode
	// var mobile Response
	// err = mapstructure.WeakDecode(result, &mobile)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(mobile.Resultcode)

	// 2. 获取 resultcode 并转换为 int
	var resultcode int
	if val, exists := result["resultcode"]; exists {
		switch v := val.(type) {
		case float64: // JSON 数字默认会被解析为 float64
			resultcode = int(v)
		case int:
			resultcode = v
		case json.Number:
			if i, err := v.Int64(); err == nil {
				resultcode = int(i)
			}
		}
	}

	fmt.Println(resultcode)
}
