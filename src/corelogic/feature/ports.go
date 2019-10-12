package feature

// FeaturePrimaryPort is the entrypoint for the feature Package
type FeaturePrimaryPort interface {
	GetAll() []Feature
}

// FolderPrimaryPort is the entrypoint for the folders
type FolderPrimaryPort interface {
	GetAll() (Object, error)
	Get(folder string) (Folder, error)
}

// FilesPrimaryPort is the entrypoint for the files
type FilesPrimaryPort interface {
	Get(id string) (string, error)
}

// DatabaseSecondaryPort is the way the business rules communicate to the external world
type DatabaseSecondaryPort interface {
	GetAll() []Feature
}

type ObjectSecondaryPort interface {
	ListAllObjects() (Object, error)
	ListObjects(fodler string) (Folder, error)
	GetObjectTemporaryURL(id string) (string, error)
}
