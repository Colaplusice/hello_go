package main

import "fmt"

func main() {

	var a[2][5]int
	for i:=0;i< len(a);i+=1{
		for j:=0;j<len(a[i]);j++{
			fmt.Print("j")
		}
	}

}