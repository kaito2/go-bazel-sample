package main

import (
	"fmt"
	"github.com/kaito2/go-bazel-sample/goland/lib/example"
	"github.com/rs/xid"
)

func main() {
	c := example.Client{}
	fmt.Println(c.Message())
	fmt.Println(xid.New())
}
