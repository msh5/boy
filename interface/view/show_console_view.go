package view

import (
	"fmt"

	"github.com/msh5/boy/ifadapter/output"
)

type ShowConsoleView struct {
	viewModel *output.ShowViewModel
}

func NewShowConsoleView(viewModel *output.ShowViewModel) *ShowConsoleView {
	return &ShowConsoleView{viewModel: viewModel}
}

func (v *ShowConsoleView) Update() {
	fmt.Println(v.viewModel.Text)
}
