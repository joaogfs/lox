package main

import (
	"fmt"
)

type Scanner struct {
	source  string
	tokens  []Token
	start   int
	current int
	line    int
}

func (s *Scanner) scanTokens() []string {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	// last token should be EOF
	s.tokens = append(s.tokens, Token{
		tokenType: EOF,
		lexeme:    "",
		literal:   "",
		line:      s.line,
	})

	var tktp []string
	for _, v := range s.tokens {
		tktp = append(tktp, v.lexeme)
	}

	return tktp
}

func (s *Scanner) scanToken() {
	char := s.advance()
	// fmt.Printf("current char: %v\n", char)
	// big switc
	switch char {
	case '(':
		s.addToken(LEFT_PAREN)
	case ')':
		s.addToken(RIGHT_PAREN)
	case '{':
		s.addToken(LEFT_BRACE)
	case '}':
		s.addToken(RIGHT_BRACE)
	case ',':
		s.addToken(COMMA)
	case '.':
		s.addToken(DOT)
	case '-':
		s.addToken(MINUS)
	case '+':
		s.addToken(PLUS)
	case ';':
		s.addToken(SEMICOLON)
	case '*':
		s.addToken(STAR)
	case '!':
		if s.match('=') {
			s.addToken(BANG_EQUAL)
		} else {
			s.addToken(BANG)
		}
	case '=':
		if s.match('=') {
			s.addToken(EQUAL_EQUAL)
		} else {
			s.addToken(EQUAL)
		}
	case '<':
		if s.match('=') {
			s.addToken(LESS_EQUAL)
		} else {
			s.addToken(LESS)
		}
	case '>':
		if s.match('=') {
			s.addToken(GREATER_EQUAL)
		} else {
			s.addToken(GREATER)
		}
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(SLASH)
		}
	case ' ':
	case '\r':
	case '\t':
	case '\n':
		s.line++
	default:
		Error(s.line, fmt.Sprintf("Unexpected character: %v", char))
	}

}

func (s *Scanner) advance() byte {
	c := s.source[s.current]
	s.current++
	return c
}

func (s *Scanner) addToken(t TokenType) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, Token{
		tokenType: t,
		lexeme:    text,
		literal:   "",
		line:      s.line,
	})
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}

	if s.source[s.current] != expected {
		return false
	}

	s.current++
	return true

}

func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return '\x00'
	}

	return s.source[s.current]
}
