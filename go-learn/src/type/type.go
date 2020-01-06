package main

import (
	"fmt"
	"unsafe"
)

func main(){
	var a = 12
// 在使用格式符打印需要变量的格式时，用 fmt.Printf
	fmt.Printf("type of a is %T, size of a is %d\n",a, unsafe.Sizeof(a))
	fmt.Println("type of a is",a)
	b, c := 5, 6
	sum := b + c
	fmt.Printf("type of sum is %T and sum is %d", sum,sum)

	b, c = 7, 8
	diff := b - c
	fmt.Println("value of diff is", diff)
	var first, last = "yu", "xiaobo"
	name := first + last
	fmt.Println("my name is",name)
// 类型转换
	var (
		first2 = 3
		last2 = 3.9
	)
	last3 := float64(first2) //强制类型转换，将last2转换为int类型，并赋值给新 声明的变量last3
	name2 := last2 + last3
	fmt.Println("his name is", name2)
}
