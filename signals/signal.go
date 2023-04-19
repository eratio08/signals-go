package signals

import (
	"fmt"

	"github.com/eratio08/signals-go/set"
)

var context []*Effect

type Signal[E any] struct {
	subscriptions set.Set[*Effect]
	value         E
}

func (s *Signal[E]) subscribe(running *Effect) {
	s.subscriptions.Add(running)
	running.dependencies.Add(&s.subscriptions)
}

func (s *Signal[E]) String() string {
	return fmt.Sprintf("Signal{subscriptions: %v, value: %v}", s.subscriptions, s.value)
}

func (s *Signal[E]) Read() E {
	if len(context) > 0 {
		running := context[len(context)-1]
		s.subscribe(running)
	}

	return s.value
}

func (s *Signal[E]) Write(nextValue E) {
	s.value = nextValue

	for _, sub := range s.subscriptions.ToSlice() {
		sub.execute()
	}
}

func CreateSignal[E any](value E) *Signal[E] {
	return &Signal[E]{
		subscriptions: set.New[*Effect](),
		value:         value,
	}
}
