package presenter

type ShowResult struct {
	Text string
}

type ShowPresenter interface {
	Present(ShowResult)
}
