package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	counts:=make(map[string]int)
	//输入流
	input:=bufio.NewScanner(os.Stdin)
	for input.Scan() {
		//调用标准输入读取内容
		counts[input.Text()]++
	}

	for line,n :=range counts{
		if n>1{
			fmt.Printf("%d\t%s\n",n,line)
		}

	}


}
