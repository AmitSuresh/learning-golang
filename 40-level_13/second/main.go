package main

import (
	"fmt"

	"github.com/AmitSuresh/learning-golang/40-level_13/second/quote"
	"github.com/AmitSuresh/learning-golang/40-level_13/second/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))
	fmt.Println(word.UseCount(quote.SunAlso))
}
