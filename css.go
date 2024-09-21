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

const (
	DisplayBlock  Keyword = "block"
	DisplayInline Keyword = "inline"
	DisplayNone   Keyword = ""
)

var (
	DodgerBlue  Color = Color{30, 144, 255, 255}
	White       Color = Color{255, 255, 255, 255}
	Yellow      Color = Color{255, 255, 0, 255}
	LightOrange Color = Color{255, 165, 0, 255}
	Purple      Color = Color{128, 0, 128, 255}
	Gray        Color = Color{128, 128, 128, 255}
	Pink        Color = Color{255, 192, 203, 255}
	LightGray   Color = Color{211, 211, 211, 255}
	LightGreen  Color = Color{144, 238, 144, 255}
	LightBlue   Color = Color{173, 216, 230, 255}
)

var (
	Margin0  EdgeSizes = EdgeSizes{}
	Margin10 EdgeSizes = EdgeSizes{10, 10, 10, 10}
	Margin20 EdgeSizes = EdgeSizes{20, 20, 20, 20}
	Margin40 EdgeSizes = EdgeSizes{40, 40, 40, 40}
)

var (
	Padding0  EdgeSizes = EdgeSizes{}
	Padding10 EdgeSizes = EdgeSizes{10, 10, 10, 10}
	Padding20 EdgeSizes = EdgeSizes{20, 20, 20, 20}
	Padding40 EdgeSizes = EdgeSizes{40, 40, 40, 40}
)

func getRules(tagName string, className string, display Keyword, height Length, width Length, color Color, margin EdgeSizes, padding EdgeSizes) *Rule {
	selector := &Selector{tagName: tagName, class: className}

	displayDeclaration := &Declaration{name: "display", value: display}
	heightDeclaration := &Declaration{name: "height", value: height}
	widthDeclaration := &Declaration{name: "width", value: width}
	colorDeclaration := &Declaration{name: "color", value: color}
	marginDeclaration := &Declaration{name: "margin", value: margin}
	paddingDeclaration := &Declaration{name: "padding", value: padding}

	var selectors []*Selector = []*Selector{selector}
	var declarations []*Declaration = []*Declaration{
		displayDeclaration,
		heightDeclaration,
		widthDeclaration,
		colorDeclaration,
		marginDeclaration,
		paddingDeclaration,
	}

	return &Rule{
		selectors:    selectors,
		declarations: declarations,
	}
}
