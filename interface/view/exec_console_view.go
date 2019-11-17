package view

import (
	"fmt"

	"github.com/msh5/boy/ifadapter/output"
)

type ExecConsoleView struct {
	viewModel *output.ExecViewModel
}

func NewExecConsoleView(viewModel *output.ExecViewModel) *ExecConsoleView {
	return &ExecConsoleView{viewModel: viewModel}
}

func (v *ExecConsoleView) Update() {
	if v.viewModel.ExitStatus == 0 {
		fmt.Printf("finish successfully\n")
	} else {
		fmt.Printf("exit with error status: %d\n", v.viewModel.ExitStatus)
	}
}
