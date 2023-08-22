package main

import (
	"fmt"

	gomodtest "github.com/chatton/go-mod-test"
)

func main() {
	foo := gomodtest.Version
	fmt.Print(foo)

	gomodtest.Foo()
}
