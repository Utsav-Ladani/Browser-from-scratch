package main

type StyledNode struct {
	node            *Node
	children        []*StyledNode
	specifiedValues PropertyMap
}

type PropertyMap map[string]Value

func (styledNode *StyledNode) getDisplay() Value {
	return styledNode.specifiedValues["display"]
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

	specifiedValues := getSpecifiedValues(node, styleSheet)

	return &StyledNode{
		node:            node,
		children:        styledChildren,
		specifiedValues: specifiedValues,
	}
}

func getSpecifiedValues(node *Node, styleSheet *StyleSheet) PropertyMap {
	var propertyMap PropertyMap = make(PropertyMap)

	for i := 0; i < len(styleSheet.rules); i++ {
		rule := styleSheet.rules[i]
		for j := 0; j < len(rule.selectors); j++ {
			selector := rule.selectors[j]
			nodeData := node.nodeType.(ElementNode)

			if _, hasClass := nodeData.attributes[selector.class]; !hasClass && selector.tagName != nodeData.tagName {
				continue
			}

			for k := 0; k < len(rule.declarations); k++ {
				declaration := rule.declarations[k]
				propertyMap[declaration.name] = declaration.value
			}
		}
	}

	return propertyMap
}
