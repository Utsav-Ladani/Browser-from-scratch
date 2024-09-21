package main

import (
	"time"
)

const TimeFormat string = "2 Jan 2006 3:4:5 PM MST"

func getHTML() *Node {
	body := getElementNode("body", map[string]string{"class": "body"})

	navigaton := getElementNode("nav", map[string]string{"class": "navigation"})

	link1 := getElementNodeWithText("link", map[string]string{"class": "link-nav-1"}, "Home")
	link2 := getElementNodeWithText("link", map[string]string{"class": "link-nav-2"}, "About")
	link3 := getElementNodeWithText("link", map[string]string{"class": "link-nav-3"}, "Career")
	link4 := getElementNodeWithText("link", map[string]string{"class": "link-nav-4"}, "Product")

	div := getElementNode("div", map[string]string{"class": "dialog"})

	heading := getElementNodeWithText("h1", map[string]string{"class": "heading"}, "Heading")
	timer := getElementNodeWithText("p", map[string]string{"class": "timer"}, "Last Updated At: "+time.Now().Format(TimeFormat))
	description := getElementNodeWithText("p", map[string]string{"class": "description"}, "Description text lorem ipsum dolor sit amet consectetur adipiscing elit")
	input := getElementNodeWithText("input", map[string]string{"class": "input"}, "Enter you name here")
	link := getElementNodeWithText("link", map[string]string{"class": "body-link"}, "Visit Here")
	button := getElementNodeWithText("button", map[string]string{"class": "button"}, "Click Me")
	notice := getElementNodeWithText("p", map[string]string{"class": "notice"}, "Copyright")

	navigaton.children = append(navigaton.children, link1, link2, link3, link4)
	div.children = append(div.children, heading, timer, description, input, link, button, notice)
	body.children = append(body.children, navigaton, div)

	return body
}

func getStyleSheet() *StyleSheet {
	body := getRules("body", "", DisplayBlock, 740, 660, White, Margin0, Padding0)

	navigation := getRules("nav", "", DisplayBlock, 0, 0, LightGreen, Margin0, Padding10)
	navigationLink := getRules("link", "", DisplayInline, 40, 60, DodgerBlue, Margin10, Padding0)
	navigationLink1 := getRules("", "link-nav-1", DisplayInline, 40, 60, Gray, Margin10, Padding0)
	navigationLink2 := getRules("", "link-nav-3", DisplayInline, 40, 60, Gray, Margin10, Padding0)

	div := getRules("", "dialog", DisplayBlock, 0, 0, LightGray, Margin10, Padding20)

	heading := getRules("", "heading", DisplayBlock, 50, 0, Yellow, Margin10, Padding0)
	timer := getRules("", "timer", DisplayBlock, 30, 0, LightBlue, Margin10, Padding0)
	description := getRules("", "description", DisplayBlock, 300, 0, Pink, Margin10, Padding0)
	input := getRules("input", "", DisplayInline, 50, 300, Gray, Margin10, Padding0)
	link := getRules("", "body-link", DisplayInline, 50, 100, LightGreen, Margin10, Padding0)
	button := getRules("button", "", DisplayInline, 50, 100, Yellow, Margin10, Padding0)
	notice := getRules("", "notice", DisplayBlock, 100, 0, LightOrange, Margin10, Padding0)

	return &StyleSheet{
		rules: []*Rule{body, navigation, navigationLink, navigationLink1, navigationLink2, div, heading, timer, description, input, link, button, notice},
	}
}

var htmlTree *Node = getHTML()
var styleSheet *StyleSheet = getStyleSheet()

func render() {
	var styledTree *StyledNode = getStyledTree(htmlTree, styleSheet)
	var layoutTree *LayoutNode = getLayoutTree(styledTree)
	computeLayout(layoutTree, 800, 1000)
	paintLayout(layoutTree)
}

func updateTimer() {
	currentTime := time.Now().Format(TimeFormat)
	htmlTree.children[1].children[1].children[0].nodeType = TextNode{text: "Last Updated At: " + currentTime}
}

func main() {
	render()

	time.Sleep(5 * time.Second)
	updateTimer()

	render()
}
