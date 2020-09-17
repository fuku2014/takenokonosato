package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("vim-go")
}

func Foo() string {
	return os.Getenv("FOO")
}
