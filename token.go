package expr

import "fmt"

type tokenType byte

type token struct {
	_type   tokenType
	literal string
}

const (
	_ tokenType = iota
	_T_NUM
	_T_ADD
	_T_SUB
	_T_MUL
	_T_INVALID
	_T_EOF
)

var literals = map[tokenType]string{
	_T_NUM:     "NUMBER",
	_T_ADD:     "+",
	_T_SUB:     "-",
	_T_MUL:     "*",
	_T_INVALID: "INVALID",
	_T_EOF:     "EOF",
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
	_T_ADD: _PREC_MEDIUM,
	_T_SUB: _PREC_MEDIUM,
	_T_MUL: _PREC_HIGH,
}
