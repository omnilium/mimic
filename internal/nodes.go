package internal

import "fmt"

type Node interface {
	String() string
	Render() (string, error)
	RenderAnnotated() (string, error)
	GetNodesByType(nodeType string) ([]Node, error)
}

type NodeList struct {
	Nodes           []Node
	ContainsNonText bool
}

type BaseNode struct {
	Node
	MustBeFirst    bool
	ChildNodeLists []NodeList
	Token          string
}

type TextNode struct {
	BaseNode
	Text string
}

func (n *TextNode) String() string {
	return fmt.Sprintf("<TextNode: %s>", n.Text)
}

func (n *TextNode) Render() (string, error) {
	return n.Text, nil
}

func (n *TextNode) RenderAnnotated() (string, error) {
	return n.Text, nil
}

type VariableNode struct {
	BaseNode
	Name string
}
