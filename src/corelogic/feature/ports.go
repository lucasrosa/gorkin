package feature

// FeaturePrimaryPort is the entrypoint for the checkout Package
type FeaturePrimaryPort interface {
	GetAll() []Feature
}

// DatabaseSecondaryPort is the way the business rules communicate to the external world
type DatabaseSecondaryPort interface {
	GetAll() []Feature
}
