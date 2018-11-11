package main

import "fmt"

func main() {
for i:=0;i<23;i++{
	fmt.Print("-")
	}
	
	println("循环完成")
	i:=23
	if i>1{
		print("这是i大于1")
	}else {
		println("else 必须接在后面")
	}
	d:=2
	switch d {
	case 1:
		fmt.Print("false")
	case 2:
		fmt.Print("right")
	case 3:
		println("error")

	}

}