package main

import (
	"fmt"
	"github.com/turbalet/go_lib"
	"unicode"
)

func main() {
	c := 'a'
	c = go_lib.ChangeCase(c)
	if unicode.IsUpper(c) {
		fmt.Println("Changed to upper case")
	}

	x1, x2 := go_lib.FindRoots(10, 15, -25)
	fmt.Println(x1, x2)

	fmt.Println(go_lib.GenerateUUID())
}
