package lib

import (
	"log"
	"strings"
)

//SyntaxError : Custom Object for Syntax
type SyntaxError struct {
	Msg string // description of error
}

//LexString :  Returns the the string type
// a valid string will start from  " and end with"
func LexString(str string) (string, string, SyntaxError) {

	var curStr string
	if string(str[0]) == string(QUOTE) {
		str = str[1:]
		for i, c := range str {

			if string(c) == string(QUOTE) {
				return strings.TrimSpace(curStr), str[i+1:], SyntaxError{Msg: ""}
			}
			curStr += string(c)
		}

		return "", "", SyntaxError{Msg: "End quote not found"}

	}

	return "", str, SyntaxError{}

}

//LexInt : lexes the interget part
func LexInt(str string) (string, string, SyntaxError) {

	validNumbers := "0123456789-E."
	curInt := ""
	for i, char := range str {
		curChar := string(char)
		if !strings.Contains(validNumbers, curChar) {
			return curInt, str[i:], SyntaxError{}
		}
		curInt += curChar
	}

	return "", str, SyntaxError{}
}

//LexBoolOrNull :  returns the boolean values
func LexBoolOrNull(str string) (string, string, SyntaxError) {
	trueStr := "true"
	falseStr := "false"
	nullStr := "null"

	if len(str) < 4 {
		return "", str, SyntaxError{}
	}

	//first 4 characters are true
	if len(str) >= 4 && str[:4] == trueStr {
		return trueStr, str[4:], SyntaxError{}
	}

	//first 5 characters are false
	if str[:5] == falseStr {
		return falseStr, str[5:], SyntaxError{}
	}

	//first 4 characters are
	if str[:4] == nullStr {
		return nullStr, str[4:], SyntaxError{}
	}
	return "", str, SyntaxError{}
}

//Lex : performs lexical analysis on the input string
func Lex(input string) ([]string, SyntaxError) {
	ln := len(input)
	tokens := []string{}

	var lexed string
	var err SyntaxError

	for ln > 0 {

		//remove whitespaces
		if strings.Contains(string(WHITESPACE), string(input[0])) {
			input = input[1:]
			ln = len(input)
			continue
		}

		//include symbols as tokens
		if strings.Contains(string(SYMBOLS), string(input[0])) {
			tokens = append(tokens, string(input[0]))
			input = input[1:]
			continue
		}

		lexed, input, err = LexString(input)
		if lexed != "" {
			tokens = append(tokens, lexed)
			ln = len(input)
			continue
		}
		if err.Msg != "" {
			log.Fatal(err.Msg)
			return tokens, err
		}

		lexed, input, err = LexInt(input)
		if lexed != "" {
			tokens = append(tokens, lexed)
			ln = len(input)
			continue
		}
		if err.Msg != "" {
			log.Fatal(err.Msg)
			return tokens, err
		}

		lexed, input, err = LexBoolOrNull(input)
		if lexed != "" {
			tokens = append(tokens, lexed)
			ln = len(input)
			continue
		}
		if err.Msg != "" {
			log.Fatal(err.Msg)
			return tokens, err
		}

		tokens = append(tokens, lexed)
		if lexed != "" {
			tokens = append(tokens, lexed)
			ln = len(input)
			continue
		}

		//code should not reach here

	}
	return tokens, SyntaxError{}

}
