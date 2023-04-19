package signals

func CreateMemento[E any](fn func() (E, error), inital E) *Signal[E] {
	signal := CreateSignal(inital)
	CreateEffect(func() error {
		val, err := fn()
		if err == nil {
			signal.Write(val)
		}

		return err
	})

	return signal
}
