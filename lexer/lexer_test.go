package lexer

import (
	"koko/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
  let five = 5;
  let ten = 10;

  let add = fn(x, y) {
    x + y;
  };

  let result = add(five, ten);

  1 - 2 / 3 * 4;
  1 < 2 > 1;
  if true {
    return !false
  } else {
    0
  }
  0 == 0
  1 != 2
  `

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.INT, "1"},
		{token.MINUS, "-"},
		{token.INT, "2"},
		{token.DIV, "/"},
		{token.INT, "3"},
		{token.MULT, "*"},
		{token.INT, "4"},
		{token.SEMICOLON, ";"},
		{token.INT, "1"},
		{token.LT, "<"},
		{token.INT, "2"},
		{token.GT, ">"},
		{token.INT, "1"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.TRUE, "true"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.BANG, "!"},
		{token.FALSE, "false"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.INT, "0"},
		{token.RBRACE, "}"},
		{token.INT, "0"},
		{token.EQ, "=="},
		{token.INT, "0"},
		{token.INT, "1"},
		{token.NOT_EQ, "!="},
		{token.INT, "2"},
		{token.EOF, ""},
	}

	l := New(input)

	for _, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("Wrong token type. Expected: %q. Got: %q", tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("Wrong literal. Expected: %q. Got: %q", tt.expectedLiteral, tok.Literal)
		}
	}
}
