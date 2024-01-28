package params

import "fmt"

type Layout struct {
	PageTitle  string
	Page       string
	IsDarkMode bool
	CSSVersion string
	IconPath   string
}

// GetPageTitle returns the formatted PageTitle when not empty, otherwise will return the default page title.
func (l Layout) GetPageTitle() string {
	if l.PageTitle == "" {
		return "Dragon Quest IX"
	}

	return fmt.Sprintf("DQIX | %s", l.PageTitle)
}

// GetIconPath returns the IconPath when not empty, otherwise will return the default path to the app's favicon.
func (l Layout) GetIconPath() string {
	if l.IconPath == "" {
		return "/static/favicon.ico"
	}

	return l.IconPath
}

type Index struct {
	LayoutParams Layout
}
