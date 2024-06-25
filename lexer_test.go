package expr

import "testing"

func TestLexer(t *testing.T) {
	input := `123 + 456 * 99 - 199s`

	l := newLexer([]byte(input))

	tokens := []token{
		{_T_NUM, "123"},
		{_T_ADD, "+"},
		{_T_NUM, "456"},
		{_T_MUL, "*"},
		{_T_NUM, "99"},
		{_T_SUB, "-"},
		{_T_NUM, "199"},
		{_T_INVALID, "s"},
		{_T_EOF, "EOF"},
	}

	for i, tn := range tokens {
		if result := l.nextToken(); tn._type != result._type {
			t.Fatalf(`wrong token type "%s", expected "%s", case#%d`, result._type, tn._type, i)
		} else if tn.literal != result.literal {
			t.Fatalf(`wrong token litreal "%s", expected "%s", case#%d`, result.literal, tn.literal, i)
		}
	}
}
