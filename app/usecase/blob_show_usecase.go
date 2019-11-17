package usecase

type BlobShowParameters struct {
	RepositoryOwner string
	RepositoryName  string
	BlobPath        string
}

type BlobShowUsecase interface {
	Run(params BlobShowParameters) error
}
