package expr

import (
	"fmt"
	"strconv"
)

type parser struct {
	l         *lexer
	currToken token
	nextToken token
	errors    []error

	prefixParser map[tokenType]func() expression
	infixParser  map[tokenType]func(expression) expression
}

func newParser(l *lexer) *parser {
	p := &parser{
		l: l,
	}
	p.proceed()
	p.proceed()

	p.prefixParser = map[tokenType]func() expression{
		t_NUM: p.parseInteger,
		t_SUB: p.parsePrefix,
	}
	p.infixParser = map[tokenType]func(expression) expression{
		t_ADD: p.parseInfix,
		t_SUB: p.parseInfix,
	}
	return p
}

func (p *parser) parse() expression {
	var exp expression = p.parseExpression()

	if p.currToken._type != t_EOF {
		p.errors = append(p.errors, fmt.Errorf(`unexpected token "%v", expected "EOF"`, p.currToken.literal))
	}

	return exp
}

func (p *parser) parseExpression() expression {
	pp := p.prefixParser[p.currToken._type]
	if pp == nil {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %q", p.currToken._type))
		return nil
	}

	exp := pp()

	ip := p.infixParser[p.currToken._type]
	if ip != nil {
		exp = ip(exp)
	}

	return exp
}

func (p *parser) parseInteger() expression {
	n, err := strconv.ParseInt(p.currToken.literal, 10, 64)
	if err != nil {
		p.errors = append(p.errors, fmt.Errorf(`unable to parse integer "%v"`, p.currToken.literal))
	}
	p.proceed()
	return integer(n)
}

func (p *parser) parsePrefix() expression {
	var pref prefix
	pref.operator = p.currToken
	p.proceed()
	pref.operand = p.parseExpression()
	return pref
}

func (p *parser) parseInfix(left expression) expression {
	var in infix
	in.left = left
	in.operator = p.currToken
	p.proceed()
	in.right = p.parseExpression()
	return in
}

func (p *parser) proceed() {
	p.currToken = p.nextToken
	p.nextToken = p.l.nextToken()
}

// tok := p.currToken

// var exp expression
// switch tok._type {
// case t_SUB:
// 	var pre prefix
// 	pre.operator = tok
// 	pre.operand = p.parse()
// 	exp = pre
// case t_NUM:
// 	n, err := strconv.ParseInt(tok.literal, 10, 64)
// 	if err != nil {
// 		p.errors = append(p.errors, fmt.Errorf(`unable to parse integer "%v", %v`, tok.literal, err.Error()))
// 	}
// 	if precedences[p.currToken._type] > precedences[tok._type] {
// 		exp = p.parseInfix(integer(n))
// 	} else {
// 		exp = integer(n)
// 	}
// default:
// 	p.errors = append(p.errors, fmt.Errorf(`unexpected token "%v"`, tok._type.String()))
// }

// return exp
