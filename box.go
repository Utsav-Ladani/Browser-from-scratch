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
		dimensions: &Dimensions{
			content: Rect{
				height: int(node.specifiedValues["height"].(Length)),
				width:  int(node.specifiedValues["width"].(Length)),
			},
			margin:  node.specifiedValues["margin"].(EdgeSizes),
			padding: node.specifiedValues["padding"].(EdgeSizes),
		},
		boxType:  node,
		children: layoutBoxChildren,
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

	if layoutTree.dimensions.content.height < height {
		layoutTree.dimensions.content.height = height
	}

	if layoutTree.dimensions.content.width < width {
		layoutTree.dimensions.content.width = width
	}
}

func computePosition(layoutTree *LayoutBox, x, y int) {
	curX, curY := x, y
	prevInlineHeight := 0

	for i := 0; i < len(layoutTree.children); i++ {
		child := layoutTree.children[i]

		computePosition(child, curX, curY)

		dimensions := child.dimensions

		if child.boxType.getDisplay() == DisplayBlock {
			dimensions.content.x = x
			dimensions.content.y = curY + prevInlineHeight

			curX = x
			curY += prevInlineHeight + dimensions.content.height

			prevInlineHeight = 0
		} else if child.boxType.getDisplay() == DisplayInline {
			dimensions.content.x = curX
			dimensions.content.y = curY

			curX += dimensions.content.width

			if prevInlineHeight < dimensions.content.height {
				prevInlineHeight = dimensions.content.height
			}
		}
	}
}
