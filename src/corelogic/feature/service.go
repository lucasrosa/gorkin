package feature

type objectPort struct {
	repo ObjectSecondaryPort
}

func NewFolderService(repo ObjectSecondaryPort) FolderPrimaryPort {
	return &objectPort{
		repo,
	}
}

func(objectport *objectPort) GetAll() []Folder {
	return objectport.repo.GetAll()
}



// To be deprecated...
type dbPort struct {
	repo DatabaseSecondaryPort
}

func NewService(repo DatabaseSecondaryPort) FeaturePrimaryPort {
	return &dbPort{
		repo,
	}
}

func(dbport *dbPort) GetAll() []Feature {
	return dbport.repo.GetAll()
}