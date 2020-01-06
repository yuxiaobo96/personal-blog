package main

import "fmt"

func main(){
	var age int
	fmt.Println("My age is",age)
	age = 28
	fmt.Println("My age is",age)
	var age2 int = 12
	fmt.Println("My age is",age2)
	var age3 = 89
	fmt.Println("you age is",age3)
	var long, wide, height = 1, 2, 3
	fmt.Println("long is", long, "wide is",wide, "height is", height)
	big,small := 10,"young"
	fmt.Println("big and small are", big,small)
	var (
		animal = "dog"
		color = "black"
		age4 = 3
		height2 = "30cm"
		other string
	)
	fmt.Println("I have a",animal,"its color is", color, "and age is", age4, "and height is", height2, "others",other)
// 测试 := 的可用性，在 := 左边只能声明并赋值新变量，
// 但当在左边至少有一个新声明的变量时，也可以对之前已赋值的变量进行重新声明并赋值，如30行的b
	a, b := 1, 2
	fmt.Println(a,b)
	b, c := 3, 4
	fmt.Println(b, c)
	a, c, d := 5, 6, 7
	fmt.Println(a, c, d)
}

