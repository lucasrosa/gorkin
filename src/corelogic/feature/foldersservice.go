package feature

type foldersPrimaryPort struct {
	repo ObjectSecondaryPort
}

func NewFolderService(repo ObjectSecondaryPort) FolderPrimaryPort {
	return &foldersPrimaryPort{
		repo,
	}
}

func (foldersprimaryport *foldersPrimaryPort) GetAll() (Object, error) {
	return foldersprimaryport.repo.ListAllObjects()
}
