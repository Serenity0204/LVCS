package models

import (
	"encoding/json"
	"errors"
)

type treeNode struct {
	Value    string
	Children []*treeNode
}

type NaryTree struct {
	root *treeNode
}

func NewNaryTree() *NaryTree {
	return &NaryTree{
		root: nil,
	}
}

func (tree *NaryTree) GetNaryTreeRoot() *treeNode {
	return tree.root
}

func (tree *NaryTree) NaryTreeString() string {
	treeStr := "all commits:\n"
	treeStr += tree.buildNaryTreeString(tree.root, "", true, true)
	return treeStr
}

func (tree *NaryTree) buildNaryTreeString(node *treeNode, prefix string, isLastChild bool, isRoot bool) string {
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
		treeString += tree.buildNaryTreeString(child, prefix, isLast, false)
	}

	return treeString
}

func (tree *NaryTree) Insert(parent *treeNode, value string) error {
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
func (tree *NaryTree) Serialize() (string, error) {
	if tree.root == nil {
		return "", nil
	}
	data, err := json.Marshal(tree.root)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (tree *NaryTree) GetNode(val string) (*treeNode, error) {
	return tree.findNode(tree.root, val)
}

func (tree *NaryTree) findNode(node *treeNode, val string) (*treeNode, error) {
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
func (tree *NaryTree) Deserialize(data string) error {
	if len(data) == 0 {
		return nil
	}
	var root treeNode
	err := json.Unmarshal([]byte(data), &root)
	if err != nil {
		return err
	}
	tree.root = tree.copyNaryTree(&root)
	return nil
}

func (tree *NaryTree) copyNaryTree(node *treeNode) *treeNode {
	copyNode := &treeNode{Value: node.Value}
	copyNode.Children = make([]*treeNode, len(node.Children))
	for i, child := range node.Children {
		copyNode.Children[i] = tree.copyNaryTree(child)
	}
	return copyNode
}

func (tree *NaryTree) LCA(version1 string, version2 string) (string, error) {
	// Find the nodes corresponding to the given values
	node1, err := tree.GetNode(version1)
	if err != nil {
		return "", err
	}

	node2, err := tree.GetNode(version2)
	if err != nil {
		return "", err
	}

	// Find the LCA of the two nodes
	lcaValue := tree.findLCA(tree.root, node1, node2)

	if lcaValue == "" {
		return "", errors.New("LCA not found")
	}

	return lcaValue, nil
}

func (tree *NaryTree) findLCA(currentNode *treeNode, node1 *treeNode, node2 *treeNode) string {
	if currentNode == nil {
		return ""
	}

	// If the current node is either of the two nodes, it is the LCA
	if currentNode == node1 || currentNode == node2 {
		return currentNode.Value
	}

	// Recursively search for the LCA in each child
	lcaFound := false
	var lcaValue string
	for _, child := range currentNode.Children {
		found := tree.findLCA(child, node1, node2)
		if found != "" {
			if lcaFound {
				// If the LCA has already been found on a previous child, it means
				// the current node is the LCA of the two nodes.
				return currentNode.Value
			}
			lcaFound = true
			lcaValue = found
		}
	}
	return lcaValue
}
