package usecase

type BlobExecParameters struct {
	RepositoryOwner string
	RepositoryName  string
	BlobPath        string
	CommandArgs     []string
}

type BlobExecUsecase interface {
	Run(params BlobExecParameters) error
}
