package main

import "fmt"

func n(){
	nameID := 5
	switch nameID {
	case 1:
		fmt.Println("小红")
	case 2:
		fmt.Println("小明")
	case 3:
		fmt.Println("小王")
	case 4, 5, 6:
		fmt.Println("均是教师")
	default:
		fmt.Println("没有此学号的同学")
	}
	num := 75
	switch { // expression is omitted
	case num >= 0 && num <= 50:
		fmt.Println("num is greater than 0 and less than 50")
	case num >= 51 && num <= 100:
		fmt.Println("num is greater than 51 and less than 100")
	case num >= 101:
		fmt.Println("num is greater than 100")
	}

}
