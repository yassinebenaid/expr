package expr

import "fmt"

type expression interface {
	String() string
}

type infix struct {
	left     expression
	operator token
	right    expression
}

func (i infix) String() string {
	return fmt.Sprintf("(%v %v %v)", i.left.String(), i.operator.literal, i.right.String())
}

type integer int

func (i integer) String() string {
	return fmt.Sprintf("%d", i)
}
