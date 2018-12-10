# hello go

## 语法

### map

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

### 初始化map

n := map[string]int{"foo": 1, "bar": 2}
删除
delete(m, "k2")



## goland 编译同一个package下的所有文件
