package tournaments

type Tournament interface {
	HasStarted() bool
	IsOver() bool
	GetWinner() string
	Init()
	NextPhase()
	Render()
}