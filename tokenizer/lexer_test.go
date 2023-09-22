package tokenizer

import (
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "This is a message for 5 of my friends in which I will [roll 2d20]."

	tests := []struct {
		expectedType  TokenType
		expectedValue string
	}{
		{IGNORED_WORD, "This"},
		{IGNORED_WORD, "is"},
		{IGNORED_WORD, "a"},
		{IGNORED_WORD, "message"},
		{IGNORED_WORD, "for"},
		{NUMBER, "5"},
		{IGNORED_WORD, "of"},
		{IGNORED_WORD, "my"},
		{IGNORED_WORD, "friends"},
		{IGNORED_WORD, "in"},
		{IGNORED_WORD, "which"},
		{IGNORED_WORD, "I"},
		{IGNORED_WORD, "will"},
		{COMMAND, "roll 2d20"},
	}

	l := NewLexer(input)

	for i, tt := range tests {
		token := l.NextToken()

		if token.Type != tt.expectedType {
			t.Fatalf(
				"tests[%d] - expected type %q, got %q",
				i,
				tt.expectedType,
				token.Type,
			)
		}

		if token.Value != tt.expectedValue {
			t.Fatalf(
				"tests[%d] - expected value %q, got %q",
				i,
				tt.expectedValue,
				token.Value,
			)
		}
	}
}
