package xtree

type XNode interface {
	GetId() any
	GetParentId() any
	SetChildren(children any)
	GetChildren() any
}

func Build[T XNode](nodes []T) []T {
	var treeNodes = []T{}

	if len(nodes) == 0 {
		return treeNodes
	}

	nodeMap := map[any]T{}
	for _, node := range nodes {
		id := node.GetId()
		if id != nil {
			nodeMap[id] = node
		}
	}

	for _, node := range nodes {
		parentId := node.GetParentId()
		if parentId != nil {
			if n, ok := nodeMap[parentId]; ok {
				if children := n.GetChildren(); children != nil {
					children = append(children.([]T), node)
					n.SetChildren(children)
				}
			} else {
				treeNodes = append(treeNodes, node)
			}
		} else {
			treeNodes = append(treeNodes, node)
		}
	}

	return treeNodes
}

func Flat[T XNode](treeNodes []T, dfs bool) []T {
	if dfs {
		return flatDFS(treeNodes)
	}
	return flatBFS(treeNodes)
}

func flatBFS[T XNode](treeNodes []T) []T {
	var nodes = []T{}

	for len(treeNodes) > 0 {
		var shiftNode T
		shiftNode, treeNodes = treeNodes[0], treeNodes[1:]

		var children []T
		_children := shiftNode.GetChildren()
		if _children != nil {
			children = _children.([]T)
		}
		if len(children) > 0 {
			treeNodes = append(treeNodes, children...)
			shiftNode.SetChildren(nil)
		}

		nodes = append(nodes, shiftNode)
	}

	return nodes
}

func flatDFS[T XNode](treeNodes []T) []T {
	var nodes = []T{}

	for len(treeNodes) > 0 {
		var popNode T
		popNode, treeNodes = treeNodes[len(treeNodes)-1], treeNodes[:len(treeNodes)-1]

		var children []T
		_children := popNode.GetChildren()
		if _children != nil {
			children = _children.([]T)
		}
		if len(children) > 0 {
			for i := len(children) - 1; i >= 0; i-- {
				treeNodes = append(treeNodes, children[i])
			}
			popNode.SetChildren(nil)
		}

		nodes = append(nodes, popNode)
	}

	return nodes
}

func Walk[T XNode](treeNodes []T, handler func(current T, parent *T)) {
	var stack = []T{}
	stack = append(stack, treeNodes...)

	i := 0
	for len(stack) > 0 {
		var popNode T
		popNode, stack = stack[0], stack[1:]

		if handler != nil && i < len(treeNodes) {
			handler(popNode, nil)
		}

		var children []T
		_children := popNode.GetChildren()
		if _children != nil {
			children = _children.([]T)
		}
		if len(children) > 0 {
			for j := len(children) - 1; j >= 0; j-- {
				stack = append(stack, children[j])

				if handler != nil {
					handler(children[j], &popNode)
				}
			}
		}

		i++
	}
}
