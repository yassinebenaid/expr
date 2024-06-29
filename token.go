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
