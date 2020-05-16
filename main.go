package main

import (
	"log"

	"github.com/rdrahul/jsonparser/lib"
)

func main() {

	var str1 string = `{ "name is good" : "rahul" ,"working" : [ 1,2,3,4 ]  } `
	// str1 := `"read" : "never"`
	// str1 := ` { "read": 1234 } `
	tokens, err := lib.Lex(str1)
	if err.Msg != "" {
		log.Fatal(err.Msg)
	}

	lib.Parser(tokens)

}
