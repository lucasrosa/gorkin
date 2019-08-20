package feature

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