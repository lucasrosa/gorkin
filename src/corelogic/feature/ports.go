package feature

// FeaturePrimaryPort is the entrypoint for the feature Package
type FeaturePrimaryPort interface {
	GetAll() []Feature
}

// FolderPrimaryPort is the entrypoint for the folders
type FolderPrimaryPort interface {
	GetAll() []Folder
}

// DatabaseSecondaryPort is the way the business rules communicate to the external world
type DatabaseSecondaryPort interface {
	GetAll() []Feature
}

type ObjectSecondaryPort interface {
	GetAll() []Folder
}