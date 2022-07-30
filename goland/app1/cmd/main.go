package main

import (
	"fmt"
	"github.com/kaito2/go-bazel-sample/goland/lib/example"
)

func main() {
	c := example.Client{}
	fmt.Println(c.Message())
}
