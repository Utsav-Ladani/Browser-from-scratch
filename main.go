package main

func getHTML() *Node {
	div := getElementNode("div", map[string]string{"class": "dialog"})

	heading := getElementNode("h1", map[string]string{"class": "heading"})
	description := getElementNode("p", map[string]string{"class": "description"})
	input := getElementNode("input", map[string]string{"class": "input"})
	link := getElementNode("link", map[string]string{"class": "link"})
	button := getElementNode("button", map[string]string{"class": "button"})
	notice := getElementNode("p", map[string]string{"class": "notice"})

	div.children = append(div.children, heading, description, input, link, button, notice)

	return div
}

func getStyleSheet() *StyleSheet {
	div := getRules("", "dialog", DisplayBlock, 600, 400, Yellow, Margin0, Padding20)
	heading := getRules("", "heading", DisplayBlock, 100, 200, Black, Margin20, Padding0)
	description := getRules("", "description", DisplayBlock, 300, 200, Red, Margin10, Padding0)
	input := getRules("input", "", DisplayInline, 30, 100, Purple, Margin0, Padding0)
	link := getRules("link", "", DisplayInline, 30, 40, Green, Margin10, Padding0)
	button := getRules("button", "", DisplayInline, 30, 50, Blue, Margin10, Padding0)
	notice := getRules("", "notice", DisplayBlock, 60, 150, Orange, Margin0, Padding0)

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
	paintLayout(layoutTree)
}
