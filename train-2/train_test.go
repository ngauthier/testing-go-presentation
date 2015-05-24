package traintwo

import "testing"

type Speeder interface {
	Speed() int
}

type Engine interface {
	Speeder
	Accel()
	Decel()
}

type engine struct {
	speed int
}

func NewEngine() Engine {
	return &engine{}
}

func (e *engine) Speed() int {
	return e.speed
}

func (e *engine) Accel() {
	e.speed += 10
}

func (e *engine) Decel() {
	e.speed -= 10
}

func expectSpeed(t *testing.T, speeder Speeder, speed int) {
	actual := speeder.Speed()
	if actual != speed {
		t.Fatal("expected", speed, "got", actual)
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

type FakeEngine struct {
	AccelCalls int
	DecelCalls int
}

func NewFakeEngine() *FakeEngine {
	return &FakeEngine{}
}

func (e *FakeEngine) Accel() {
	e.AccelCalls += 1
}

func (e *FakeEngine) Decel() {
	e.DecelCalls += 1
}

func (e *FakeEngine) Speed() int {
	return 0
}

type Train struct {
	engine Engine
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

func trainWithFakeEngine() (*Train, *FakeEngine) {
	t := NewTrain()
	e := NewFakeEngine()
	t.engine = e
	return t, e
}

func TestTrainGo(t *testing.T) {
	tr, e := trainWithFakeEngine()
	tr.Go()
	if e.AccelCalls != 2 {
		t.Error("expected", 2, "got", e.AccelCalls)
	}
}

func TestTrainStop(t *testing.T) {
	tr, e := trainWithFakeEngine()
	tr.Stop()
	if e.DecelCalls != 2 {
		t.Error("expected", 2, "got", e.DecelCalls)
	}
}
