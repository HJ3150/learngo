package main

import (
	"fmt"

	"github.com/hj3150/learngo/tutorial1/mydict"
)

func main() {
	// dictionary := mydict.Dictionary{"first": "First word"}
	// definition, err := dictionary.Search("first")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }

	// word := "hello"
	// definition = "Greeting"
	// err2 := dictionary.Add(word, definition)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }
	// definition, err = dictionary.Search(word)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("found", word, "definition:", definition)
	// }
	// err2 = dictionary.Add(word, definition)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	err := dictionary.Update("baseWord", "Second")
	if err != nil {
		fmt.Println(err)
	}
	definition2, _ := dictionary.Search(baseWord)
	fmt.Println(definition2)
	err2 := dictionary.Delete("baseWord")
	if err2 != nil {
		fmt.Println(err2)
	}
	def, err3 := dictionary.Search(baseWord)
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println("found", baseWord, "definitiion:", def)
	}
}
