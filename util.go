package gquery

import (
	"log"
	"strings"

	"golang.org/x/net/html"
)

// const (
//     ErrorNode NodeType = iota
//     TextNode			1 text
//     DocumentNode		2 html
//     ElementNode		3 node
//     CommentNode
//     DoctypeNode
// )

func nodePrev(node *html.Node) *html.Node {
	for prev := node.PrevSibling; prev != nil; prev = prev.PrevSibling {
		if prev.Type == html.ElementNode {
			return prev
		} else if prev.Type == html.TextNode {
			if len(strings.TrimSpace(prev.Data)) == 0 {
				continue
			}
			return prev
		}
	}
	return nil
}
func nodeNext(node *html.Node) *html.Node {
	//node.Type
	//1	text and space,if space then skip
	//3	tag
	for next := node.NextSibling; next != nil; next = next.NextSibling {
		if next.Type == html.ElementNode {
			return next
		} else if next.Type == html.TextNode {
			d := strings.TrimSpace(next.Data)
			if len(d) == 0 {
				continue
			}
			return next
		}
	}
	return nil
}
func nodeChild(node *html.Node) []*html.Node {
	nodes := make([]*html.Node, 0)
	for n := node.FirstChild; n != nil; n = n.NextSibling {
		if n.Type == html.ElementNode {
			nodes = append(nodes, n)

		} else if n.Type == html.TextNode {
			if len(strings.TrimSpace(n.Data)) == 0 {
				continue
			}
			nodes = append(nodes, n)
		}
	}

	return nodes
}

func nodeCopy(src *html.Node) *html.Node {
	dst := new(html.Node)

	dst.Type = src.Type
	dst.DataAtom = src.DataAtom
	dst.Data = src.Data
	dst.Namespace = src.Namespace
	dst.Attr = append(dst.Attr, src.Attr...)

	si := src.FirstChild
	if si != nil {
		for {
			dst.AppendChild(nodeCopy(si))
			si = si.NextSibling

			if si != nil {
				continue
			} else {
				break
			}
		}
	}

	return dst
}

func nodeSwitch(n interface{}) *html.Node {
	switch n.(type) {
	case string:
		str, _ := n.(string)
		return nodeNewByString(str)
	case *html.Node:
		node, _ := n.(*html.Node)
		return node
	case *Elements:
		elms, _ := n.(*Elements)
		if len(elms.nodes) > 0 {
			return elms.nodes[0]
		}
	}
	return nil
}

func nodeNewByString(src string) *html.Node {
	htmlnode, err := html.Parse(strings.NewReader(src))
	if err != nil {
		log.Panicln(err)
	}
	parent := htmlnode.LastChild.LastChild
	node := htmlnode.LastChild.LastChild.FirstChild

	parent.RemoveChild(node)
	return node
}

func nodesNewByString(src string) []*html.Node {
	htmlnode, err := html.Parse(strings.NewReader(src))
	if err != nil {
		log.Panicln(err)
	}

	parent := htmlnode.LastChild.LastChild
	child := nodeChild(parent)
	for _, node := range child {
		parent.RemoveChild(node)
	}

	return child
}

func newElements(root *html.Node, nodes []*html.Node) *Elements {
	elements := new(Elements)
	if root == nil || len(nodes) == 0 {
		return nil
	}
	elements.root = root
	elements.nodes = nodes

	return elements
}

func newElement(root *html.Node, node *html.Node) *Element {
	element := new(Element)
	if root == nil {
		return nil
	}
	element.root = root
	element.nodes = []*html.Node{node}

	return element
}

func toElements(element *Element) *Elements {
	elements := new(Elements)
	elements.root = element.root
	elements.nodes = []*html.Node{element.nodes[0]}

	return elements
}

func toElement(elements *Elements) *Element {
	element := new(Element)
	element.root = elements.root
	element.nodes = []*html.Node{elements.nodes[0]}

	return element
}

func elementsAppendNode(elements *Elements, node *html.Node) {
	elements.nodes = append(elements.nodes, node)
}
func elementsAppendNodes(elements *Elements, nodes []*html.Node) {
	elements.nodes = append(elements.nodes, nodes...)
}
