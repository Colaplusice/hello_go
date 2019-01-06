package sample

import (
	"log"
	"os"
	"hello_go/go_in_action/sample/search"
)

//在main函数之前调用
func init()  {
	log.SetOutput(os.Stdout)
	
}
func main() {
	// 搜索总统
 search.Run("president")

}