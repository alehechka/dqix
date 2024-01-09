package types

import (
	"net/url"
	"strings"

	"github.com/a-h/templ"
)

type SortOrder string

const (
	SortOrderNone SortOrder = ""
	SortOrderAsc  SortOrder = "ascending"
	SortOrderDesc SortOrder = "descending"
)

type Sort struct {
	Field string
	Order SortOrder
}

func (s Sort) NextSortOrder() SortOrder {
	switch s.Order {
	case SortOrderAsc:
		return SortOrderNone
	case SortOrderDesc:
		return SortOrderAsc
	case SortOrderNone:
		fallthrough
	default:
		return SortOrderDesc
	}
}

func (s *Sort) SetNextSortOrder() {
	s.Order = s.NextSortOrder()
}

func (s Sort) String() string {
	switch s.Order {
	case SortOrderDesc:
		return "-" + s.Field
	case SortOrderAsc:
		return s.Field
	case SortOrderNone:
		fallthrough
	default:
		return ""
	}
}

type Sorts []Sort

type SortMap map[string]SortOrder

func (s Sorts) Get(field string) Sort {
	for _, sort := range s {
		if field == sort.Field {
			return sort
		}
	}

	return Sort{Field: field}
}

func (s Sorts) ToMap() SortMap {
	sortMap := make(SortMap)

	for _, sort := range s {
		sortMap[sort.Field] = sort.Order
	}

	return sortMap
}

func (s Sorts) Set(newSort Sort) Sorts {
	newSorts := make(Sorts, 0)
	var isSet bool
	for _, sort := range s {
		if sort.Field == newSort.Field {
			if newSort.Order != SortOrderNone {
				newSorts = append(newSorts, newSort)
				isSet = true
			}
		} else {
			newSorts = append(newSorts, sort)
		}
	}

	if !isSet && newSort.Order != SortOrderNone {
		newSorts = append(newSorts, newSort)
	}

	return newSorts
}

func (s Sorts) GetStringSorts() (strSorts []string) {
	for _, sort := range s {
		strSort := sort.String()
		if strSort != "" {
			strSorts = append(strSorts, strSort)
		}
	}
	return
}

func (s Sorts) String() string {
	return strings.Join(s.GetStringSorts(), ",")
}

func ParseSortingQuery(sortQuery string) Sorts {
	parts := strings.Split(sortQuery, ",")

	sorts := make([]Sort, 0)
	for _, part := range parts {
		order := SortOrderAsc
		trimmed := strings.TrimPrefix(part, "-")
		if len(trimmed) < len(part) {
			order = SortOrderDesc
		}
		sorts = append(sorts, Sort{Field: trimmed, Order: order})
	}

	return sorts
}

func PrepareSortPath(uri url.URL) func(sortField string) templ.SafeURL {
	query := uri.Query()
	sortQuery := query.Get("sort")

	sorts := ParseSortingQuery(sortQuery)

	return func(sortField string) templ.SafeURL {
		sort := sorts.Get(sortField)
		sort.SetNextSortOrder()

		newSorts := sorts.Set(sort)

		query.Set("sort", newSorts.String())

		uri.RawQuery = query.Encode()

		return templ.SafeURL(uri.String())
	}
}

func PrepareSimpleSortPath(uri url.URL) func(sortField string) templ.SafeURL {
	query := uri.Query()
	sortQuery := query.Get("sort")

	sorts := ParseSortingQuery(sortQuery)

	return func(sortField string) templ.SafeURL {
		sort := sorts.Get(sortField)
		sort.SetNextSortOrder()

		newSorts := Sorts{sort}

		query.Set("sort", newSorts.String())

		uri.RawQuery = query.Encode()

		return templ.SafeURL(uri.String())
	}
}

func GetSortOrder(uri *url.URL) func(sortField string) string {
	sortQuery := uri.Query().Get("sort")

	sortMap := ParseSortingQuery(sortQuery).ToMap()

	return func(sortField string) string {
		order, _ := sortMap[sortField]
		return string(order)
	}
}
