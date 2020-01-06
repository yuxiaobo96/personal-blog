package main

import "fmt"

func main() {
	// break 是终止循环，并跳出循环，执行循环外的语句（即在符合条件后，终止循环）
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break
		}
		fmt.Println("value of i is", i)
	}
	fmt.Println("test the break")
	// continue 是不执行符合条件的当前循环，执行下一循环，当i=6时，i>5符合条件，执行continue，即不执行当前循环，所以不输出
	for i := 1; i <= 10; i++ {
		if i > 5 {
			continue
		}
		fmt.Println("value of i is", i)
	}
	fmt.Println("test the continue")
	n := 5
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Println("value of i,j is", i, j)
		}
	}

	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
			if i == j {
				break // 当i=j=1时，break打破当前的内循环，但是外循环还是会执行，所以依然会输出i=2,并继续执行内循环，输出j=1
			}
		}

	}
	// 可以使用标签直接破坏外部循环
outer:
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("\ni = %d , j = %d\n", i, j)
			if i == j {
				break outer // 当i=j=i时，直接破坏外部循环，那么程序将不再执行，退出
			}
		}

	}
}

