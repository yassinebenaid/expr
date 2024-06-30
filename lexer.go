package expr

import "fmt"

type tokenType byte

type token struct {
	_type   tokenType
	literal string
}

const (
	_ tokenType = iota
	_T_INT
	_T_FLOAT
	_T_ADD
	_T_SUB
	_T_MUL
	_T_DEV
	_T_LPAR
	_T_RPAR
	_T_BINAND
	_T_BINOR
	_T_BINLSHIFT
	_T_BINRSHIFT
	_T_INVALID
	_T_EOF
)

var literals = map[tokenType]string{
	_T_INT:       "INT",
	_T_FLOAT:     "FLOAT",
	_T_ADD:       "+",
	_T_SUB:       "-",
	_T_MUL:       "*",
	_T_DEV:       "/",
	_T_LPAR:      "(",
	_T_RPAR:      ")",
	_T_BINOR:     "|",
	_T_BINAND:    "&",
	_T_BINLSHIFT: "<<",
	_T_BINRSHIFT: ">>",
	_T_INVALID:   "INVALID",
	_T_EOF:       "EOF",
}

func (t tokenType) String() string {
	if l, ok := literals[t]; ok {
		return l
	}
	panic(fmt.Sprintf("token is unknown [%d]", t))
}

const (
	_ = iota
	_PREC_LOW
	_PREC_MEDIUM
	_PREC_HIGH
)

var precedences = map[tokenType]int{
	_T_ADD:       _PREC_MEDIUM,
	_T_SUB:       _PREC_MEDIUM,
	_T_BINOR:     _PREC_MEDIUM,
	_T_MUL:       _PREC_HIGH,
	_T_DEV:       _PREC_HIGH,
	_T_BINAND:    _PREC_HIGH,
	_T_BINLSHIFT: _PREC_HIGH,
	_T_BINRSHIFT: _PREC_HIGH,
}

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
	var skip_reading bool

	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' {
		l.readCh()
	}

	switch {
	case l.ch == 0:
		tok._type, tok.literal = _T_EOF, "EOF"
	case l.ch == '+':
		tok._type, tok.literal = _T_ADD, string(l.ch)
	case l.ch == '-':
		tok._type, tok.literal = _T_SUB, string(l.ch)
	case l.ch == '*':
		tok._type, tok.literal = _T_MUL, string(l.ch)
	case l.ch == '/':
		tok._type, tok.literal = _T_DEV, string(l.ch)
	case l.ch == '(':
		tok._type, tok.literal = _T_LPAR, string(l.ch)
	case l.ch == ')':
		tok._type, tok.literal = _T_RPAR, string(l.ch)
	case l.ch == '|':
		tok._type, tok.literal = _T_BINOR, string(l.ch)
	case l.ch == '&':
		tok._type, tok.literal = _T_BINAND, string(l.ch)
	case l.ch == '>' && l.peek == '>':
		l.readCh()
		tok._type, tok.literal = _T_BINRSHIFT, ">>"
	case l.ch == '<' && l.peek == '<':
		l.readCh()
		tok._type, tok.literal = _T_BINLSHIFT, "<<"
	case (l.ch <= '9' && l.ch >= '0') || (l.ch == '.' && (l.peek <= '9' && l.peek >= '0')):
		var num []byte
		tok._type = _T_INT

		for {
			if l.ch == '.' {
				if tok._type == _T_FLOAT {
					skip_reading = true
					break
				}
				tok._type = _T_FLOAT
			}
			num = append(num, l.ch)

			if (l.peek > '9' || l.peek < '0') && l.peek != '.' {
				break
			}
			l.readCh()
		}

		tok.literal = string(num)
	default:
		tok._type, tok.literal = _T_INVALID, string(l.ch)
	}

	if !skip_reading {
		l.readCh()
	}

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
