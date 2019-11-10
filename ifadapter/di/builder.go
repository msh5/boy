package di

import (
	"github.com/sarulabs/di"

	"github.com/msh5/boy/app/usecase"
)

type Builder struct {
	*di.Builder
}

func NewBuilder() Builder {
	builder, err := di.NewBuilder()
	if err != nil {
		panic(err)
	}

	return Builder{builder}
}

func (b *Builder) BuildContainer() Container {
	return Container{b.Build()}
}

func (b *Builder) RegisterSnippetExecUsecase(usecase usecase.SnippetExecUsecase) {
	if err := b.Add(di.Def{
		Name: snippetExecUsecaseDIObject,
		Build: func(ctn di.Container) (interface{}, error) {
			return usecase, nil
		},
	}); err != nil {
		panic(err)
	}
}

func (b *Builder) RegisterSnippetShowUsecase(usecase usecase.SnippetShowUsecase) {
	if err := b.Add(di.Def{
		Name: snippetShowUsecaseDIObject,
		Build: func(ctn di.Container) (interface{}, error) {
			return usecase, nil
		},
	}); err != nil {
		panic(err)
	}
}
