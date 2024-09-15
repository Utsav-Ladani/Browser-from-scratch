package main

func getHTML() *Node {
	body := getElementNode("body", map[string]string{"class": "body"})

	navigaton := getElementNode("nav", map[string]string{"class": "navigation"})

	link1 := getElementNode("link", map[string]string{"class": "link-nav"})
	link2 := getElementNode("link", map[string]string{"class": "link-nav"})
	link3 := getElementNode("link", map[string]string{"class": "link-nav"})
	link4 := getElementNode("link", map[string]string{"class": "link-nav"})

	div := getElementNode("div", map[string]string{"class": "dialog"})

	heading := getElementNode("h1", map[string]string{"class": "heading"})
	description := getElementNode("p", map[string]string{"class": "description"})
	input := getElementNode("input", map[string]string{"class": "input"})
	link := getElementNode("link", map[string]string{"class": "link"})
	button := getElementNode("button", map[string]string{"class": "button"})
	notice := getElementNode("p", map[string]string{"class": "notice"})

	navigaton.children = append(navigaton.children, link1, link2, link3, link4)
	div.children = append(div.children, heading, description, input, link, button, notice)
	body.children = append(body.children, navigaton, div)

	return body
}

func getStyleSheet() *StyleSheet {
	body := getRules("body", "", DisplayBlock, 740, 660, White, Margin0, Padding0)

	navigation := getRules("nav", "", DisplayBlock, 70, 660, Green, Margin0, Padding10)
	navigationLink := getRules("", "link-nav", DisplayInline, 30, 60, Blue, Margin10, Padding0)

	div := getRules("", "dialog", DisplayBlock, 650, 640, Blue, Margin10, Padding20)

	heading := getRules("", "heading", DisplayBlock, 100, 600, Black, Margin10, Padding0)
	description := getRules("", "description", DisplayBlock, 300, 600, Red, Margin10, Padding0)
	input := getRules("input", "", DisplayInline, 50, 300, Purple, Margin10, Padding0)
	link := getRules("", "link", DisplayInline, 50, 100, Green, Margin10, Padding0)
	button := getRules("button", "", DisplayInline, 50, 100, Yellow, Margin10, Padding0)
	notice := getRules("", "notice", DisplayBlock, 100, 600, Orange, Margin10, Padding0)

	return &StyleSheet{
		rules: []*Rule{body, navigation, navigationLink, div, heading, description, input, link, button, notice},
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
