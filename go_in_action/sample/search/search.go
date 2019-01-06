package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

func Run(searchTerm string) {
	//拿到所有的序列化类型?
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}
	// create an channel receive results
	results := make(chan *Result)

	//set up a wait group
	var waitGroup sync.WaitGroup
	//set the number of goroutines we need to wait
	// wait group is a counting semaphore and we will use it
	//to count off goroutines as they finish their work
	waitGroup.Add(len(feeds))

	// 看是feed否在匹配其中存在
	for _, feed := range feeds {
		matcher, exists := matchers[feed.TYPE]
		//如果不存在就选择默认的
		if !exists {
			matcher = matchers["default"]
		}
		//闭包函数，可以直接使用value
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			// 当match 方法完成后，减少信号量的value
			waitGroup.Done()
			//	matcher feed 作为值传递得到匿名函数中
		}(matcher, feed)
	}
	// 运行个goroutine 等所有工作都结束，然后关闭
	// 让goroutine 阻塞直到wait group减到0
	go func() {
		waitGroup.Wait()
		close(results)
	}()
	//显示结果
	Display(results)
}

func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}
	log.Println("register", feedType, "matcher")
	matchers[feedType] = matcher
}
