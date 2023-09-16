package builder

// construct is a function type that is used to apply a property to the keyboard
type construct func(keyboard *Keyboard)

/*
FunctionalKeyboardBuilder is a struct
that tracks all the actions that are needed to construct the keyboard
*/
type FunctionalKeyboardBuilder struct {
	actions []construct
}

// NewFunctionalKeyboardBuilder is a constructor function to initialize FunctionalKeyboardBuilder
func NewFunctionalKeyboardBuilder() *FunctionalKeyboardBuilder {
	return &FunctionalKeyboardBuilder{}
}

// AddPlate is a function to append an action to add a plate to the keyboard
func (fkb *FunctionalKeyboardBuilder) AddPlate(plate Plate) *FunctionalKeyboardBuilder {
	fkb.actions = append(fkb.actions, func(kb *Keyboard) {
		kb.Plate = plate
	})
	return fkb
}

// AddKeyCaps is a function to append an action to add keycaps to the keyboard
func (fkb *FunctionalKeyboardBuilder) AddKeyCaps(keycaps KeyCap) *FunctionalKeyboardBuilder {
	fkb.actions = append(fkb.actions, func(kb *Keyboard) {
		kb.KeyCaps = keycaps
	})
	return fkb
}

// AddSwitches is a function to append an action to add switches to the keyboard
func (fkb *FunctionalKeyboardBuilder) AddSwitches(switches Switch) *FunctionalKeyboardBuilder {
	fkb.actions = append(fkb.actions, func(kb *Keyboard) {
		kb.Switches = switches
	})
	return fkb
}

// AddPCB is a function to append an action to add a PCB to the keyboard
func (fkb *FunctionalKeyboardBuilder) AddPCB(pcb PCB) *FunctionalKeyboardBuilder {
	fkb.actions = append(fkb.actions, func(kb *Keyboard) {
		kb.PCB = pcb
	})
	return fkb
}

// AddStabilizer is a function to append an action to add stabilizers to the keyboard
func (fkb *FunctionalKeyboardBuilder) AddStabilizer(stabilizers Stabilizer) *FunctionalKeyboardBuilder {
	fkb.actions = append(fkb.actions, func(kb *Keyboard) {
		kb.Stabilizers = stabilizers
	})
	return fkb
}

/*
Build is a function that iterates over all the actions and
constructs the keyboard
*/
func (fkb *FunctionalKeyboardBuilder) Build() Keyboard {
	keyboard := Keyboard{}
	for _, construct := range fkb.actions {
		construct(&keyboard)
	}
	return keyboard
}
