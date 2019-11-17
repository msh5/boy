package output

import (
	"github.com/msh5/boy/app/presenter"
)

type ExecOutput struct {
	viewModel *ExecViewModel
}

type ExecViewModel struct {
	ExitStatus int
}

func NewExecOutput(viewModel *ExecViewModel) presenter.ExecPresenter {
	return &ExecOutput{viewModel: viewModel}
}

func (o *ExecOutput) Present(result presenter.ExecResult) {
	o.viewModel.ExitStatus = result.ExitStatus
}
