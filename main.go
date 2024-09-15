package main

import "fmt"

func getHTML() *Node {
	heading := getElementNode("h1", map[string]string{"class": "heading"})
	description := getElementNode("p", map[string]string{"class": "description"})
	input := getElementNode("input", map[string]string{"class": "input"})
	link := getElementNode("link", map[string]string{"class": "link"})
	button := getElementNode("button", map[string]string{"class": "button"})
	notice := getElementNode("p", map[string]string{"class": "notice-info"})

	div := getElementNode("div", map[string]string{"class": "div"})
	div.children = append(div.children, heading, description, input, link, button, notice)

	fmt.Println(div)

	return div
}

func getStyleSheet() *StyleSheet {
	div := getRules("div", "dialog", DisplayBlock, 600, 400, Yellow)
	heading := getRules("h1", "heaing", DisplayBlock, 100, 200, Black)
	description := getRules("p", "", DisplayBlock, 300, 200, Red)
	input := getRules("input", "", DisplayInline, 30, 100, Purple)
	link := getRules("link", "", DisplayInline, 30, 40, Green)
	button := getRules("button", "", DisplayInline, 30, 50, Green)
	notice := getRules("notice", "", DisplayBlock, 60, 150, Orange)

	return &StyleSheet{
		rules: []*Rule{div, heading, description, input, link, button, notice},
	}
}

func main() {
	var rootNode *Node = getHTML()
	var styleSheet *StyleSheet = getStyleSheet()
	var styledTree *StyledNode = getStyledTree(rootNode, styleSheet)
	var layoutTree *LayoutBox = getLayoutTree(styledTree)
	computeLayout(layoutTree)

	fmt.Println(layoutTree)
}
