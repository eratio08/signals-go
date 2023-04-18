package signals

func CreateMemento(fn func() (any, error)) *Signal {
	signal := CreateSignal(nil)
	CreateEffect(func() error {
		val, err := fn()
		if err == nil {
			signal.Write(val)
		}

		return err
	})

	return signal
}
