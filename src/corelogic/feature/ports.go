package feature

// FolderPrimaryPort is the entrypoint for the folders
type FolderPrimaryPort interface {
	GetAll() (Object, error)
	Get(folder string) (Folder, error)
}

// FilesPrimaryPort is the entrypoint for the files
type FilesPrimaryPort interface {
	Get(id string) (string, error)
}

type ObjectSecondaryPort interface {
	ListAllObjects() (Object, error)
	ListObjects(fodler string) (Folder, error)
	GetObjectTemporaryURL(id string) (string, error)
}
