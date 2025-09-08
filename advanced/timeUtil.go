package main

import "time"

var (
	cst *time.Location
)

// CSTLayout China Standard Time Layout
const CSTLayout = "2006-01-02 15:04:05"

func init() {
	var err error
	if cst, err = time.LoadLocation("Asia/Shanghai"); err != nil {
		panic(err)
	}
}

//  convert rfc3339 value to china standard time layout
func RFC3339ToCSTLayout(value string) (string, error) {
	ts, err := time.Parse(time.RFC3339, value) // 指定时间格式
	if err != nil {
		return "", err
	}

	return ts.In(cst).Format(CSTLayout), nil
}
