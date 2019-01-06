package search

type defaultMatcher struct {
}

func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// defaultMatcher类型 实现了matcher接口的方法 声明 defaultMatcher类型为接收者
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
