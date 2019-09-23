package main

import (
	"fmt"
	"time"
)

func CalFuncExecTime(start time.Time)  {
	elapsed := time.Since(start)

	fmt.Println("运行时间是 ", elapsed)
}

func main() {
	defer CalFuncExecTime(time.Now())

	time.Sleep(1*time.Second)
}
