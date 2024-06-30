package expr

import "fmt"

type Number float64

func (n Number) ToString() string {
	return fmt.Sprint(n)
}

func eval(exp expression) Number {
	switch v := exp.(type) {
	case numberLiteral:
		return Number(v)
	case prefix:
		return evalPrefix(v)
	default:
		panic("eval: expression not supported")
	}
}

func evalPrefix(exp prefix) Number {
	value := eval(exp.operand)

	switch exp.operator._type {
	case _T_SUB:
		return -value
	case _T_ADD:
		return value
	default:
		panic("evalPrefix: prefix is undefined")
	}
}
