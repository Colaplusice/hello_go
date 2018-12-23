# hello go

## 语法

### map

判断dict key 是否存在
if _, ok := map1[key1]; ok {
// ...
}

while  用for代替

int2string 	`y:=fmt.Sprintf("%d",x) print(strconv.Itoa(a))`
string2int `	a,_:=strconv.Atoi("123")`

类型判断  t := reflect.TypeOf(3) // a reflect.Type
遍历字符串
```
previous := "123"
for i := 0; i < len(previous); i ++ {
	fmt.Printf("%q",previous[i])
}
```

字符串index
 `fmt.Println(string([]rune("Hello, 世界")[8]))`
`fmt.Println(string("123"[1]))`

字符串复制:
`fmt.Println(strings.Repeat(a,2))`

### 初始化map

slice

向上取整：
官方的math 包中提供了取整的方法，向上取整math.Ceil() ，向下取整math.Floor()


n := map[string]int{"foo": 1, "bar": 2}
删除
delete(m, "k2")

## goland 编译同一个package下的所有文件
