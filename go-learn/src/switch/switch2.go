package main

import "fmt"

func main(){
	switch total := math(5, 5); {
	case total <20:
		fmt.Printf("%d of value lesser than 20\n",total)
	case total <15:
		fmt.Printf("%d of value lesser than 15\n",total)
// switch语句遇到符合条件的case后，打印并退出；
// 当在符合条件的case后添加	fallthrough（位置应是case中的最后一个语句）后，会继续执行下一个case语句（无论case是否符合条件）
// 但fallthrough 生效的前提是fallthrough当前所在的case是符合条件的
		fallthrough
	case total <30:
 		fmt.Printf("%d of value lesser than 30\n",total)
		fallthrough
	case total <200:
		fmt.Printf("%d of value lesser than 200\n",total)
	}
}
func math(a, b int)int{
	math := a * b
	return math
}

