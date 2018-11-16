package parser

import (
	"testing"

	"es3/ast"
	"es3/lexer"
)

func TestVarStatements(t *testing.T) {
	input := `
	var x = 5;
	var y = 10;
	var $_foobar = 34692;
	`

	lex := lexer.New(input)
	parse := New(lex)

	program := parse.ParseProgram()
	checkParserErrors(t, parse)
	
	if program == nil {
		t.Fatalf("have (nil), want (program)")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("var statements should have 3 components. got %d statements", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"$_foobar"},
	}

	for i, testToken := range tests {
		stmt := program.Statements[i]
		if !testVarStatement(t, stmt, testToken.expectedIdentifier) {
			return
		}
	}
}

func TestReturnStatements(t *testing.T) {
	input := `
	return 5;
	return x;
	return 3 + 2;
	`

	lex := lexer.New(input)
	parse := New(lex)

	program := parse.ParseProgram()
	checkParserErrors(t, parse)

	if len(program.Statements) != 3{
		t.Fatalf("have %d want 3 program statements", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		//type assertion; statement should be a pointer to a ReturnStatement
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("have %T want type *ast.ReturnStatement", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("have literal %q want literal 'return'", returnStmt.TokenLiteral())
		}
	}
}

func checkParserErrors(t *testing.T, parse *Parser) {
	errors := parse.Errors()

	if len(errors) == 0 {
		return 
	}

	t.Errorf("found %d parser errors:", len(errors))

	for _, msg := range errors {
		t.Errorf("%q", msg)
	}
	t.FailNow()
}

func testVarStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "var" {
		t.Errorf("expected var statment got literal %q", s.TokenLiteral())
		return false
	}

	varStmt, ok := s.(*ast.VarStatement)

	if !ok {
		t.Errorf("expected type *ast.VarStatement, got %T", s)
		return false
	}

	if varStmt.Name.Value != name {
		t.Errorf("expected identifier value  %s, got %s", name, varStmt.Name.Value)
		return false
	}

	if varStmt.Name.TokenLiteral() != name {
		t.Errorf("expected statement.Name.TokenLiteral() %s, got %s", name, varStmt.Name)
		return false
	}

	return true
}
