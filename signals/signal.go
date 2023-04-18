package signals

import (
	"fmt"

	"github.com/eratio08/signals-go/set"
)

var context []*Effect

type Signal struct {
	subscriptions set.Set[*Effect]
	value         any
}

func (s *Signal) subscribe(running *Effect) {
	s.subscriptions.Add(running)
	running.dependencies.Add(&s.subscriptions)
}

func (s *Signal) String() string {
	return fmt.Sprintf("Signal{subscriptions: %v, value: %q}", s.subscriptions, s.value)
}

func (s *Signal) Read() any {
	if len(context) > 0 {
		running := context[len(context)-1]
		s.subscribe(running)
	}

	return s.value
}

func (s *Signal) Write(nextValue any) {
	s.value = nextValue

	for _, sub := range s.subscriptions.ToSlice() {
		sub.execute()
	}
}

func CreateSignal(value any) *Signal {
	return &Signal{
		subscriptions: set.New[*Effect](),
		value:         value,
	}
}
