package main

import (
	"fmt"
	"reflect"
)

func main(){
	var value interface{}
	value = 45.3
	fmt.Println("Tipe data: ",reflect.TypeOf(value))
}