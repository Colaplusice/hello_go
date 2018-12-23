package main

import (
	"fmt"
	"strconv"
	"reflect"
)

func main() {
	fmt.Println(string("123"[1]))
	a,_:=strconv.Atoi("123")
	fmt.Println(reflect.TypeOf(a))

	a=243
	b:=245
	a-=b-2
	fmt.Println(a)
}
