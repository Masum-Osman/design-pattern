package memento

import "fmt"

func mementoInit() {
	caretaker := &Caretaker{
		mementoArray: make([]*Memento, 0),
	}

	originator := &Originator{
		state: "Initial",
	}

	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.setState("Stage 2")
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.setState("Stage 3")
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.restoreMemento(caretaker.getMemento(1))
	fmt.Printf("Restored to State: %s\n", originator.getState())

	originator.restoreMemento(caretaker.getMemento(0))
	fmt.Printf("Restored to State: %s\n", originator.getState())

	originator.setState("Stage 4")
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.setState("Stage 5")
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.setState("Stage 6")
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.restoreMemento(caretaker.getMemento(4))
	fmt.Printf("Restored to State: %s\n", originator.getState())
}
