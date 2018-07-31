package main

import (
	"fmt"
	"github.com/leb9882/stringutil"
)

type T struct {
	name  string // name of the object
	value int    // its value
}

func main() {
	fmt.Printf(stringutil.Reverse("hello, world3\n"))
}
