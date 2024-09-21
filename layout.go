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

type LayoutNode struct {
	dimensions *Dimensions
	boxType    *StyledNode
	children   []*LayoutNode
}

func getLayoutTree(node *StyledNode) *LayoutNode {
	var layoutBoxChildren []*LayoutNode = []*LayoutNode{}

	for i := 0; i < len(node.children); i++ {
		layoutBoxChildren = append(layoutBoxChildren, getLayoutTree(node.children[i]))
	}

	return &LayoutNode{
		dimensions: &Dimensions{
			content: Rect{
				height: int(node.styles["height"].(Length)),
				width:  int(node.styles["width"].(Length)),
			},
			margin:  node.styles["margin"].(EdgeSizes),
			padding: node.styles["padding"].(EdgeSizes),
		},
		boxType:  node,
		children: layoutBoxChildren,
	}
}

func computeLayout(layoutTree *LayoutNode, viewportHeight, viewportWidth int) {
	computeWidth(layoutTree, viewportWidth)
	computeHeight(layoutTree, viewportHeight)
	computePosition(layoutTree, 0, 0)
}

func computeWidth(layoutNode *LayoutNode, parentWidth int) {
	if layoutNode.boxType.getDisplay() == DisplayBlock {
		computeBlockWidth(layoutNode, parentWidth)
	} else if layoutNode.boxType.getDisplay() == DisplayInline {
		computeInlineWidth(layoutNode, parentWidth)
	}
}

func computeBlockWidth(layoutNode *LayoutNode, parentWidth int) {
	if layoutNode.dimensions.content.width == 0 {
		layoutNode.dimensions.content.width = parentWidth
	} else {
		layoutNode.dimensions.content.width = min(layoutNode.dimensions.content.width, parentWidth)
	}

	for i := 0; i < len(layoutNode.children); i++ {
		computeWidth(layoutNode.children[i], layoutNode.dimensions.content.width)
	}
}

func computeInlineWidth(layoutNode *LayoutNode, parentWidth int) {
	width := parentWidth
	if layoutNode.dimensions.content.width != 0 {
		width = min(layoutNode.dimensions.content.width, parentWidth)
	}

	for i := 0; i < len(layoutNode.children); i++ {
		computeWidth(layoutNode.children[i], width)
	}

	if layoutNode.dimensions.content.width == 0 {
		childrenWidth := 0
		for i := 0; i < len(layoutNode.children); i++ {
			childrenWidth += layoutNode.children[i].dimensions.content.width
		}

		layoutNode.dimensions.content.width = min(childrenWidth, parentWidth)
	} else {
		layoutNode.dimensions.content.width = min(layoutNode.dimensions.content.width, parentWidth)
	}
}

func computeHeight(layoutNode *LayoutNode, parentHeight int) {
	height := parentHeight
	if layoutNode.dimensions.content.height != 0 {
		height = min(layoutNode.dimensions.content.height, parentHeight)
	}

	for i := 0; i < len(layoutNode.children); i++ {
		computeHeight(layoutNode.children[i], height)
	}

	if layoutNode.dimensions.content.height == 0 {
		width := layoutNode.dimensions.content.width
		blockHeight := 0
		inlineElements := []int{}
		filledWidth := 0

		for i := 0; i < len(layoutNode.children); i++ {
			child := layoutNode.children[i]
			childWidth := child.dimensions.content.width

			if child.boxType.getDisplay() == DisplayBlock {
				childWidth = width
			}

			if filledWidth+childWidth > width {
				maxHeight := 0
				for _, v := range inlineElements {
					maxHeight = max(maxHeight, v)
				}
				blockHeight += maxHeight

				inlineElements = []int{}
				filledWidth = 0
			}

			inlineElements = append(inlineElements, child.dimensions.content.height)
			filledWidth += childWidth
		}

		maxHeight := 0
		for _, v := range inlineElements {
			maxHeight = max(maxHeight, v)
		}
		blockHeight += maxHeight

		layoutNode.dimensions.content.height = min(blockHeight, parentHeight)
	}
}

func computePosition(layoutNode *LayoutNode, parentX int, parentY int) {
	layoutNode.dimensions.content.x = parentX
	layoutNode.dimensions.content.y = parentY

	width := layoutNode.dimensions.content.width
	blockHeight := 0
	inlineElements := []int{}
	filledWidth := 0

	for i := 0; i < len(layoutNode.children); i++ {
		child := layoutNode.children[i]
		childWidth := child.dimensions.content.width

		if child.boxType.getDisplay() == DisplayBlock {
			childWidth = width
		}

		if filledWidth+childWidth > width {
			maxHeight := 0
			for _, v := range inlineElements {
				maxHeight = max(maxHeight, v)
			}
			blockHeight += maxHeight

			inlineElements = []int{}
			filledWidth = 0
		}

		computePosition(
			layoutNode.children[i],
			parentX+filledWidth,
			parentY+blockHeight,
		)

		inlineElements = append(inlineElements, child.dimensions.content.height)
		filledWidth += childWidth
	}
}
