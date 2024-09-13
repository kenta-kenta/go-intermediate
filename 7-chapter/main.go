package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	_, err0 := strconv.Atoi("a")
	fmt.Printf("err0: [%T] %v\n", err0, err0)

	// err0の中に含まれているエラーを取り出して、err1に代入
	err1 := errors.Unwrap(err0)
	fmt.Printf("err1: [%T] %v\n", err1, err1)

	// err1の中に含まれているエラーを取り出して、err2に代入
	err2 := errors.Unwrap(err1)
	fmt.Printf("err2: [%T] %v\n", err2, err2)
}
