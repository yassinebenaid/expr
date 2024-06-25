package expr

type lexer struct {
	input []byte
	pos   int
	ch    byte
	peek  byte
}

func newLexer(in []byte) *lexer {
	l := &lexer{input: in}
	l.readCh()
	l.readCh()
	return l
}

func (l *lexer) nextToken() token {
	var tok token

	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' {
		l.readCh()
	}

	switch {
	case l.ch == 0:
		tok._type, tok.literal = _T_EOF, "EOF"
	case l.ch == '+':
		tok._type, tok.literal = _T_ADD, "+"
	case l.ch == '-':
		tok._type, tok.literal = _T_SUB, "-"
	case l.ch == '*':
		tok._type, tok.literal = _T_MUL, "*"
	case l.ch <= '9' && l.ch >= '0':
		var num []byte
		for {
			num = append(num, l.ch)

			if l.peek > '9' || l.peek < '0' {
				break
			}
			l.readCh()
		}
		tok._type, tok.literal = _T_NUM, string(num)
	default:
		tok._type, tok.literal = _T_INVALID, string(l.ch)
	}

	l.readCh()
	return tok
}

func (l *lexer) readCh() {
	l.ch = l.peek
	if l.pos >= len(l.input) {
		l.peek = 0
	} else {
		l.peek = l.input[l.pos]
	}
	l.pos++
}
