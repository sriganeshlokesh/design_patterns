package builder

// KeyboardBuilder is a builder struct used to create a keyboard
type KeyboardBuilder struct {
	Keyboard Keyboard
}

// NewKeyboardBuilder is a function that initializes the KeyboardBuilder struct
func NewKeyboardBuilder() *KeyboardBuilder {
	return &KeyboardBuilder{
		Keyboard: Keyboard{},
	}
}

// Plate is a type to determine the type of plate the keyboard is built with
type Plate string

const (
	Aluminium     Plate = "Aluminium"
	Polycarbonate Plate = "Polycardbonate"
	Brass         Plate = "Brass"
	FR4           Plate = "FR4"
	POM           Plate = "POM"
	CarbonFiber   Plate = "Carbon Fiber"
)

// KeyCap is a type to determine the type of keycaps we need to use for the keyboard
type KeyCap string

const (
	OEM    KeyCap = "OEM"
	DSA    KeyCap = "DSA"
	SA     KeyCap = "SA"
	Cherry KeyCap = "Cherry"
)

// Switch is a type to determine the type of switch to use on the board
type Switch string

const (
	Tactile Switch = "Tactile"
	Linear  Switch = "Linear"
	Clicky  Switch = "Clicky"
)

// PCB is a type to determine the type of PCB Board built needed by the user
type PCB string

const (
	HotSwap PCB = "Hot Swap"
	Solder  PCB = "Solder"
)

// Stabilizer is a type to determine the different kinds of stabilizers available on the board.
type Stabilizer string

const (
	PlateMounted Stabilizer = "Plate Mounted"
	ScrewIn      Stabilizer = "Screw In"
	SnapIn       Stabilizer = "Snap In"
)

type Keyboard struct {
	Plate       Plate
	KeyCaps     KeyCap
	Switches    Switch
	PCB         PCB
	Stabilizers Stabilizer
}
