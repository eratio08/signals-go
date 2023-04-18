package signals

import (
	"fmt"
	"github.com/eratio08/signals-go/set"
	"log"
)

type Effect struct {
	dependencies set.Set[*set.Set[*Effect]]
	fn           func() error
}

func (e *Effect) String() string {
	return fmt.Sprintf("Effect{dependencies: %v, fn: %p}", e.dependencies, e.fn)
}

func (e *Effect) cleanup() {
	for _, dep := range e.dependencies.ToSlice() {
		dep.Remove(e)
	}

	e.dependencies = set.New[*set.Set[*Effect]]()
}

func (e *Effect) execute() {
	e.cleanup()
	context = append(context, e)

	err := e.fn()
	if err != nil {
		log.Print(err)
	}

	context = context[:len(context)-1]
}

func CreateEffect(fn func() error) {
	running := Effect{
		dependencies: set.New[*set.Set[*Effect]](),
		fn:           fn,
	}

	running.execute()
}
