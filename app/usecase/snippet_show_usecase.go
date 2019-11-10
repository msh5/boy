package usecase

type SnippetShowParameters struct {
	UserID        string
	GistEntryName string
}

type SnippetShowUsecase interface {
	Run(params SnippetShowParameters) error
}
