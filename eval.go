package expr

import "fmt"

type Value interface {
	ToString() string
}

type Number float64

func (n Number) ToString() string {
	return fmt.Sprint(n)
}

type Boolean bool

func (n Boolean) ToString() string {
	return fmt.Sprintf("%t", n)
}

func eval(exp expression) Value {
	switch v := exp.(type) {
	case integer:
		return Number(v)
	case prefix:
		return evalPrefix(v)
	default:
		panic("eval: expression not supported")
	}
}

func evalPrefix(exp prefix) Value {
	value := eval(exp.operand)

	switch exp.operator._type {
	case _T_SUB:
		num, ok := value.(Number)
		if !ok {
			panic("evalPrefix: operand is not a number")
		}
		return Number(-num)
	case _T_ADD:
		num, ok := value.(Number)
		if !ok {
			panic("evalPrefix: operand is not a number")
		}
		return Number(num)
	default:
		panic("evalPrefix: prefix is undefined")
	}
}
