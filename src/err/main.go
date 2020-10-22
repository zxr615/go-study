package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("error info")

	if err != nil {
		fmt.Println(err)
	}
}
