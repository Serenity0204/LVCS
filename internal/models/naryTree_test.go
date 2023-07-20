package models

import (
	"fmt"
	"testing"
)

const debug bool = false

func TestNaryTree(t *testing.T) {
	tree := NewNaryTree()
	if tree.root != nil {
		t.Errorf("expected root to be nil but it's not nil")
	}

	root := tree.GetNaryTreeRoot()
	if root != nil {
		t.Errorf("GetDirectoryTreeRoot failed. Expected nil root for empty tree. Got root='%v'", root)
	}

	// Test Insert
	err := tree.Insert(nil, "v0")
	if err != nil {
		t.Errorf("Insert failed. Expected error for nil parent, but got nil.")
	}

	parent := tree.GetNaryTreeRoot()
	err = tree.Insert(parent, "v1")
	if err != nil {
		t.Errorf("Insert failed. Expected no error, but got: %v", err)
	}
	if len(parent.Children) != 1 || parent.Children[0].Value != "v1" {
		t.Errorf("Insert failed. Parent's children are not updated correctly.")
	}

	parent, err = tree.GetNode("v0")
	if err != nil {
		t.Errorf("expected v1 exists but not found")
	}
	err = tree.Insert(parent, "v2")
	if err != nil {
		t.Errorf("Insert failed. Expected no error, but got: %v", err)
	}
	err = tree.Insert(parent, "v3")
	if err != nil {
		t.Errorf("Insert failed. Expected no error, but got: %v", err)
	}

	parent, err = tree.GetNode("v2")
	if err != nil {
		t.Errorf("expected v2 exists but not found")
	}

	err = tree.Insert(parent, "v4")
	if err != nil {
		t.Errorf("Insert failed. Expected no error, but got: %v", err)
	}

	// Test Serialize
	serializedData, err := tree.Serialize()
	if err != nil {
		t.Errorf("Serialize failed. Expected no error, but got: %v", err)
	}

	// Test Deserialize
	err = tree.Deserialize(serializedData)
	if err != nil {
		t.Errorf("Deserialize failed. Expected no error, but got: %v", err)
	}
	if debug {
		treeString := tree.NaryTreeString()
		fmt.Println(treeString)
	}
}
