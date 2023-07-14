package main

// import (
// 	"fmt"

// 	"github.com/Serenity0204/LVCS/helper"
// )

func main() {
	// if helper.AlreadyInit(helper.LvcsDir) {
	// 	fmt.Println("Already exists")
	// 	return
	// }

	// // add files
	// err := helper.Init(helper.LvcsDir)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// path := "test_data/a.txt"
	// err = helper.Add(path, helper.LvcsDir)
	// if err != nil {
	// 	fmt.Printf("Failed to add %s", path)
	// 	return
	// }
	// path = "test_data/b.txt"
	// err = helper.Add(path, helper.LvcsDir)
	// if err != nil {
	// 	fmt.Printf("Failed to add %s", path)
	// 	return
	// }

	// path = "test_data/ok/abc.txt"
	// err = helper.Add(path, helper.LvcsDir)
	// if err != nil {
	// 	fmt.Printf("Failed to add %s", path)
	// 	return
	// }

	// // create branch
	// master := "master"
	// commitFolderPath := helper.LvcsDir + "/commits/"
	// if !helper.BranchExists(commitFolderPath, master) {
	// 	err = helper.CreateBranch(commitFolderPath, master)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// }
	// // commit
	// err = helper.Commit(helper.LvcsDir, master)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// // remove stage content
	// err = helper.RemoveStageContent(helper.LvcsDir)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// // add again
	// path = "test_data/b.txt"
	// err = helper.Add(path, helper.LvcsDir)
	// if err != nil {
	// 	fmt.Printf("Failed to add %s", path)
	// 	return
	// }

	// path = "test_data/ok/abc.txt"
	// err = helper.Add(path, helper.LvcsDir)
	// if err != nil {
	// 	fmt.Printf("Failed to add %s", path)
	// 	return
	// }

	// // commit again
	// err = helper.Commit(helper.LvcsDir, master)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
}
