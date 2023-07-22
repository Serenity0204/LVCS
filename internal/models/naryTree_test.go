package models_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/Serenity0204/LVCS/internal/models"
)

const debug bool = false

func TestNaryTree(t *testing.T) {
	tree := models.NewNaryTree()
	if tree.GetNaryTreeRoot() != nil {
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

func buildTestTree() (*models.NaryTree, error) {
	tree := models.NewNaryTree()
	err := tree.Insert(nil, "v0")
	if err != nil {
		return nil, err
	}
	parent, err := tree.GetNode("v0")
	if err != nil {
		return nil, err
	}
	for i := 1; i <= 4; i++ {
		err = tree.Insert(parent, "v"+strconv.Itoa(i))
		if err != nil {
			return nil, err
		}
	}
	child := 5
	for i := 1; i <= 4; i++ {
		parent, err = tree.GetNode("v" + strconv.Itoa(i))
		if err != nil {
			return nil, err
		}
		for j := 0; j < 2; j++ {
			err = tree.Insert(parent, "v"+strconv.Itoa(child))
			if err != nil {
				return nil, err
			}
			child++
		}
	}
	parent, err = tree.GetNode("v11")
	if err != nil {
		return nil, err
	}

	for i := 0; i < 3; i++ {
		err = tree.Insert(parent, "v"+strconv.Itoa(child))
		if err != nil {
			return nil, err
		}
		child++
	}
	if debug {
		fmt.Println(tree.NaryTreeString())
	}
	return tree, nil
}

func TestNaryTreeLCA(t *testing.T) {
	tree, err := buildTestTree()
	if err != nil {
		t.Errorf(err.Error())
	}

	// Equal Node
	version, err := tree.LCA("v4", "v4")
	if err != nil {
		t.Errorf(err.Error())
	}
	if version != "v4" {
		t.Errorf("version not equal to v4")
	}

	// General Cases
	version, err = tree.LCA("v5", "v6")
	if err != nil {
		t.Errorf(err.Error())
	}
	if version != "v1" {
		t.Errorf("version not equal to v1")
	}
	version, err = tree.LCA("v2", "v8")
	if err != nil {
		t.Errorf(err.Error())
	}
	if version != "v2" {
		t.Errorf("version not equal to v2")
	}

	version, err = tree.LCA("v12", "v10")
	if err != nil {
		t.Errorf(err.Error())
	}
	if version != "v0" {
		t.Errorf("version not equal to v0")
	}

	version, err = tree.LCA("v15", "v12")
	if err != nil {
		t.Errorf(err.Error())
	}
	if version != "v4" {
		t.Errorf("version not equal to v4")
	}

	// Non existing Nodes
	version, err = tree.LCA("v666", "v1")
	if err == nil {
		t.Errorf("expected error to be nil but not none nil")
	}
	version, err = tree.LCA("v6", "v111")
	if err == nil {
		t.Errorf("expected error to be nil but not none nil")
	}

	// Empty tree
	tree = models.NewNaryTree()
	version, err = tree.LCA("v0", "v0")
	if err == nil {
		t.Errorf("expected error to be nil but not none nil")
	}
}
