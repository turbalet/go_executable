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
		fmt.Print("Changed to upper case")
	}

	
}
