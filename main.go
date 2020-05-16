package main

import (
	"fmt"
	"log"

	"github.com/rdrahul/jsonparser/lib"
)

func main() {

	var str1 string = `{ "name" : "rahul" ,"working" : [ 1,2,3,4 ]  } `
	// str1 := `"read" : "never"`
	// str1 := ` { "read": 1234 } `
	tokens, err := lib.Lex(str1)
	if err.Msg != "" {
		log.Fatal(err.Msg)
	}
	for _, c := range tokens {
		fmt.Printf("%s | ", c)
	}

}
