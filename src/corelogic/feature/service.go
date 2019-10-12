package feature

type foldersPrimaryPort struct {
	repo ObjectSecondaryPort
}

type filesPrimaryPort struct {
	repo ObjectSecondaryPort
}

func NewFolderService(repo ObjectSecondaryPort) FolderPrimaryPort {
	return &foldersPrimaryPort{
		repo,
	}
}

func NewFilesService(repo ObjectSecondaryPort) FilesPrimaryPort {
	return &filesPrimaryPort{
		repo,
	}
}

func (foldersprimaryport *foldersPrimaryPort) GetAll() (Object, error) {
	return foldersprimaryport.repo.ListAllObjects()
}

func (foldersprimaryport *foldersPrimaryPort) Get(folder string) (Folder, error) {
	return foldersprimaryport.repo.ListObjects(folder)
}

func (filesprimaryport *filesPrimaryPort) Get(id string) (string, error) {
	return filesprimaryport.repo.GetObjectTemporaryURL(id)
}
