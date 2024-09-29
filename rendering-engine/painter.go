package main

import (
	"image/color"

	"github.com/fogleman/gg"
)

func paintLayout(layoutTree *LayoutNode) {
	width := layoutTree.dimensions.content.width
	height := layoutTree.dimensions.content.height

	drawingContext := gg.NewContext(width, height)

	drawLayoutTree(layoutTree, drawingContext)

	drawingContext.SavePNG("browser.png")
}

func drawLayoutTree(layoutTree *LayoutNode, drawingContext *gg.Context) {
	switch layoutTree.boxType.node.nodeType.(type) {
	case TextNode:
		drawTextNodeLayoutTree(layoutTree, drawingContext)
	case ElementNode:
		drawElementNodeLayoutTree(layoutTree, drawingContext)
	}
}

func drawTextNodeLayoutTree(layoutTree *LayoutNode, drawingContext *gg.Context) {
	// fmt.Println("Name: " + layoutTree.boxType.node.nodeType.(TextNode).text)
	// fmt.Println("Dimensions: ", layoutTree.dimensions.content)
	// fmt.Println()

	drawingContext.SetColor(color.Black)
	drawingContext.DrawStringAnchored(
		layoutTree.boxType.node.nodeType.(TextNode).text,
		float64(layoutTree.dimensions.content.x),
		float64(layoutTree.dimensions.content.y),
		0,
		1,
	)
}

func drawElementNodeLayoutTree(layoutTree *LayoutNode, drawingContext *gg.Context) {
	content := layoutTree.dimensions.content
	drawingContext.DrawRectangle(float64(content.x), float64(content.y), float64(content.width), float64(content.height))

	color := layoutTree.boxType.styles["color"].(Color)
	drawingContext.SetRGB(float64(color.r)/256.0, float64(color.g)/256.0, float64(color.b)/256.0)

	drawingContext.Fill()

	// fmt.Println("Name: " + layoutTree.boxType.node.nodeType.(ElementNode).tagName)
	// fmt.Println("Color: %v", color)
	// fmt.Println("Dimensions: ", layoutTree.dimensions.content)
	// fmt.Println()

	for i := 0; i < len(layoutTree.children); i++ {
		drawLayoutTree(layoutTree.children[i], drawingContext)
	}
}
