package trainthree

import (
	"testing"
	"time"
)

type Speeder interface {
	Speed() int
}

type Engine interface {
	Speeder
	Accel()
	Decel()
	Stop()
	Run()
}

type accelCommand struct {
	done chan bool
}

func newAccelCommand() *accelCommand {
	return &accelCommand{done: make(chan bool)}
}

type engine struct {
	speed    int
	stop     chan bool
	done     chan bool
	accelcmd chan *accelCommand
}

func newEngine() *engine {
	return &engine{
		accelcmd: make(chan *accelCommand),
		stop:     make(chan bool),
		done:     make(chan bool),
	}
}

func NewEngine() Engine {
	return newEngine()
}

func (e *engine) Run() {
	for {
		select {
		case c := <-e.accelcmd:
			e.speed += 10
			time.Sleep(100 * time.Millisecond)
			close(c.done)
		case <-e.stop:
			close(e.done)
			return
		}
	}
}

func (e *engine) Stop() {
	e.stop <- true
	<-e.done
}

func (e *engine) Speed() int {
	return e.speed
}

func (e *engine) Accel() {
	cmd := newAccelCommand()
	e.accelcmd <- cmd
	<-cmd.done
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
	go e.Run()
	expectSpeed(t, e, 0)
	e.Accel()
	expectSpeed(t, e, 10)
	e.Decel()
	expectSpeed(t, e, 0)
	e.Stop()
}

func TestEngineAccel(t *testing.T) {
	e := newEngine()

	done := make(chan bool)
	go func() {
		e.Accel()
		close(done)
	}()

	i, ok := <-e.accelcmd
	if i == nil || !ok {
		t.Fatal("expected acceleration command")
	}

	close(i.done)
	<-done
}
