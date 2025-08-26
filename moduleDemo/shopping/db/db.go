package db // 包名和文件夹名一致

type Item struct {
	Price float64
}

func LoadItem(id int) *Item {
	return &Item{
		Price: 9.001,
	}
}
