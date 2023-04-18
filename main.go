package main

import (
	"fmt"

	"github.com/eratio08/signals-go/signals"
)

func main() {
	signal := signals.CreateSignal(3)
	signal.Write(5)
	signal.Write(signal.Read().(int) * 2)

	fmt.Println()

	fmt.Println("1. Create Signal")
	signalWithEffect := signals.CreateSignal(0)

	fmt.Println("2. Create Reaction")
	signals.CreateEffect(func() error {
		fmt.Printf("The count is %d\n", signalWithEffect.Read())

		return nil
	})

	fmt.Println("3. Set the count to 5")
	signalWithEffect.Write(5)

	fmt.Println("4. Set the count to 10")
	signalWithEffect.Write(10)

	fmt.Println()

	fmt.Println("1. Create")
	firstName := signals.CreateSignal("John")
	lastName := signals.CreateSignal("Smith")
	showFullName := signals.CreateSignal(true)

	displayName := signals.CreateMemento(func() (any, error) {
		if !showFullName.Read().(bool) {
			return firstName.Read(), nil
		}

		return firstName.Read().(string) + " " + lastName.Read().(string), nil
	})

	signals.CreateEffect(func() error {
		fmt.Println("My name is", displayName.Read())
		return nil
	})

	fmt.Println("2. Set showFullName: false")
	showFullName.Write(false)

	fmt.Println("3. Change lastName")
	lastName.Write("Legend")

	fmt.Println("4. Set showFullName: true")
	showFullName.Write(true)
}
