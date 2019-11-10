package usecase

type SnippetExecParameters struct {
	UserID        string
	GistEntryName string
	CommandArgs   []string
}

type SnippetExecUsecase interface {
	Run(params SnippetExecParameters) error
}
