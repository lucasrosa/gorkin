package feature

type filesPrimaryPort struct {
	repo ObjectSecondaryPort
}

func NewFilesService(repo ObjectSecondaryPort) FilesPrimaryPort {
	return &filesPrimaryPort{
		repo,
	}
}

func (filesprimaryport *filesPrimaryPort) Get(key string) (string, error) {
	return filesprimaryport.repo.GetObjectTemporaryURL(key)
}
