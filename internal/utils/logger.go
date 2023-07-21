package utils

type LVCSLoggerManager struct {
	lvcsBaseManager
}

// // creates a new LVCSLoggerManager instance
// func NewLVCSLoggerManager(lvcsPath string) *LVCSLoggerManager {
// 	return &LVCSLoggerManager{
// 		lvcsBaseManager: newLVCSBaseManager(lvcsPath),
// 	}
// }

// func (lvcsLogger *LVCSLoggerManager) RemoveStageContent() error {
// 	stageFile, err := os.OpenFile(lvcsLogger.lvcsStagePath, os.O_WRONLY|os.O_TRUNC, 0644)
// 	if err != nil {
// 		return err
// 	}
// 	defer stageFile.Close()
// 	return nil
// }

// func (lvcsLogger *LVCSLoggerManager) GetStageContent() (string, error) {
// 	return "", nil
// }
