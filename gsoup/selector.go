package gsoup

import (
	"strings"

	"fmt"

	"golang.org/x/net/html"
)

type Selection struct {
	root       *Node
	combinator uint8
	matcher    *AndMatcher // 当前计算器
	results    []*Node     // 当前结果节点
	workers    []*Node     // 当前工作节点
}

func (s *Selection) doSelect() {
	if len(s.results) > 0 {
		s.results = make([]*Node, 0)
	}
	var rescur func(root *Node, depth int)
	rescur = func(node *Node, depth int) {
		if s.matcher.matches(s.root, node) {
			if s.combinator == '>' && len(s.workers) > 0 {
				for _, w := range s.workers {
					hd := html.Node(*w)
					if node.Parent == &hd {
						s.results = append(s.results, node)
					}
				}
			} else {
				s.results = append(s.results, node)
			}
		}
		if depth < 0 {
			for n := node.FirstChild; n != nil; n = n.NextSibling {
				nd := Node(*n)
				rescur(&nd, depth)
			}
		} else if depth == 0 {
			return
		} else {
			depth = depth - 1
			for n := node.FirstChild; n != nil; n = n.NextSibling {
				nd := Node(*n)
				rescur(&nd, depth)
			}
		}
	}
	if len(s.workers) == 0 {
		s.workers = append(s.workers, s.root)
	}
	for _, w := range s.workers {
		rescur(w, -1)
	}
	if s.combinator == ' ' {
		s.workers = s.results
	} else if s.combinator == ',' {
		//s.workers = append()
	} else if s.combinator == '>' {

	} else if s.combinator == '+' {

	}
	s.matcher.matchers = make([]Matcher, 0)
}

func compile(root *Node, cssQuery string) (slct *Selection) {
	slct = &Selection{}
	slct.root = root
	cssQuery = strings.TrimSpace(cssQuery)
	tk := &TokenQueue{cssQuery: cssQuery, pos: 0}
	fmt.Printf("cssQuery: %s, posotion: %d\n", cssQuery, tk.pos)
	matcher := &AndMatcher{}
	for !tk.isEmpty() {
		if tk.matches("#") {
			id := tk.consumeId()
			mch := &IdMatcher{id}
			matcher.matchers = append(matcher.matchers, mch)
		} else if tk.matches(".") {
			className := tk.consumeClass()
			mch := &ClassMatcher{class: className}
			matcher.matchers = append(matcher.matchers, mch)
		} else if tk.matches("[") {
			attribute := tk.consumeAttribute()
			fmt.Println(attribute)
			mch := &AttributeMatcher{attribute: attribute}
			matcher.matchers = append(matcher.matchers, mch)
		} else if tk.matchCombinator() {
			if len(matcher.matchers) > 0 {
				slct.matcher = matcher
				slct.doSelect()
				slct.combinator = tk.consumeCombinator()
			}
		} else {
			tagName := tk.consumeTag()
			mch := &TagMatcher{tagName: tagName}
			matcher.matchers = append(matcher.matchers, mch)
		}
	}
	slct.matcher = matcher
	slct.doSelect()
	//slct.combinator = tk.consumeCombinator()
	return
}
