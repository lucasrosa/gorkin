package feature

type foldersPrimaryPort struct {
	repo ObjectSecondaryPort
}

// NewFolderService instantiates the file service with a ObjectSecondaryPort adapter
func NewFolderService(repo ObjectSecondaryPort) FolderPrimaryPort {
	return &foldersPrimaryPort{
		repo,
	}
}

func (foldersprimaryport *foldersPrimaryPort) GetAll() (Object, error) {
	return foldersprimaryport.repo.ListAllObjects()
}
