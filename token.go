package expr

import "fmt"

type tokenType byte

type token struct {
	_type   tokenType
	literal string
}

const (
	_ tokenType = iota
	t_NUM
	t_INVALID
	t_EOF
)

var literals = map[tokenType]string{
	t_NUM:     "NUMBER",
	t_INVALID: "INVALID",
	t_EOF:     "EOF",
}

func (t tokenType) String() string {
	if l, ok := literals[t]; ok {
		return l
	}
	panic(fmt.Sprintf("token is unknown [%d]", t))
}
