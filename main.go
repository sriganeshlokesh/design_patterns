package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/sriganeshlokesh/designpatterns/internal/patterns/builder"
)

func main() {
	// Builder - Implementation I - Builder
	kb := builder.KeyboardBuilder{}
	kb.
		SetPlate(builder.POM).
		SetPCB(builder.HotSwap).
		SetSwitches(builder.Tactile).
		SetKeyCaps(builder.Cherry).
		SetStabilizers(builder.PlateMounted)

	spew.Dump(kb.Build())

	// Builder - Implementation II - Builder Facet
	mkb := builder.NewKeyboardFacetBuilder()
	mkb.
		Internal().With(builder.HotSwap, builder.PlateMounted, builder.Aluminium).
		External().Look(builder.DSA, builder.Tactile)

	spew.Dump(mkb.Build())

	// Builder - Implementation III - Builder Parameter
	builder.BuildKeyboard(func(b *builder.KeyboardParameterBuilder) {
		b.
			PCB(builder.HotSwap).
			KeyCap(builder.OEM).
			Switch(builder.Clicky).
			Plate(builder.Brass).
			Stabilizer(builder.ScrewIn)
	})

	// Builder - Implementation IV - Functional Builder
	fkb := builder.NewFunctionalKeyboardBuilder()
	keyboard := fkb.
		AddPCB(builder.HotSwap).
		AddKeyCaps(builder.DSA).
		AddSwitches(builder.Clicky).
		AddPlate(builder.Aluminium).
		AddStabilizer(builder.SnapIn).
		Build()

	spew.Dump(keyboard)
}
