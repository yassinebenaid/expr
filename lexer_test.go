package expr

import "testing"

func TestLexer(t *testing.T) {
	input := `123 ,+`

	l := newLexer([]byte(input))

	tokens := []token{
		{t_NUM, "123"},
		{t_INVALID, ""},
		{t_INVALID, ""},
		{t_EOF, ""},
	}

	for i, tn := range tokens {
		if result := l.nextToken(); tn._type != result._type {
			t.Fatalf(`wrong token type "%s", expected "%s", case#%d`, result._type, tn._type, i)
		} else if tn.literal != result.literal {
			t.Fatalf(`wrong token litreal "%s", expected "%s", case#%d`, result.literal, tn.literal, i)
		}
	}
}
