package main

type Dimensions struct {
	content Rect
	margin  EdgeSizes
	padding EdgeSizes
}

type Rect struct {
	x, y, width, height int
}

type EdgeSizes struct {
	top, bottom, left, right int
}

type LayoutBox struct {
	dimensions *Dimensions
	boxType    *StyledNode
	children   []*LayoutBox
}

func getLayoutTree(node *StyledNode) *LayoutBox {
	var layoutBoxChildren []*LayoutBox = []*LayoutBox{}

	for i := 0; i < len(node.children); i++ {
		layoutBoxChildren = append(layoutBoxChildren, getLayoutTree(node.children[i]))
	}

	return &LayoutBox{
		dimensions: &Dimensions{},
		boxType:    node,
		children:   layoutBoxChildren,
	}
}

func computeLayout(layoutTree *LayoutBox) {
	computeHeightAndWidth(layoutTree)
	computePosition(layoutTree, 0, 0)
}

func computeHeightAndWidth(layoutTree *LayoutBox) {
	for i := 0; i < len(layoutTree.children); i++ {
		computeHeightAndWidth(layoutTree.children[i])
	}

	height, width := 0, 0
	blockWidth := 0

	for i := 0; i < len(layoutTree.children); i++ {
		child := layoutTree.children[i]
		childDimensions := child.dimensions

		if child.boxType.getDisplay() == DisplayBlock {
			height += childDimensions.content.height
			blockWidth = childDimensions.content.width
		} else if child.boxType.getDisplay() == DisplayInline {
			blockWidth += childDimensions.content.width
		}

		if blockWidth > width {
			width = blockWidth
		}
	}

	layoutTree.dimensions.content.height = height
	layoutTree.dimensions.content.width = width
}

func computePosition(layoutTree *LayoutBox, x, y int) {
	posX, posY := x, y

	for i := 0; i < len(layoutTree.children); i++ {
		child := layoutTree.children[i]

		computePosition(child, posX, posY)

		dimensions := child.dimensions
		dimensions.content.x = posX
		dimensions.content.y = posY

		if child.boxType.getDisplay() == DisplayBlock {
			posX = x
			posY += dimensions.content.height
		} else if child.boxType.getDisplay() == DisplayInline {
			posX += dimensions.content.width
		}
	}
}
