package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	defer MeasureTime(context.TODO(), "aaa")()
	var arr []string
	arr = append(arr, "a")
	arr = append(arr, "b")
	arr = append(arr, "c")
	fmt.Println(arr)
}

func MeasureTime(ctx context.Context, funcName string, extraData ...interface{}) func() {
	start := time.Now()
	return func() {
		dur := time.Since(start)
		fmt.Println(dur.String())
	}
}
