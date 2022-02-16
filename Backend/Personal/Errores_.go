package Personal

import "github.com/antlr/antlr4/runtime/Go/antlr"

type CustomSyntaxError struct {
	Line, Column int
	Msg          string
}

type CustomErrorListener struct {
	*antlr.DefaultErrorListener // Embed default which ensures we fit the interface
	Errors                      []CustomSyntaxError
}

func (c *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	c.Errors = append(c.Errors, CustomSyntaxError{
		Line:   line,
		Column: column,
		Msg:    msg,
	})
}
