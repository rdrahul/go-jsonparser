package main

import (
	"log"

	"github.com/rdrahul/jsonparser/lib"
)

func main() {

	// var str1 string = `{ "name" : "rahul" ,"working" : [ 1,2,3,4 ]  } `
	// var str1 string = `{ "name" : "rahul" , "working" : { "nested" : true  } } `
	var str1 string = `{ "name" : "rahul" , "working" : null } `
	// str1 := `{"read" : "never"}`
	// str1 := ` { "read": 1234 } `
	tokens, err := lib.Lex(str1)
	if err.Msg != "" {
		log.Fatal(err.Msg)
	}

	lib.Parser(tokens)

}
