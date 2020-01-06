package main

import "fmt"

func main(){
	sum := calculate(10,7)
	fmt.Println("单价乘以数量等于总价",sum)
	//因为此处调用的函数有两个返回值，所以声明变量去接返回值时也需要声明两个变量
	chinese, foreign := name("ma","xiaolong")
	fmt.Println("chinese name is",chinese, "foreign name is", foreign)
	// 当只想返回一个值时，可以用下划线_来表示不用返回的值
	chinese2,_ := name("zhang","ling")
	fmt.Println("He is",chinese2)
	area, perimeter := math(3,6)
	fmt.Printf("area and perimeter are %f, %f", area,perimeter)
}

func calculate(price int,amount int)int {
	sum := price * amount
	return sum
}

func name(first, last string)(string,string ){
	chinesename := first + last
	foreignname := last + first
	return chinesename, foreignname
}

// 在函数中命名返回值，相当于将其声明为变量，可以直接赋值，不用再次声明
func math(length, width float64)(area, perimeter float64){
	area = length * width
	perimeter = (length + width) * 2
	return
}