package expr

import (
	"fmt"
	"strconv"
)

type parser struct {
	l      *lexer
	errors []error
}

func (p *parser) parse() expression {
	tok := p.l.nextToken()

	var exp expression
	switch tok._type {
	case t_NUM:
		n, err := strconv.ParseInt(tok.literal, 10, 64)
		if err != nil {
			p.errors = append(p.errors, fmt.Errorf(`unable to parse integer "%v", %v`, tok.literal, err.Error()))
		}
		if precedences[p.l.readToken()._type] > precedences[tok._type] {
			exp = p.parseInfix(integer(n))
		} else {
			exp = integer(n)
		}
	}

	if tok = p.l.nextToken(); tok._type != t_EOF {
		p.errors = append(p.errors, fmt.Errorf(`unexpected token "%v", expected "EOF"`, tok.literal))
	}

	return exp
}

func (p *parser) parseInfix(left expression) expression {
	tok := p.l.nextToken()
	var in infix
	in.left = left
	in.operator = tok
	in.right = p.parse()
	return in
}
