package main

import (
	_ "embed"
	"fmt"
)

//go:embed sample-image.jpg
var file []byte

func main() {
	// print the file size
	fmt.Println(len(file))
}
