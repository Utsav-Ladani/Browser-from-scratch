package main

import (
	"github.com/fogleman/gg"
)

func paintLayout(layoutTree *LayoutBox) {
	width := layoutTree.dimensions.content.width
	height := layoutTree.dimensions.content.height

	drawingContext := gg.NewContext(width, height)

	drawLayoutTree(layoutTree, drawingContext)

	drawingContext.SavePNG("browser.png")
}

func drawLayoutTree(layoutTree *LayoutBox, drawingContext *gg.Context) {
	content := layoutTree.dimensions.content
	drawingContext.DrawRectangle(float64(content.x), float64(content.y), float64(content.width), float64(content.height))

	color := layoutTree.boxType.specifiedValues["color"].(Color)
	drawingContext.SetRGB(float64(color.r)/256.0, float64(color.g)/256.0, float64(color.b)/256.0)

	drawingContext.Fill()

	// fmt.Println("Name: " + layoutTree.boxType.node.nodeType.(ElementNode).tagName)
	// fmt.Println("Color: %v", color)
	// fmt.Println("Dimensions: %v", layoutTree.dimensions.content)

	for i := 0; i < len(layoutTree.children); i++ {
		drawLayoutTree(layoutTree.children[i], drawingContext)
	}
}
