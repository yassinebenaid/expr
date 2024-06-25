package expr

import (
	"testing"
)

func TestCanParseInteger(t *testing.T) {
	input := `123`
	l := newLexer([]byte(input))
	p := newParser(l)
	exp := p.parse()

	if len(p.errors) > 0 {
		t.Fatalf("parser has %d errors#0: %v", len(p.errors), p.errors[0])
	}

	if v, ok := exp.(integer); !ok {
		t.Fatalf(`expected expression to be of type "integer", its of type "%T"`, exp)
	} else if v != 123 {
		t.Fatalf(`expected integer to be 123, its "%v"`, v)
	}
}

func TestCanParseInfix(t *testing.T) {
	input := `1 + 2`
	l := newLexer([]byte(input))
	p := newParser(l)
	exp := p.parse()

	if len(p.errors) > 0 {
		t.Fatalf("parser has %d errors#0: %v", len(p.errors), p.errors[0])
	}

	if inf, ok := exp.(infix); !ok {
		t.Fatalf(`expected expression to be of type "infix", its of type "%T"`, exp)
	} else if inf.operator._type != _T_ADD {
		t.Fatalf(`expected infix.operator to be +, its "%v"`, inf.operator._type)
	} else if v, ok := inf.left.(integer); !ok {
		t.Fatalf(`expected infix.left to be of type "integer", its of type "%T"`, inf.left)
	} else if v != 1 {
		t.Fatalf(`expected infix.left to be 1, its "%v"`, inf.left)
	} else if v, ok := inf.right.(integer); !ok {
		t.Fatalf(`expected infix.right to be of type "integer", its of type "%T"`, inf.right)
	} else if v != 2 {
		t.Fatalf(`expected infix.right to be 2, its "%v"`, inf.right)
	}
}

func TestCanParsePrefix(t *testing.T) {
	input := `-1`
	l := newLexer([]byte(input))
	p := newParser(l)
	exp := p.parse()

	if len(p.errors) > 0 {
		t.Fatalf("parser has %d errors#0: %v", len(p.errors), p.errors[0])
	}

	if inf, ok := exp.(prefix); !ok {
		t.Fatalf(`expected expression to be of type "prefix", its of type "%T"`, exp)
	} else if inf.operator._type != _T_SUB {
		t.Fatalf(`expected prefix.operator to be -, its "%v"`, inf.operator._type)
	} else if v, ok := inf.operand.(integer); !ok {
		t.Fatalf(`expected prefix.operand to be of type "integer", its of type "%T"`, inf.operand)
	} else if v != 1 {
		t.Fatalf(`expected prefix.operand to be 1, its "%v"`, inf.operand)
	}
}

func TestCanParse(t *testing.T) {
	testCases := []struct {
		input              string
		expectedExpression string
	}{
		{"1", "1"},
		{"-1", "(-1)"},
		{"--1", "(-(-1))"},
		{"1 + 2", "(1 + 2)"},
		{"1 - 2", "(1 - 2)"},
		{"1 + 2 - 3 + 4 - 5", "((((1 + 2) - 3) + 4) - 5)"},
		{"-1 + -2 - -3 + -4 - -5", "(((((-1) + (-2)) - (-3)) + (-4)) - (-5))"},
		{"1 * 2", "(1 * 2)"},
		{"-1 * -2", "((-1) * (-2))"},
		{"1 + 2 * 3 - -2", "((1 + (2 * 3)) - (-2))"},
	}

	for i, tc := range testCases {
		l := newLexer([]byte(tc.input))
		p := newParser(l)
		n := p.parse()

		if len(p.errors) > 0 {
			t.Fatalf("parser has %d errors#0: %v", len(p.errors), p.errors[0])
		}

		if exp := n.String(); tc.expectedExpression != exp {
			t.Fatalf(`case#%d: expected "%s", got "%s"`, i, tc.expectedExpression, exp)
		}
	}
}

func TestParserErrors(t *testing.T) {
	testCases := []struct {
		input         string
		expectedError string
	}{
		{"x", `unexpected token "x"`},
		{"1 + x", `unexpected token "x"`},
		{"1 + ", `unexpected token "EOF"`},
	}

	for i, tc := range testCases {
		l := newLexer([]byte(tc.input))
		p := newParser(l)
		p.parse()

		if len(p.errors) == 0 {
			t.Fatalf("parser has no errors")
		}

		if err := p.errors[0]; tc.expectedError != err.Error() {
			t.Fatalf(`case#%d: expected error "%s", got "%s"`, i, tc.expectedError, err)
		}
	}
}
