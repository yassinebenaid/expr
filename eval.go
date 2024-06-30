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
	case infix:
		return evalInfix(v)
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

func evalInfix(exp infix) Number {
	l, r := eval(exp.left), eval(exp.right)

	switch exp.operator._type {
	case _T_SUB:
		return l - r
	case _T_ADD:
		return l + r
	case _T_MUL:
		return l * r
	case _T_DEV:
		return l / r
	case _T_BINOR:
		return Number(int64(l) | int64(r))
	case _T_BINAND:
		return Number(int64(l) & int64(r))
	case _T_BINLSHIFT:
		return Number(int64(l) << int64(r))
	case _T_BINRSHIFT:
		return Number(int64(l) >> int64(r))
	default:
		panic("evalPrefix: operator is undefined")
	}
}
