package expr

type lexer struct {
	input []byte
	pos   int
	ch    byte
}

func newLexer(in []byte) *lexer {
	l := &lexer{input: in}
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
		l.readCh()
	default:
		if l.ch <= '9' && l.ch >= '0' {
			var num []byte
			for ; l.ch <= '9' && l.ch >= '0'; l.readCh() {
				num = append(num, l.ch)
			}
			tok._type, tok.literal = t_NUM, string(num)
		} else {
			tok._type = t_INVALID
			l.readCh()
		}
	}

	return tok
}

func (l *lexer) readCh() {
	if l.pos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.pos]
		l.pos++
	}
}
