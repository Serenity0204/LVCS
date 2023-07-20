package utils

type lvcsBaseManager struct {
	lvcsPath           string
	lvcsObjPath        string
	lvcsCommitPath     string
	lvcsStagePath      string
	lvcsCurrentRefPath string
	lvcsTreePath       string
}

func newLVCSBaseManager(lvcsPath string) lvcsBaseManager {
	return lvcsBaseManager{
		lvcsPath:           lvcsPath,
		lvcsObjPath:        lvcsPath + "/objects",
		lvcsCommitPath:     lvcsPath + "/commits",
		lvcsTreePath:       lvcsPath + "/trees",
		lvcsStagePath:      lvcsPath + "/stage.txt",
		lvcsCurrentRefPath: lvcsPath + "/currentRef.txt",
	}
}
