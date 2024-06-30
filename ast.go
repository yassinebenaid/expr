package expr

import "fmt"

type expression interface {
	ToString() string
}

type infix struct {
	left     expression
	operator token
	right    expression
}

func (i infix) ToString() string {
	return fmt.Sprintf("(%v %v %v)", i.left.ToString(), i.operator.literal, i.right.ToString())
}

type prefix struct {
	operator token
	operand  expression
}

func (i prefix) ToString() string {
	return fmt.Sprintf("(%v%v)", i.operator.literal, i.operand.ToString())
}

type numberLiteral int

func (i numberLiteral) ToString() string {
	return fmt.Sprint(i)
}

type float float64

func (i float) ToString() string {
	return fmt.Sprint(i)
}
