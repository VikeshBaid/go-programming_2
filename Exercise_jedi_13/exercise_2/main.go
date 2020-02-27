package main

import (
	"fmt"
	"github.com/go-programming_2/Exercise_jedi_13/exercise_2/quote"
	"github.com/go-programming_2/Exercise_jedi_13/exercise_2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
