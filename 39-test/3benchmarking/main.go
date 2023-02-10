package main

import (
	"fmt"

	"github.com/AmitSuresh/learning-golang/39-test/3benchmarking/speak"
)

func main() {
	str := speak.Hello("hola")
	fmt.Println(str)
}
