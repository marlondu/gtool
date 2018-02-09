package gsoup

import (
	"fmt"
	"strings"
	"unicode"
)

var combinators = [5]uint8{' ', ',', '>', '+', '~'}

type TokenQueue struct {
	cssQuery string
	pos      int
}

func (t *TokenQueue) matches(seq string) bool {
	return strings.HasPrefix(t.cssQuery[t.pos:], seq)
}

func (t *TokenQueue) consumeId() string {
	start := t.pos + 1
	for t.pos = start; t.pos < len(t.cssQuery); t.pos = t.pos + 1 {
		if !isLetterOrDigit(t.cssQuery[t.pos]) {
			break
		}
	}
	return t.cssQuery[start:t.pos]
}

func (t *TokenQueue) consumeClass() string {
	return t.consumeId()
}

func (t *TokenQueue) consumeAttribute() string {
	start := t.pos + 1
	for ; t.pos < len(t.cssQuery); t.pos++ {
		if t.cssQuery[t.pos] == ']' {
			break
		}
	}
	attr := t.cssQuery[start:t.pos]
	t.pos = t.pos + 1
	return attr
}

func (t *TokenQueue) consumeTag() string {
	start := t.pos
	for t.pos = start; t.pos < len(t.cssQuery); t.pos = t.pos + 1 {
		if !isLetterOrDigit(t.cssQuery[t.pos]) {
			break
		}
	}
	return t.cssQuery[start:t.pos]
}

func (t *TokenQueue) matchCombinator() bool {
	for _, com := range combinators {
		if com == t.cssQuery[t.pos] {
			return true
		}
	}
	return false
}

func (t *TokenQueue) consumeCombinator() uint8 {
	start := t.pos
	t.pos = t.pos + 1
	fmt.Printf("start: %d, pos: %d\n", start, t.pos)
	return t.cssQuery[start]
}

func (t *TokenQueue) isEmpty() bool {
	l := len(t.cssQuery)
	return t.pos >= l
}

func isLetterOrDigit(r uint8) bool {
	return unicode.IsDigit(rune(r)) || unicode.IsLetter(rune(r)) || r == '-' || r == '_'
}
