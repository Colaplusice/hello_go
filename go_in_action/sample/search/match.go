package search

import (
	"log"
	"fmt"
)

//支持不同匹配器接口

//保存结果
type Result struct {
	Field   string
	Content string
}

//接受一个feed类型 和string 作为参数
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

//为每个数据源单独启动goroutine来执行这个函数
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}
	//将结果写入通道
	for _, result := range searchResults {
		results <- result
	}

}

//在终端窗口输出
func Display(results chan *Result) {
	for result := range results {
		fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
