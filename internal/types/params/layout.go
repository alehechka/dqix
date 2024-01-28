package params

type Layout struct {
	PageTitle  string
	Page       string
	IsDarkMode bool
	CSSVersion string
	IconPath   string
}

// GetIconPath returns the IconPath when it not empty, otherwise will return the default path to the app's favicon.
func (l Layout) GetIconPath() string {
	if l.IconPath != "" {
		return l.IconPath
	}

	return "/static/favicon.ico"
}

type Index struct {
	LayoutParams Layout
}
