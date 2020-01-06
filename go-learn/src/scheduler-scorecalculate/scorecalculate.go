package main

import (
	"fmt"
	"math"
)

func main(){
	var totalMilliCPU = 2.0
	var nodeCPUCapacity =4.0
	score :=  math.Abs(totalMilliCPU/nodeCPUCapacity)
	fmt.Println("socre is",score)
}
