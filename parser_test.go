package expr

import (
	"testing"
)

func TestCanParseInteger(t *testing.T) {
	input := `123`
	l := newLexer([]byte(input))
	p := parser{l: l}

	n := p.parse()

	if len(p.errors) > 0 {
		t.Fatalf("parser has %d errors: %v", len(p.errors), p.errors)
	}

	if v, ok := n.(integer); !ok {
		t.Fatalf(`expected expression to be of type "integer", its of type "%T"`, n)
	} else if v != 123 {
		t.Fatalf(`expected integer to be 123, its "%v"`, v)
	}
}

func TestCanParseInfix(t *testing.T) {
	input := `1 + 2 + 3`
	l := newLexer([]byte(input))
	p := parser{l: l}

	n := p.parse()

	if len(p.errors) > 0 {
		t.Fatalf("parser has %d errors#0: %v", len(p.errors), p.errors[0])
	}

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

func TestCanParsePrefix(t *testing.T) {
	input := `-1`
	l := newLexer([]byte(input))
	p := parser{l: l}

	n := p.parse()

	if len(p.errors) > 0 {
		t.Fatalf("parser has %d errors: %v", len(p.errors), p.errors)
	}

	if inf, ok := n.(prefix); !ok {
		t.Fatalf(`expected expression to be of type "prefix", its of type "%T"`, n)
	} else if inf.operator._type != t_SUB {
		t.Fatalf(`expected prefix.operator to be -, its "%v"`, inf.operator._type)
	} else if v, ok := inf.operand.(integer); !ok {
		t.Fatalf(`expected prefix.operand to be of type "integer", its of type "%T"`, inf.operand)
	} else if v != 1 {
		t.Fatalf(`expected prefix.operand to be 1, its "%v"`, inf.operand)
	}
}
