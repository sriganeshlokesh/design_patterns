package builder

func (kb *KeyboardBuilder) SetPlate(plate Plate) *KeyboardBuilder {
	kb.Keyboard.Plate = plate
	return kb
}

func (kb *KeyboardBuilder) SetKeyCaps(caps KeyCap) *KeyboardBuilder {
	kb.Keyboard.KeyCaps = caps
	return kb
}

func (kb *KeyboardBuilder) SetSwitches(switches Switch) *KeyboardBuilder {
	kb.Keyboard.Switches = switches
	return kb
}

func (kb *KeyboardBuilder) SetPCB(pcb PCB) *KeyboardBuilder {
	kb.Keyboard.PCB = pcb
	return kb
}

func (kb *KeyboardBuilder) SetStabilizers(stabilizer Stabilizer) *KeyboardBuilder {
	kb.Keyboard.Stabilizers = stabilizer
	return kb
}

func (kb *KeyboardBuilder) Build() Keyboard {
	return kb.Keyboard
}
