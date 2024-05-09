package xtree

type XNode interface {
	GetId() any
	GetParentId() any
	SetChildren(children any)
	GetChildren() any
}

func FromList[T XNode](nodes []T) []T {
	nodeMap := map[any]T{}

	for _, node := range nodes {
		nodeMap[node.GetId()] = node
	}

	var treeNodes []T
	for _, node := range nodes {
		if node.GetParentId() != nil {
			if n, ok := nodeMap[node.GetParentId()]; ok {
				children := n.GetChildren()
				children = append(children.([]T), node)
				n.SetChildren(children)
			} else {
				treeNodes = append(treeNodes, node)
			}
		} else {
			treeNodes = append(treeNodes, node)
		}
	}

	return treeNodes
}

func ToListBFS[T XNode](treeNodes []T) []T {
	var nodes []T

	for len(treeNodes) > 0 {
		var shiftNode T
		shiftNode, treeNodes = treeNodes[0], treeNodes[1:]

		children := shiftNode.GetChildren().([]T)
		if len(children) > 0 {
			treeNodes = append(treeNodes, children...)
			shiftNode.SetChildren(nil)
		}

		nodes = append(nodes, shiftNode)
	}

	return nodes
}

func ToListDFS[T XNode](treeNodes []T) []T {
	var nodes []T

	for len(treeNodes) > 0 {
		var popNode T
		popNode, treeNodes = treeNodes[len(treeNodes)-1], treeNodes[:len(treeNodes)-1]

		children := popNode.GetChildren().([]T)
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
