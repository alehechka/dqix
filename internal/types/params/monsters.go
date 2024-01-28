package params

import (
	"dqix/internal/types"

	"github.com/a-h/templ"
)

type MonsterFamily struct {
	Family          string
	FamilyTitle     string
	Monsters        types.MonsterSlice
	DisplayMode     string
	LayoutParams    Layout
	SortPathGetter  func(sortField string) templ.SafeURL
	SortOrderGetter func(sortField string) string
}

type Monster struct {
	Monster      types.Monster
	Getter       types.IGetThingFromID
	LayoutParams Layout
}
