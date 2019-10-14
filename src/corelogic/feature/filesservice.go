package feature

type filesPrimaryPort struct {
	repo ObjectSecondaryPort
}

// NewFilesService instantiates the file service with a ObjectSecondaryPort adapter
func NewFilesService(repo ObjectSecondaryPort) FilesPrimaryPort {
	return &filesPrimaryPort{
		repo,
	}
}

func (filesprimaryport *filesPrimaryPort) Get(key string) (string, error) {
	return filesprimaryport.repo.GetObjectTemporaryURL(key)
}
