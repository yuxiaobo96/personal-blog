package main

import (
	"fmt"
	"math"
)

// 常量即赋予一常数固定值
func main(){
	var a = 5
	fmt.Println("value of a is",a)
	a =6
	fmt.Println("value of a is",a)// 由此可见，声明的变量的值是可以重新赋值的
	const b = 7
	fmt.Println("const value of b is",b)// b已经是变量，不能二次赋值

	var c = math.Sqrt(0.16)
	fmt.Println("c is", c)
//	const d  = math.Sqrt(4)
//	fmt.Println("d is",d)
// 常量是在编译前就需要定义的，
// 而math.Sqrt是运行时计算出来的结果，所以不能赋值给常量
	const hello  = "hello"
	fmt.Printf("type of hello is %T\n", hello)
	var defaultname = "wiki"
	newname := defaultname
	fmt.Println("newname is",newname)
	type defaultstring string
	var newname2 defaultstring = "2"
	fmt.Println("newname2 is", newname2)
// 常量声明后可以赋值给其他类型的变量，而不需要强制类型转换，如：
    const assige = 2
    var intVar int = assige
    var int32Var int32 = assige
    var float64Var float64 = assige
    fmt.Println("these values are",intVar, int32Var, float64Var)
    var assign2 = 3
    var replace float64 = float64(assign2)//强制类型转换
    fmt.Println("replace is",replace)
// 数值表达式中，常量可以自由组合计算，没有类型的障碍，
// 但当它们被分配给变量或需要类型的代码时需要使用类型转换
    var sum = 5*5 + 2.4/2
    fmt.Println("value of sum is",sum)
}
