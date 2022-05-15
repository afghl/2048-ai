package lib

type Evaluator interface {
	Evaluate(state GameState) int
}

func BaseEvaluator() Evaluator {
	return &base{}
}

type base struct {
}

func (b *base) Evaluate(state GameState) int {
	return 1
}
