package lib

type Agent struct {
}

func NewAgent() *Agent {
	return &Agent{}
}

var k = 0

func (a *Agent) GetAction(gameState GameState) Action {
	k = k + 1
	return actions[k%len(actions)]
}
