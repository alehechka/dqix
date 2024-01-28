package params

import (
	"dqix/internal/types"

	"github.com/a-h/templ"
)

type InventoryClassification struct {
	Classification  string
	Inventories     types.InventorySlice
	Stats           types.HasInventoryStats
	DisplayMode     string
	LayoutParams    Layout
	SortPathGetter  func(sortField string) templ.SafeURL
	SortOrderGetter func(sortField string) string
}

type Inventory struct {
	Inventory    types.Inventory
	Getter       types.IGetThingFromID
	LayoutParams Layout
}
