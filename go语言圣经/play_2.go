package main

import (
	"net"
	"log"
	"time"
	"fmt"
	"strings"
	"bufio"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	// 返回输入结果的大写值
	fmt.Fprintf(c, "\t", strings.ToUpper(shout))

}

func handleconn(c net.Conn) {
	// 读取输入 输出
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)
	}

}

func main() {

	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy()
}
