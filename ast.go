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

type prefix struct {
	operator token
	operand  expression
}

func (i prefix) String() string {
	return fmt.Sprintf("(%v%v)", i.operator.literal, i.operand.String())
}

type integer int

func (i integer) String() string {
	return fmt.Sprintf("%d", i)
}
