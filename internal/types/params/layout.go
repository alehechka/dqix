package params

type Layout struct {
	PageTitle  string
	Page       string
	IsDarkMode bool
	CSSVersion string
}

type Index struct {
	LayoutParams Layout
}
