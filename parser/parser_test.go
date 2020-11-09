package parser

import (
  "testing"
  "koko/ast"
  "koko/lexer"
)

func checkParserErrors(t *testing.T, p *Parser) {
  errors := p.Errors()
  if len(errors) == 0 {
    return
  }

  t.Errorf("parser has %d errors", len(errors))
  for _, msg := range errors {
    t.Errorf("parser error: %q", msg)
  }
  t.FailNow()
}

func TestLetStatement(t *testing.T) {
  input := `
  let x = 5;
  let y = 2;
  let foobar = 123;
  `

  lexer := lexer.New(input)
  parser := New(lexer)

  program := parser.ParseProgram()
  checkParserErrors(t, parser)
  if program == nil {
    t.Fatalf("ParseProgram returned nil")
  }
  if len(program.Statements) != 3 {
    t.Fatalf("program.Statements contains %d statements instead of 3", len(program.Statements))
  }

  tests := []struct {
    expectedIdentifier string
  }{
    {"x"},
    {"y"},
    {"foobar"},
  }

  for i, test := range tests {
    statement := program.Statements[i]
    if !testLetStatement(t, statement, test.expectedIdentifier) {
      return
    }
  }
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
  if s.TokenLiteral() != "let" {
    t.Fatalf("Wrong statement type %q expected 'let'.", s.TokenLiteral())
    return false
  }

  letStmt, ok := s.(*ast.LetStatement)
  if !ok {
    t.Fatalf("Not a LetStatement.")
    return false
  }

  if letStmt.Name.TokenLiteral() != name {
    t.Fatalf("Wrong token literal. expected=%q got=%q", name, letStmt.Name.TokenLiteral())
    return false
  }

  return true
}
