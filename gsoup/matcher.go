package gsoup

import (
	"log"
	"strings"

	"golang.org/x/net/html"
)

type Matcher interface {
	matches(root *Node, node *Node) bool
}

type IdMatcher struct {
	id string
}

func (m IdMatcher) matches(root *Node, node *Node) bool {
	for _, attr := range node.Attr {
		if strings.ToLower(attr.Key) == "id" && attr.Val == m.id {
			return true
		}
	}
	return false
}

type ClassMatcher struct {
	class string
}

func (m ClassMatcher) matches(root *Node, node *Node) bool {
	for _, attr := range node.Attr {
		if strings.ToLower(attr.Key) == "class" && attr.Val == m.class {
			return true
		}
	}
	return false
}

type AttributeMatcher struct {
	attribute string // [attribute=val]不带[]
}

func (m AttributeMatcher) matches(root *Node, node *Node) bool {
	var (
		attrName  string
		opt       string
		attrValue string
	)
	if m.attribute[0] == '=' {
		log.Panic("attribute matcher invalid")
	}
	i, pos0, pos1 := 0, 0, 0
	for _, c := range m.attribute {
		if pos0 == 0 && (c == '=' || c == '~' || c == '|' || c == '*' || c == '$' || c == '^') {
			pos0 = i
			pos1 = i
		}
		if pos0 > 0 && i > pos0 && c != '=' && c != '~' && c != '|' && c != '*' && c != '$' && c != '^' {
			pos1 = i
			break
		}
		i++
	}
	if pos0 == 0 || pos0 == pos1 {
		attrName = m.attribute
		for _, attr := range node.Attr {
			if attr.Key == attrName {
				return true
			}
		}
		return false
	} else {
		attrName = m.attribute[:pos0]
		opt = m.attribute[pos0:pos1]
		attrValue = m.attribute[pos1:]
		if opt == "=" {
			for _, attr := range node.Attr {
				if attr.Key == attrName && attr.Val == attrValue {
					return true
				}
			}
			return false
		} else if opt == "~=" || opt == "*=" {
			for _, attr := range node.Attr {
				if attr.Key == attrName && strings.Contains(attr.Val, attrValue) {
					return true
				}
			}
			return false
		} else if opt == "|=" || opt == "^=" {
			for _, attr := range node.Attr {
				if attr.Key == attrName && strings.HasPrefix(attr.Val, attrValue) {
					return true
				}
			}
		} else if opt == "$=" {
			for _, attr := range node.Attr {
				if attr.Key == attrName && strings.HasSuffix(attr.Val, attrValue) {
					return true
				}
			}
		}
	}
	return false
}

type TagMatcher struct {
	tagName string
}

func (m TagMatcher) matches(root *Node, node *Node) bool {
	return node.Data == m.tagName
}

type AndMatcher struct {
	matchers []Matcher
}

func (m AndMatcher) matches(root *Node, node *Node) bool {
	for _, mch := range m.matchers {
		if !mch.matches(root, node) {
			return false
		}
	}
	return node.Type == html.ElementNode
}
