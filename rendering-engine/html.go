package main

import (
	"github.com/fogleman/gg"
)

type Node struct {
	children []*Node
	nodeType NodeType
}

type NodeType interface {
	getDimensions() *Dimensions
}

type TextNode struct {
	text string
}

func (textNode TextNode) getDimensions() *Dimensions {
	drawingContext := gg.NewContext(100, 200)
	width, height := drawingContext.MeasureString(textNode.text)

	return &Dimensions{
		content: Rect{
			height: int(height),
			width:  int(width),
		},
	}
}

type ElementNode struct {
	tagName    string
	attributes map[string]string
}

func (elementNode ElementNode) getDimensions() *Dimensions {
	return &Dimensions{}
}

func getTextNode(text string) *Node {
	return &Node{nodeType: TextNode{text: text}}
}

func getElementNode(tagName string, attributes map[string]string) *Node {
	return &Node{nodeType: ElementNode{tagName: tagName, attributes: attributes}, children: []*Node{}}
}

func getElementNodeWithText(tagName string, attributes map[string]string, text string) *Node {
	textNode := getTextNode(text)
	return &Node{nodeType: ElementNode{tagName: tagName, attributes: attributes}, children: []*Node{textNode}}
}
