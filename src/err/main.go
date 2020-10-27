package main

import (
	"customErr"
	"errors"
	"fmt"
)

func main() {
	err := customErr.CustomError{Name: "自定义错误", Msg: "错误信息"}

	fmt.Println(err.Msg)
	//a := -1
	//b := -1

	//rs, err := add(&a, &b)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(rs)
	//}
	// HelloWorld
}

func add(a, b *int) (c int, err error) {
	if (*a < 0 || *b < 0) {
		err = errors.New("不能为负数")
		return
	}

	*a *= 2
	*b *= 3
	c = *a + *b
	return
}
