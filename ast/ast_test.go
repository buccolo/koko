package ast

import (
	"koko/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "testing"},
					Value: "testing",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "someVar"},
					Value: "someVar",
				},
			},
		},
	}

	if program.String() != "let testing = someVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
