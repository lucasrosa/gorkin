package feature

// FolderPrimaryPort is the entrypoint for the folders
type FolderPrimaryPort interface {
	GetAll() (Object, error)
}

// FilesPrimaryPort is the entrypoint for the files
type FilesPrimaryPort interface {
	Get(key string) (string, error)
}

type ObjectSecondaryPort interface {
	ListAllObjects() (Object, error)
	GetObjectTemporaryURL(key string) (string, error)
}
