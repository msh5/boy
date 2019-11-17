package presenter

type ExecResult struct {
	ExitStatus int
}

type ExecPresenter interface {
	Present(ExecResult)
}
