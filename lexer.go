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

	switch l.ch {
	case 0:
		tok._type = t_EOF
	case '+':
		tok._type, tok.literal = t_ADD, "+"
	case '-':
		tok._type, tok.literal = t_SUB, "-"
	default:
		if l.ch <= '9' && l.ch >= '0' {
			var num []byte
			for ; l.ch <= '9' && l.ch >= '0'; l.readCh() {
				num = append(num, l.ch)
			}
			tok._type, tok.literal = t_NUM, string(num)
		} else {
			tok._type = t_INVALID
		}
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

func (l *lexer) readToken() token {
	var peek, ch, pos = l.peek, l.ch, l.pos
	tok := l.nextToken()
	l.peek, l.ch, l.pos = peek, ch, pos
	return tok
}
