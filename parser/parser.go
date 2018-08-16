package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	lexer *lexer.Lexer

	currentToken token.Token
	peekToken token.Token
}

func New(lex *lexer.Lexer) *Parser {
	prs := &Parser{lexer: lex}

	prs.nextToken()
	prs.nextToken()

	return prs
}

func (prs *Parser) nextToken() {
	prs.currentToken = prs.peekToken
	prs.peekToken = prs.lexer.NextToken()
}

func (prs *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for prs.currentToken.Type != token.EOF {
		stmt := prs.parseStatement()

		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		prs.nextToken()
	}

	return program
}

func (prs *Parser) parseStatement() ast.Statement {
	switch prs.currentToken.Type {
	case token.LET:
		return prs.parseLetStatement()
	default:
		return nil
	}
}

func (prs *Parser) parseLetStatement() ast.Statement {
	stmt := &ast.LetStatement{Token: prs.currentToken}

	if !prs.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: prs.currentToken, Value: prs.currentToken.Literal}

	if !prs.expectPeek(token.ASSIGN) {
		return nil
	}

	for !prs.currentTokenIs(token.SEMICOLON) {
		prs.nextToken()
	}

	return stmt
}

func (prs *Parser) currentTokenIs(t token.TokenType) bool {
	return prs.currentToken.Type == t
}

func (prs *Parser) peekTokenIs(t token.TokenType) bool {
	return prs.peekToken.Type == t
}

func (prs *Parser) expectPeek(t token.TokenType) bool {
	if prs.peekTokenIs(t) {
		prs.nextToken()
		return true
	} else {
		return false
	}
}
