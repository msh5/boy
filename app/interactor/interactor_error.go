package interactor

type noFileError struct{}

func (e *noFileError) Error() string {
	return "no files in gist entry"
}
