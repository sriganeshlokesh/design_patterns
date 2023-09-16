package builder

import "github.com/davecgh/go-spew/spew"

/*
KeyboardParameterBuilder is a struct that holds the keyboard information
and restricts outside packages to write to it
*/
type KeyboardParameterBuilder struct {
	keyboard Keyboard
}

/*
build is a type of function that builds the keyboard structure
without exposing the fields
*/
type build func(kb *KeyboardParameterBuilder)

// BuildKeyboard is the function that users can use to build their keyboard
func BuildKeyboard(build build) {
	builder := KeyboardParameterBuilder{}
	build(&builder)
	previewKeyboard(&builder.keyboard)
}

// previewKeyboard is a internal function to display the final keyboard build
func previewKeyboard(keyboard *Keyboard) {
	spew.Dump(keyboard)
}

// Plate is a function to assign the plate build for the keyboard
func (kb *KeyboardParameterBuilder) Plate(plate Plate) *KeyboardParameterBuilder {
	kb.keyboard.Plate = plate
	return kb
}

// KeyCap is a function to assign the keycaps for the keyboard
func (kb *KeyboardParameterBuilder) KeyCap(keycaps KeyCap) *KeyboardParameterBuilder {
	kb.keyboard.KeyCaps = keycaps
	return kb
}

// Switch is a function to assign the switches for the keyboard
func (kb *KeyboardParameterBuilder) Switch(switches Switch) *KeyboardParameterBuilder {
	kb.keyboard.Switches = switches
	return kb
}

// PCB is a function to assign the type of PCB board for the keyboard
func (kb *KeyboardParameterBuilder) PCB(pcb PCB) *KeyboardParameterBuilder {
	kb.keyboard.PCB = pcb
	return kb
}

// Stabilizer is a function to assign the stabilizer for the keyboard
func (kb *KeyboardParameterBuilder) Stabilizer(stabilizers Stabilizer) *KeyboardParameterBuilder {
	kb.keyboard.Stabilizers = stabilizers
	return kb
}
