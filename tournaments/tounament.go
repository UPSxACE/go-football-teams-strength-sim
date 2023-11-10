package tournaments

type Tournament struct {
	CurrentPhase int
	TotalPhases int
	Started bool
	Winner string
	Init func()
	NextPhase func()
	Render func()
}