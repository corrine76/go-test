package main

import (
	"fmt"
	"reflect"
)

type as struct {
	a int
}

func main() {
	// spider.V3()

	v := &as{}
	fmt.Println(reflect.DeepEqual(v, &as{}))
}
