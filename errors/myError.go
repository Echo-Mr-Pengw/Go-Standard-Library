package main

import(
	"fmt"
	"errors"
)

func main() {

	err := errors.New("错误提示");
	fmt.Println(err);
}