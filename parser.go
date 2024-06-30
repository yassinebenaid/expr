package expr

import (
	"fmt"
	"strconv"
)

type parser struct {
	l         *lexer
	currToken token
	errors    []error

	prefixParser map[tokenType]func() expression
	infixParser  map[tokenType]func(expression) expression
}

func newParser(l *lexer) *parser {
	p := &parser{
		l: l,
	}
	p.proceed()

	p.prefixParser = map[tokenType]func() expression{
		_T_INT:   p.parseNumberLiteral,
		_T_FLOAT: p.parseFloat,
		_T_SUB:   p.parsePrefix,
		_T_ADD:   p.parsePrefix,
		_T_LPAR:  p.parseGroupedExp,
	}
	p.infixParser = map[tokenType]func(expression) expression{
		_T_ADD:       p.parseInfix,
		_T_SUB:       p.parseInfix,
		_T_MUL:       p.parseInfix,
		_T_DEV:       p.parseInfix,
		_T_BINAND:    p.parseInfix,
		_T_BINOR:     p.parseInfix,
		_T_BINLSHIFT: p.parseInfix,
		_T_BINRSHIFT: p.parseInfix,
	}
	return p
}

func (p *parser) parse() expression {
	var exp expression = p.parseExpression(_PREC_LOW)

	if p.currToken._type != _T_EOF {
		p.errors = append(p.errors, fmt.Errorf(`unexpected token "%v", expected end of expression`, p.currToken.literal))
	}

	return exp
}

func (p *parser) parseExpression(precedence int) expression {
	pp := p.prefixParser[p.currToken._type]
	if pp == nil {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %q", p.currToken.literal))
		return nil
	}

	exp := pp()

	for precedence < precedences[p.currToken._type] {
		ip := p.infixParser[p.currToken._type]
		if ip == nil {
			return exp
		}
		exp = ip(exp)
	}

	return exp
}

func (p *parser) parseNumberLiteral() expression {
	n, err := strconv.ParseInt(p.currToken.literal, 10, 64)
	if err != nil {
		p.errors = append(p.errors, fmt.Errorf(`unable to parse number "%v"`, p.currToken.literal))
	}
	p.proceed()
	return numberLiteral(n)
}

func (p *parser) parseFloat() expression {
	n, err := strconv.ParseFloat(p.currToken.literal, 64)
	if err != nil {
		p.errors = append(p.errors, fmt.Errorf(`unable to parse float "%v"`, p.currToken.literal))
	}
	p.proceed()
	return numberLiteral(n)
}

func (p *parser) parsePrefix() expression {
	var pref prefix
	pref.operator = p.currToken
	p.proceed()
	pref.operand = p.parseExpression(_PREC_HIGH)
	return pref
}

func (p *parser) parseGroupedExp() expression {
	p.proceed()
	exp := p.parseExpression(_PREC_LOW)

	if p.currToken._type != _T_RPAR {
		p.errors = append(p.errors, fmt.Errorf("unclosed grouped expression, expected %q", _T_RPAR))
	}

	p.proceed()
	return exp
}

func (p *parser) parseInfix(left expression) expression {
	var in infix
	in.left = left
	in.operator = p.currToken
	precedence := precedences[p.currToken._type]
	p.proceed()

	if p.currToken._type == _T_EOF {
		p.errors = append(p.errors, fmt.Errorf("unexpected end of expression"))
		return nil
	}

	in.right = p.parseExpression(precedence)
	return in
}

func (p *parser) proceed() {
	p.currToken = p.l.nextToken()
}
