package axolotl

import "github.com/lindsaygelle/animalcrossing/languages"

type Axolotl struct{}

func (a Axolotl) Id() string                            { return "axolotl" }
func (a Axolotl) Special() bool                         { return true }
func (a Axolotl) Translations() []languages.Translation { return []languages.Translation{} }
