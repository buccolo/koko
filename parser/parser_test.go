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

func TestReturnStatement(t *testing.T) {
  input := `
  return 2;
  return foo(420);
  `

  lexer := lexer.New(input)
  parser := New(lexer)

  program := parser.ParseProgram()
  checkParserErrors(t, parser)
  if program == nil {
    t.Fatalf("ParseProgram returned nil")
  }
  if len(program.Statements) != 2 {
    t.Fatalf("program.Statements contains %d statements instead of 3", len(program.Statements))
  }

  for _, stmt := range program.Statements {
    returnStmt, ok := stmt.(*ast.ReturnStatement)
    if !ok {
      t.Errorf("stmt not *ast.ReturnStatement, got=%q", stmt)
    }
    if returnStmt.TokenLiteral() != "return" {
      t.Errorf("returnStmt.TokenLiteral is not return, got=%q", returnStmt.TokenLiteral())
    }
  }
}

func TestParserError(t *testing.T) {
  input := `
  let x 5;
  let = x 5;
  `

  lexer := lexer.New(input)
  parser := New(lexer)
  parser.ParseProgram()

  expectedErrors := []string{
    "expected next token to be =, got INT instead",
    "expected next token to be IDENT, got = instead",
  }

  for i, error := range expectedErrors {
    if error != parser.Errors()[i] {
      t.Fatalf("Expected error '%q', got: '%q'", error, parser.Errors()[i])
    }
  }
}
