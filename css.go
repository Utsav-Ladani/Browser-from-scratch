package main

type StyleSheet struct {
	rules []*Rule
}

type Rule struct {
	selectors    []*Selector
	declarations []*Declaration
}

type Selector struct {
	tagName string
	class   string
}

type Declaration struct {
	name  string
	value Value
}

type Value interface{}

type Keyword string
type Length int
type ColorValue Color

type Color struct {
	r, g, b, a int
}

var (
	Red    Color = Color{255, 0, 0, 255}
	Green  Color = Color{0, 255, 0, 255}
	Blue   Color = Color{0, 0, 255, 255}
	Black  Color = Color{0, 0, 0, 255}
	White  Color = Color{255, 255, 255, 255}
	Yellow Color = Color{255, 255, 0, 255}
	Orange Color = Color{255, 165, 0, 255}
	Purple Color = Color{128, 0, 128, 255}
)

const (
	DisplayBlock  Keyword = "block"
	DisplayInline Keyword = "inline"
	DisplayNone   Keyword = ""
)

func getRules(tagName string, className string, display Keyword, height Length, width Length, color Color) *Rule {
	selector := &Selector{tagName: tagName, class: className}

	displayDeclaration := &Declaration{name: "display", value: display}
	heightDeclaration := &Declaration{name: "height", value: height}
	widthDeclaration := &Declaration{name: "width", value: width}
	colorDeclaration := &Declaration{name: "color", value: color}

	var selectors []*Selector = []*Selector{selector}
	var declarations []*Declaration = []*Declaration{
		displayDeclaration,
		heightDeclaration,
		widthDeclaration,
		colorDeclaration,
	}

	return &Rule{
		selectors:    selectors,
		declarations: declarations,
	}
}
