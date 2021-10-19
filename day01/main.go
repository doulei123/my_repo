package main // 声明 main 包，表明当前是一个可执行程序
import (
	"fmt"
	"math"
) // 导入内置 fmt 包

// 全局变量m 声明了可以不使用  局部变量申明了必须使用
var m bool
var l string = "dssd"

// var m = 100   变量m不支持重复声明

//常量定义  在const关键字出现时将被重置为0。const中每新增一行常量声明将使
const (
	n1 = iota //0
	n2 = 100  //100
	n3 = iota //2
	_
	n4 //4
)
const n5 = iota //0

func foo() (int, string) {
	return 10, "Q1mi"
}
func main() { // main函数，是程序执行的入口
	fmt.Println("Hello World!") // 在终端打印 Hello World!
	//局部变量定义
	n := 10
	t := "sjjsaj"
	tt := t[1:4]
	fmt.Printf("tpye:%T,value:%v", tt, tt)
	// m := 200 // 此处声明局部变量m
	fmt.Println(m, n, t)
	//_为匿名变量，多用于占位，表示忽略值
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)

	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.3f\n", math.Pi)

	//字符串转义
	fmt.Println("str := \"c:\\Code\\lesson1\\go.exe\"")
	// 多行字符串
	s1 := `第一行
                   第二行 
        \n  第三行 /t aaa  
`
	fmt.Println(s1)
	var s5 string
	fmt.Printf("s5:%v\n" , s5)

	slice01 := []int{1,4,2,1,4,6,24,2} 
// 	fmt.Println(copy(slice01[2:],slice01[3:]))
// 	fmt.Println(slice01)
	slice01[2],slice01[3] = slice01[3],slice01[2]  // 交换元素

	s3 := make([]int,2,4)
	
	fmt.Println(slice01,s3)

	aa:= 1
	bb:= 2
	aa,bb = bb,aa
	fmt.Println(aa,bb)
}
