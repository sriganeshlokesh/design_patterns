package builder

// KeyboardBuilderFacet is a struct to hold the final keyboard build
type KeyboardBuilderFacet struct {
	keyboard *Keyboard
}

/*
KeyboardInternalBuilder is a struct that holds information
for internal keyboard properties
*/
type KeyboardInternalBuilder struct {
	KeyboardBuilderFacet
}

/*
KeyboardExternalBuilder is a struct that holds information
for external keyboard properties
*/
type KeyboardExternalBuilder struct {
	KeyboardBuilderFacet
}

// Internal is a function that builds the internal properties for a keyboard
func (kb *KeyboardBuilderFacet) Internal() *KeyboardInternalBuilder {
	return &KeyboardInternalBuilder{*kb}
}

// External is a function that builds the external properties for a keyboard
func (kb *KeyboardBuilderFacet) External() *KeyboardExternalBuilder {
	return &KeyboardExternalBuilder{*kb}
}

// NewKeyboardFacetBuilder is a function that initializes the KeyboardBuildFacet
func NewKeyboardFacetBuilder() *KeyboardBuilderFacet {
	return &KeyboardBuilderFacet{&Keyboard{}}
}

// Look is a function to assign the required external property values
func (ikb *KeyboardExternalBuilder) Look(keycaps KeyCap, switches Switch) *KeyboardExternalBuilder {
	ikb.keyboard.KeyCaps = keycaps
	ikb.keyboard.Switches = switches
	return ikb
}

// With is a function to assign the required internal property values
func (exb *KeyboardInternalBuilder) With(pcb PCB, stabilizers Stabilizer, plate Plate) *KeyboardInternalBuilder {
	exb.keyboard.PCB = pcb
	exb.keyboard.Stabilizers = stabilizers
	exb.keyboard.Plate = plate
	return exb
}

// Build is a function that returns the final build of the keyboard to the user
func (b *KeyboardBuilderFacet) Build() *Keyboard {
	return b.keyboard
}
