package main

import "fmt"

func main(){
	array := [3]int{4,5,9}
	fmt.Println("value of array is", array)
	fmt.Println(array[2])
	//切片：引用类型
	a := []int{1,2,3,4}
	b := a
	fmt.Println("a b",a,b)
	b[0]= 5
	fmt.Println("a b",a,b)
// 数组：值类型，当不指定数组的长度时，应用...代替，（若为空，就是切片了）
	a2 := [...]int{1,2,3,4}
	b2 := a2
	fmt.Println("a b",a2,b2)
	b2[0]= 5
	fmt.Println("a b",a2,b2)
	// 数组在调用函数前和调用函数后的值是相同的，
	// 说明在调用change函数时，数组num2是通过值传递的，本身的数组值不会变
	num2 := [...]int{1, 2, 3}
	fmt.Println("call func before", num2)
	change(num2)
	fmt.Println("current num2",num2)
	//遍历数组
	num6 := [...]int{11,22,33}
	for i:=0;i<len(num6);i++ {
		fmt.Printf("num6[%d]=%d\n",i,num6[i])
	}
	for i,v := range num6{
		fmt.Printf("a[%d]=%d\n",i, v)
	}
	a5 := [...]float64{67.7, 89.8, 21, 78}
	for i := 0; i < len(a5); i++ { //looping from 0 to the length of the array
		fmt.Printf("%d th element of a is %.2f\n", i, a5[i])
	}
}

func change(num [3]int){
	num[0] = 11
	fmt.Println("in change func of num", num)
}
