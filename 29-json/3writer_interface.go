package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello")
	fmt.Fprint(os.Stdout, "Hello")
	io.WriteString(os.Stdout, "Hello")
}
