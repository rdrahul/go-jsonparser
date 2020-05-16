package lib

import "fmt"

//ParseArray : parses the tokens and returns array
func ParseArray(tokens []string) ([]interface{}, int) {
	curArray := []interface{}{}

	//end condition for our recursice function we stop when we find closing brace
	if tokens[0] == string(RIGHTBRACKET) {
		return curArray, 1
	}

	var i int
	//go through all the tokens insert them into the array
	for i = 0; i < len(tokens); i++ {
		token := tokens[i]
		if token == string(RIGHTBRACKET) {
			return curArray, i + 1
		}
		if token == string(COMMA) {
			continue
		}

		if token == string(LEFTBRACE) {
			value, indx := ParseObject(tokens[i:])
			i += indx
			curArray = append(curArray, value)
			continue
		}

		if token == string(LEFTBRACKET) {
			value, indx := ParseArray(tokens[i:])
			i += indx
			curArray = append(curArray, value)
			continue
		}

		curArray = append(curArray, token)

	}
	return curArray, i + 1

}

//ParseObject : returns the object from from the token
func ParseObject(tokens []string) (map[string]interface{}, int) {
	var curobj map[string]interface{} = make(map[string]interface{})

	var curKey string
	var lastValue interface{}
	var indx int
	var i int
	for i = 0; i < len(tokens); i++ {

		token := tokens[i]
		if token == string(RIGHTBRACE) {
			curobj[curKey] = lastValue
			return curobj, i + 1
		}
		if token == string(LEFTBRACE) {
			lastValue, indx = ParseObject(tokens[i+1:])
			i += indx
			continue
		}
		if token == string(LEFTBRACKET) {
			lastValue, indx = ParseArray(tokens[i+1:])
			i += indx
			continue
		}
		if token == string(COLON) {
			curKey = tokens[i-1]
			continue
		}
		//we encountered comma put key value parr
		if token == string(COMMA) {
			curobj[curKey] = lastValue
		}
		//its the last key
		if i == len(tokens)-1 {
			curobj[curKey] = lastValue
		}
		lastValue = token
	}
	return curobj, i
}

//Parser parses the tokens
func Parser(tokens []string) map[string]interface{} {
	result, _ := ParseObject(tokens[1:])
	fmt.Println(result)
	return result
}
