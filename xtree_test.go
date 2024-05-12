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
