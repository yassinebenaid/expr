package expr

import "strconv"

type parser struct {
	l *lexer
}

func (p *parser) parse() expression {
	tok := p.l.nextToken()

	var exp expression
	switch tok._type {
	case t_NUM:
		n, _ := strconv.ParseInt(tok.literal, 2, 64)
		if precedences[p.l.readToken()._type] > precedences[tok._type] {
			exp = p.parseInfix(integer(n))
		} else {
			exp = integer(n)
		}
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
