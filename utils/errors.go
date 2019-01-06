package utils
//声明errorString结构体
type errorString struct {
	s string
}
//重写Error方法
func (e *errorString) Error() string {
	return e.s
}
//定义New方法来使用
func New(text string) error {
	return &errorString{text}

}
