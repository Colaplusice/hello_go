package main

import (
	"os"
	"bufio"
	"fmt"
)

func countlines(f *os.File, counts map[string]int) {
	//读取文件
	input := bufio.NewScanner(f)
	for input.Scan() {
		print(len(counts))
		counts[input.Text()]++
	}
	print(counts)
}

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	//如果文件为0，将控制台输入作为标准输入流
	if len(files) == 0 {
		//将 counts map作为参数传到统计行数的
		countlines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			//如果有错误的话，打印错误
			if err != nil {
				fmt.Print(os.Stderr, "dumps2： %v\n", err)
				continue
			}
			//统计行数
			countlines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}

	}

}
