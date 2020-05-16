# JSONPARSER
A simple json lexical analysis and parsing done with go

## lexical analysis
lexing or tokenization is the process of converting a sequence of characters into a sequence of tokens.
Here we read the string and then convert them into tokens

## Syntactic analysis or Parser
the job of the syntactic analyzer will be to iterate over the tokens received after a call to lex and try to match the tokens to objects, lists, or plain values.

TODO :
- validate json
- return actual token instead of just string
- implement something like bindjson 
- encode a object to json