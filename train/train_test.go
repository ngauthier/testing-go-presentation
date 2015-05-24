package train

import "testing"

type Engine struct {
	speed int
}

func NewEngine() *Engine {
	return &Engine{}
}

func (e *Engine) Speed() int {
	return e.speed
}

func (e *Engine) Accel() {
	e.speed += 10
}

func (e *Engine) Decel() {
	e.speed -= 10
}

func expectSpeed(t *testing.T, e *Engine, s int) {
	actual := e.Speed()
	if actual != s {
		t.Fatal("expected", s, "got", actual)
	}
}

func TestEngine(t *testing.T) {
	e := NewEngine()
	expectSpeed(t, e, 0)
	e.Accel()
	expectSpeed(t, e, 10)
	e.Decel()
	expectSpeed(t, e, 0)
}

type Train struct {
	engine *Engine
}

func NewTrain() *Train {
	return &Train{
		engine: NewEngine(),
	}
}

func (t *Train) Go() {
	t.engine.Accel()
	t.engine.Accel()
}

func (t *Train) Stop() {
	t.engine.Decel()
	t.engine.Decel()
}

func (t *Train) Speed() int {
	return t.engine.Speed()
}

func expectTrainSpeed(t *testing.T, tr *Train, s int) {
	actual := tr.Speed()
	if actual != s {
		t.Fatal("expected", s, "got", actual)
	}
}

func TestTrain(t *testing.T) {
	tr := NewTrain()
	tr.Go()
	expectTrainSpeed(t, tr, 20)
	tr.Stop()
	expectTrainSpeed(t, tr, 0)
}
