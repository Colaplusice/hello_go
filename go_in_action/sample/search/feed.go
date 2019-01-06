package search

import (
	"os"
	"encoding/json"
)

// 读取json
const datafile string = "data/data.json"

type Feed struct {
	Name string `json:"site"`
	URL  string `json:"link"`
	TYPE string `json:"type"`
}

//读取并反序列化文件
func RetrieveFeeds() ([]*Feed, error) {
	file, err := os.Open(datafile)
	if err != nil {
		return nil, err
	}
	//当function 完成时文件时关闭的
	defer file.Close()
	//解析文件到Feed struct中
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	return feeds, err
}
