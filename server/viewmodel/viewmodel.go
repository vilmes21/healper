package viewmodel

type JsonRes struct {
	Ok bool `json:"ok"`
	Id int `json:"id"`
	Msg string `json:"msg"`
}

type Src struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Author string `json:"author"`
}

type Cat struct {
	Id int `json:"id"`
	CatTypeId int `json:"catTypeId"`
	Name string `json:"name"`
	Namezh string `json:"namezh"`
	Pinyin string `json:"pinyin"`
}