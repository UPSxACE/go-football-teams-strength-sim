package tournaments

type Tournament interface {
	HasStarted() bool
	GetWinner() string
	Init()
	NextPhase()
	Render()
}