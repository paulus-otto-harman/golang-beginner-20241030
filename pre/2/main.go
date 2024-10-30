package main

import (
	"fmt"
	"reflect"
)

func coba(tes interface{}) {
	//
	fmt.Println(tes)
	fmt.Println(reflect.TypeOf(tes))
}

func main() {
	coba(struct {
		name  string
		phone string
	}{})
}
