package tokenizer

type Lexer struct {
	Input string

	Character byte
	Index     int
	PeekIndex int
}

func NewLexer(input string) *Lexer {
	lexer := &Lexer{
		Input: input,
	}

	lexer.readNext()

	return lexer
}

func (l *Lexer) readNext() {
	if l.PeekIndex >= len(l.Input) {
		l.Character = 0
	} else {
		l.Character = l.Input[l.PeekIndex]
	}

	l.Index = l.PeekIndex
	l.PeekIndex++
}

func (l *Lexer) peek() byte {
	if l.PeekIndex >= len(l.Input) {
		return 0
	} else {
		return l.Input[l.PeekIndex]
	}
}

func isLetter(c byte) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || c == '\''
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func (l *Lexer) readNumber() string {
	idx := l.Index

	for isDigit(l.Character) {
		l.readNext()
	}

	return l.Input[idx:l.Index]
}

func (l *Lexer) readWord() string {
	idx := l.Index

	for isLetter(l.Character) {
		l.readNext()
	}

	return l.Input[idx:l.Index]
}

func (l *Lexer) readCommand() string {
	idx := l.Index

	for l.Character != ']' {
		l.readNext()
	}

	return l.Input[idx:l.Index]
}

func (l *Lexer) skipWhiteSpace() {
	for l.Character == ' ' || l.Character == '\t' || l.Character == '\n' || l.Character == '\r' {
		l.readNext()
	}
}

func newToken(tokenType TokenType, characters ...byte) Token {
	value := ""

	for _, ch := range characters {
		value += string(ch)
	}

	return Token{Type: tokenType, Value: value}
}

func (l *Lexer) NextToken() *Token {
	var token Token

	l.skipWhiteSpace()

	switch l.Character {
	case '[':
		l.readNext()
		token.Type = COMMAND
		token.Value = l.readCommand()
	case ']':
		l.readNext()
	case 0:
		token.Type = EOF
		token.Value = ""
	default:
		if isLetter(l.Character) {
			// TODO: Use db and wiki to determine if word is important
			token.Type = IGNORED_WORD
			token.Value = l.readWord()
		} else if isDigit(l.Character) {
			token.Type = NUMBER
			token.Value = l.readNumber()
		}
	}

	return &token
}
