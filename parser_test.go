package expr

import "testing"

func TestCanParseInfix(t *testing.T) {
	input := `1 + 2 + 3`
	l := newLexer([]byte(input))
	p := parser{l}

	n := p.parse()

	if inf, ok := n.(infix); !ok {
		t.Fatalf(`expected expression to be of type "infix", its of type "%T"`, n)
	} else if inf.operator._type != t_ADD {
		t.Fatalf(`expected infix.operator to be +, its "%v"`, inf.operator._type)
	} else if v, ok := inf.left.(integer); !ok {
		t.Fatalf(`expected infix.left to be of type "integer", its of type "%T"`, inf.left)
	} else if v != 1 {
		t.Fatalf(`expected infix.left to be 1, its "%v"`, inf.left)
	} else if inf, ok := inf.right.(infix); !ok {
		t.Fatalf(`expected inf.right to be of type "infix", its of type "%T"`, n)
	} else if v, ok := inf.left.(integer); !ok {
		t.Fatalf(`expected infix.left to be of type "integer", its of type "%T"`, inf.left)
	} else if v != 2 {
		t.Fatalf(`expected infix.left to be 2, its "%v"`, inf.left)
	} else if v, ok := inf.right.(integer); !ok {
		t.Fatalf(`expected infix.right to be of type "integer", its of type "%T"`, inf.right)
	} else if v != 3 {
		t.Fatalf(`expected infix.right to be 3, its "%v"`, inf.right)
	} else if inf.operator._type != t_ADD {
		t.Fatalf(`expected infix.operator to be +, its "%v"`, inf.operator._type)
	}
}
