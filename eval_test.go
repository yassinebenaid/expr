package expr

import (
	"testing"
)

func TestCanEval(t *testing.T) {
	testCases := []struct {
		input              string
		expectedExpression Value
	}{
		{"1", Number(1)},
		{"-1", Number(-1)},
		{"--1", Number(1)},
		{"+1", Number(1)},
		// {"1 + 2", Number(3)},
		// {"1 - 2", "(1 - 2)"},
		// {"1 + 2 - 3 + 4 - 5", "((((1 + 2) - 3) + 4) - 5)"},
		// {"-1 + -2 - -3 + -4 - -5", "(((((-1) + (-2)) - (-3)) + (-4)) - (-5))"},
		// {"1 * 2", "(1 * 2)"},
		// {"-1 * -2", "((-1) * (-2))"},
		// {"1 + 2 * 3 - -2", "((1 + (2 * 3)) - (-2))"},
		// {"1 / 2", "(1 / 2)"},
		// {"-1 / -2", "((-1) / (-2))"},
		// {"1 + 2 / 3 - -2", "((1 + (2 / 3)) - (-2))"},
		// {"(1 + 2) * (3 * 4)", "((1 + 2) * (3 * 4))"},
		// {"1 | 2", "(1 | 2)"},
		// {"1 & 2", "(1 & 2)"},
		// {"1 << 2", "(1 << 2)"},
		// {"1 >> 2", "(1 >> 2)"},
		// {"1 | 2 * 3", "(1 | (2 * 3))"},
		// {"1 & 2 * 3", "((1 & 2) * 3)"},
		// {"1 << 2 * 3", "((1 << 2) * 3)"},
		// {"1 >> 2 * 3", "((1 >> 2) * 3)"},
		// {"1 + 2 | 3 & 4 * 5", "((1 + 2) | ((3 & 4) * 5))"},
		// {"1 + 2.2 | .3 & 4. * 5 << 6 / 7 - 8 >> 9", "(((1 + 2.2) | ((((0.3 & 4) * 5) << 6) / 7)) - (8 >> 9))"},
	}

	for i, tc := range testCases {
		l := newLexer([]byte(tc.input))
		p := newParser(l)
		exp := p.parse()

		if len(p.errors) > 0 {
			t.Fatalf("parser has %d errors#0: %v", len(p.errors), p.errors[0])
		}

		v := eval(exp)

		if str := v.ToString(); tc.expectedExpression.ToString() != str {
			t.Fatalf(`case#%d: expected "%s", got "%s"`, i, tc.expectedExpression.ToString(), str)
		}
	}
}
