package task

type BackupTask struct {
}

func NewBackupTask() (*BackupTask, error) {
	return &BackupTask{}, nil
}
