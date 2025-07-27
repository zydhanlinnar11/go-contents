package main

import (
	_ "embed"
	"fmt"
)

//go:embed file-to-import.txt
var file string

func main() {
	fmt.Println(file)
}
