package output

import (
	"github.com/msh5/boy/app/presenter"
)

type ShowOutput struct {
	viewModel *ShowViewModel
}

type ShowViewModel struct {
	Text string
}

func NewShowOutput(viewModel *ShowViewModel) presenter.ShowPresenter {
	return &ShowOutput{viewModel: viewModel}
}

func (o *ShowOutput) Present(result presenter.ShowResult) {
	o.viewModel.Text = result.Text
}
