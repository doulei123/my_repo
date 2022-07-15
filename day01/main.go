package main // 声明 main 包，表明当前是一个可执行程序
import (
	"fmt"
)

// 全局变量m 声明了可以不使用  局部变量申明了必须使用
var m bool
var l string = "dssd"

const (
	n1 = iota //0
	n2 = 100  //100
	n3 = iota //2
	_
	n4 //4
)
const n5 = iota //0

func main() {
	aa := 1
	bb := 2
	aa, bb = bb, aa
	fmt.Println(aa, bb)

	///////////////////////////
}
