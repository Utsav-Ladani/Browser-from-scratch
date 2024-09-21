package main

type StyledNode struct {
	node     *Node
	children []*StyledNode
	styles   StyleMap
}

type StyleMap map[string]Value

func (styledNode *StyledNode) getDisplay() Value {
	return styledNode.styles["display"]
}

func getStyledTree(node *Node, styleSheet *StyleSheet) *StyledNode {
	if _, ok := node.nodeType.(ElementNode); !ok {
		return &StyledNode{
			node: node,
		}
	}

	var styledChildren []*StyledNode = []*StyledNode{}

	for i := 0; i < len(node.children); i++ {
		styledChildren = append(styledChildren, getStyledTree(node.children[i], styleSheet))
	}

	styles := getStyles(node, styleSheet)

	return &StyledNode{
		node:     node,
		children: styledChildren,
		styles:   styles,
	}
}

func getStyles(node *Node, styleSheet *StyleSheet) StyleMap {
	var styles StyleMap = make(StyleMap)

	for i := 0; i < len(styleSheet.rules); i++ {
		rule := styleSheet.rules[i]

		for j := 0; j < len(rule.selectors); j++ {
			selector := rule.selectors[j]
			nodeData := node.nodeType.(ElementNode)

			if selector.class != nodeData.attributes["class"] && selector.tagName != nodeData.tagName {
				continue
			}

			for k := 0; k < len(rule.declarations); k++ {
				declaration := rule.declarations[k]
				styles[declaration.name] = declaration.value
			}
		}
	}

	return styles
}
