package lib

//exported constants
var (
	QUOTE        byte = '"'
	LEFTBRACKET  byte = '{'
	RIGHTBRACKET byte = '}'
	LEFTBRACE    byte = '['
	RIGHTBRACE   byte = ']'
	COMMA        byte = ','
	COLON        byte = ':'
	WHITESPACE        = []byte{' ', '\t', '\b', '\n', '\r'}
)

//SYMBOLS list of valid json symbols
var SYMBOLS = []byte{LEFTBRACE, RIGHTBRACE, LEFTBRACKET, RIGHTBRACKET, COMMA, COLON}
