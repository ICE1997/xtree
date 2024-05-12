package xtree

import (
	"reflect"
	"testing"
)

type MockNode struct {
	Id       int
	ParentId int
	Children []*MockNode
}

func (m *MockNode) GetId() any {
	return m.Id
}

func (m *MockNode) GetParentId() any {
	return m.ParentId
}

func (m *MockNode) SetChildren(children any) {
	if children == nil {
		m.Children = nil
		return
	}
	m.Children = children.([]*MockNode)
}

func (m *MockNode) GetChildren() any {
	return m.Children
}

func TestFromList(t *testing.T) {
	// Mock data
	nodes := []*MockNode{
		{Id: 2, ParentId: 1},
		{Id: 1, ParentId: 0},
		{Id: 3, ParentId: 1},
		{Id: 5, ParentId: 3},
		{Id: 4, ParentId: 2},
	}

	// Expected result
	expectedTree := []*MockNode{
		{Id: 1, ParentId: 0, Children: []*MockNode{
			{Id: 2, ParentId: 1, Children: []*MockNode{
				{Id: 4, ParentId: 2},
			}},
			{Id: 3, ParentId: 1, Children: []*MockNode{
				{Id: 5, ParentId: 3},
			}},
		}},
	}

	// Call the Build function
	tree := Build(nodes)

	// Check if the result matches the expected tree
	if !reflect.DeepEqual(tree, expectedTree) {
		t.Errorf("Build() returned incorrect result, expected: %v, got: %v", expectedTree, tree)
	}
}

func TestToListBFS(t *testing.T) {
	// Mock data
	treeNodes := []*MockNode{
		{Id: 1, ParentId: 0, Children: []*MockNode{
			{Id: 2, ParentId: 1, Children: []*MockNode{
				{Id: 4, ParentId: 2},
			}},
			{Id: 3, ParentId: 1, Children: []*MockNode{
				{Id: 5, ParentId: 3},
			}},
		}},
	}

	// Expected result
	expectedList := []*MockNode{
		{Id: 1, ParentId: 0},
		{Id: 2, ParentId: 1},
		{Id: 3, ParentId: 1},
		{Id: 4, ParentId: 2},
		{Id: 5, ParentId: 3},
	}

	// Call the flatBFS function
	list := flatBFS(treeNodes)

	// Check if the result matches the expected tree
	if !reflect.DeepEqual(list, expectedList) {
		t.Errorf("flatBFS() returned incorrect result, expected: %v, got: %v", expectedList, list)
	}
}

func TestToListDFS(t *testing.T) {
	// Mock data
	treeNodes := []*MockNode{
		{Id: 1, ParentId: 0, Children: []*MockNode{
			{Id: 2, ParentId: 1, Children: []*MockNode{
				{Id: 4, ParentId: 2},
			}},
			{Id: 3, ParentId: 1, Children: []*MockNode{
				{Id: 5, ParentId: 3},
			}},
		}},
	}

	// Expected result
	expectedList := []*MockNode{
		{Id: 1, ParentId: 0},
		{Id: 2, ParentId: 1},
		{Id: 4, ParentId: 2},
		{Id: 3, ParentId: 1},
		{Id: 5, ParentId: 3},
	}

	// Call the flatDFS function
	list := flatDFS(treeNodes)

	// Check if the result matches the expected tree
	if !reflect.DeepEqual(list, expectedList) {
		t.Errorf("flatDFS() returned incorrect result, expected: %v, got: %v", expectedList, list)
	}
}

func TestWalk(t *testing.T) {
	// Mock data
	treeNodes := []*MockNode{
		{Id: 1, ParentId: 0, Children: []*MockNode{
			{Id: 2, ParentId: 1, Children: []*MockNode{
				{Id: 4, ParentId: 2},
				{Id: 6, ParentId: 2, Children: []*MockNode{
					{Id: 7, ParentId: 6, Children: []*MockNode{
						{Id: 10, ParentId: 7},
						{Id: 11, ParentId: 7},
					}},
					{Id: 8, ParentId: 6},
					{Id: 9, ParentId: 6},
				}},
			}},
			{Id: 3, ParentId: 1, Children: []*MockNode{
				{Id: 5, ParentId: 3},
				{Id: 12, ParentId: 3, Children: []*MockNode{
					{Id: 13, ParentId: 12},
					{Id: 14, ParentId: 12},
				}},
			}},
		}},
		{Id: 15, ParentId: 0, Children: []*MockNode{
			{Id: 16, ParentId: 15},
			{Id: 17, ParentId: 15},
			{Id: 18, ParentId: 15},
		}},
	}

	parentIdMap := map[int]int{
		1:  0,
		2:  1,
		3:  1,
		4:  2,
		5:  3,
		6:  2,
		7:  6,
		8:  6,
		9:  6,
		10: 7,
		11: 7,
		12: 3,
		13: 12,
		14: 12,
		15: 0,
		16: 15,
		17: 15,
		18: 15,
	}

	Walk(treeNodes, func(current *MockNode, parent **MockNode) {
		if parent == nil {
			if current.ParentId != 0 {
				t.Errorf("walk() returned incorrect result, expected: %v, got: %v\"", 0, current.ParentId)
			}
		} else {
			if parentIdMap[current.Id] != (*parent).Id {
				t.Errorf("walk() returned incorrect result, expected: %v, got: %v\"", parentIdMap[current.Id], current.ParentId)
			}
		}
	})
}
