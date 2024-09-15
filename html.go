package main

type Node struct {
	children []*Node
	nodeType NodeType
}

type NodeType interface{}

type ElementNode struct {
	tagName    string
	attributes map[string]string
}

func getElementNode(tagName string, attributes map[string]string) *Node {
	return &Node{nodeType: ElementNode{tagName: tagName, attributes: attributes}, children: []*Node{}}
}
