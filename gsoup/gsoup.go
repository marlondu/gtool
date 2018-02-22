package gsoup

import (
	"bytes"
	"io"
	"log"
	"strings"

	"golang.org/x/net/html"
)

// Node is html.Node wrapper
type Node html.Node

// Parse parse reader to html Node
func Parse(r io.Reader) (*Node, error) {
	n, err := html.Parse(r)
	nd := Node(*n)
	return &nd, err
}

// Parse parse string to html Node
func ParseString(str string) (*Node, error) {
	n, err := html.Parse(strings.NewReader(str))
	nd := Node(*n)
	return &nd, err
}

// GetElementById get element node by id
func (n *Node) GetElementById(id string) *Node {
	results := n.Select("#" + id)
	if results == nil || len(results) == 0 {
		return nil
	} else {
		return results[0]
	}
}

// Attribute return attr's value by attr name
func (n *Node) Attribute(key string) (val string) {
	for _, attr := range n.Attr {
		if attr.Key == key {
			val = attr.Val
			break
		}
	}
	return
}

// Html return element's inner html
func (n *Node) Html() string {
	buf := bytes.NewBuffer([]byte{})
	hd := html.Node(*n)
	err := html.Render(buf, &hd)
	if err != nil {
		log.Panic(err)
	}
	return buf.String()
}

// Text return element's inner text
func (n *Node) Text() string {
	if n.Type == html.TextNode {
		return n.Data
	} else {
		if n.FirstChild != nil {
			buf := bytes.NewBuffer([]byte{})
			for n := n.FirstChild; n != nil; n = n.NextSibling {
				nd := Node(*n)
				buf.WriteString(nd.Text())
			}
			buf.String()
		}
	}
	return ""
}

// Select select elements by css query string
func (d *Node) Select(cssQuery string) []*Node {
	slct := compile(d, cssQuery)
	nds := slct.results
	results := make([]*Node, len(nds))
	for i := 0; i < len(nds); i++ {
		res := Node(*nds[i])
		//results = append(results, &res)
		results[i] = &res
	}
	return results
}
