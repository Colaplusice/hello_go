# hello go


## 语法


判断dict key 是否存在
if _, ok := map1[key1]; ok {
// ...
}

while  用for代替

int to string 	y:=fmt.Sprintf("%d",x)  	print(strconv.Itoa(a))

类型判断  t := reflect.TypeOf(3) // a reflect.Type

遍历字符串
	previous := "123"
	for i := 0; i < len(previous); i ++ {
		fmt.Printf("%q",previous[i])
	}

字符串复制:
	fmt.Println(strings.Repeat(a,2))

slice