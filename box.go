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
	computePosition(layoutTree, 0, 0)
}

func computePosition(layoutTree *LayoutBox, x, y int) {
	x = x + layoutTree.dimensions.padding.left
	y = y + layoutTree.dimensions.padding.top

	layoutTree.dimensions.content.x = x
	layoutTree.dimensions.content.y = y

	curX, curY := x, y
	prevInlineHeight := 0

	for i := 0; i < len(layoutTree.children); i++ {
		child := layoutTree.children[i]

		computePosition(child, curX, curY)

		dimensions := child.dimensions

		if child.boxType.getDisplay() == DisplayBlock {
			dimensions.content.x = x + child.dimensions.margin.left
			dimensions.content.y = curY + prevInlineHeight + child.dimensions.margin.top

			curX = x + child.dimensions.margin.left + child.dimensions.margin.right
			curY += prevInlineHeight + dimensions.content.height + child.dimensions.margin.top + child.dimensions.margin.bottom

			prevInlineHeight = 0
		} else if child.boxType.getDisplay() == DisplayInline {
			dimensions.content.x = curX + child.dimensions.margin.left
			dimensions.content.y = curY + child.dimensions.margin.top

			curX += dimensions.content.width + child.dimensions.margin.left + child.dimensions.margin.right

			if prevInlineHeight < dimensions.content.height+child.dimensions.margin.top+child.dimensions.margin.bottom {
				prevInlineHeight = dimensions.content.height + child.dimensions.margin.top + child.dimensions.margin.bottom
			}
		}
	}
}
