package models

import (
	"encoding/json"
	"errors"
)

type treeNode struct {
	Value    string
	Children []*treeNode
}

type DirectoryTree struct {
	lvcsPath string
	root     *treeNode
}

func NewDirectoryTree(lvcsPath string) *DirectoryTree {
	return &DirectoryTree{
		lvcsPath: lvcsPath,
		root:     nil,
	}
}

func (tree *DirectoryTree) GetDirectoryTreeRoot() *treeNode {
	return tree.root
}

func (tree *DirectoryTree) DirectoryTreeString() string {
	treeStr := "all commits:\n"
	treeStr += tree.buildDirectoryTreeString(tree.root, "", true, true)
	return treeStr
}

func (tree *DirectoryTree) buildDirectoryTreeString(node *treeNode, prefix string, isLastChild bool, isRoot bool) string {
	treeString := ""

	if !isRoot {
		if isLastChild {
			treeString += prefix + "└── "
			prefix += "    "
		} else {
			treeString += prefix + "├── "
			prefix += "|   "
		}
	}

	treeString += node.Value + "\n"

	childCount := len(node.Children)
	for i, child := range node.Children {
		isLast := i == childCount-1
		treeString += tree.buildDirectoryTreeString(child, prefix, isLast, false)
	}

	return treeString
}

func (tree *DirectoryTree) Insert(parent *treeNode, value string) error {
	// if empty tree
	if tree.root == nil {
		tree.root = &treeNode{Value: "v0"}
		return nil
	}
	node := &treeNode{Value: value}
	if parent == nil {
		return errors.New("parent is nil")
	}
	parent.Children = append(parent.Children, node)
	return nil
}

// Serialize serializes an N-ary tree to a JSON string
func (tree *DirectoryTree) Serialize() (string, error) {
	if tree.root == nil {
		return "", nil
	}
	data, err := json.Marshal(tree.root)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (tree *DirectoryTree) GetNode(val string) (*treeNode, error) {
	return tree.findNode(tree.root, val)
}

func (tree *DirectoryTree) findNode(node *treeNode, val string) (*treeNode, error) {
	if node == nil {
		return nil, errors.New("node not found")
	}

	if node.Value == val {
		return node, nil
	}

	for _, child := range node.Children {
		foundNode, err := tree.findNode(child, val)
		if err == nil {
			return foundNode, nil
		}
	}

	return nil, errors.New("node not found")
}

// Deserialize deserializes a JSON string to an N-ary tree
func (tree *DirectoryTree) Deserialize(data string) error {
	var root treeNode
	err := json.Unmarshal([]byte(data), &root)
	if err != nil {
		return err
	}
	tree.root = tree.copyDirectoryTree(&root)
	return nil
}

func (tree *DirectoryTree) copyDirectoryTree(node *treeNode) *treeNode {
	copyNode := &treeNode{Value: node.Value}
	copyNode.Children = make([]*treeNode, len(node.Children))
	for i, child := range node.Children {
		copyNode.Children[i] = tree.copyDirectoryTree(child)
	}
	return copyNode
}
