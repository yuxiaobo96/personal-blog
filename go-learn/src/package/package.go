package main

import (
	"fmt"
	"github.com/yuxiaobo96/personal-blog/go-learn/src/package/app"
	_ "github.com/yuxiaobo96/personal-blog/go-learn/src/package/app"//可以用下划线_来代表当前没有用到的引用包，阻止报错
)

func main() {
	areasum := app.Sumarea(5,2,3)
	fmt.Println("areasum value is", areasum)
}