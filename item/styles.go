package item

import "github.com/charmbracelet/lipgloss"

const (
	ellipsis = "…"
)

// DefaultItemStyles defines styling for a default list item.
// See DefaultItemView for when these come into play.
type DefaultItemStyles struct {
	// The Normal state.
	NormalTitle lipgloss.Style
	NormalDesc  lipgloss.Style

	// The selected item state.
	SelectedTitle lipgloss.Style
	SelectedDesc  lipgloss.Style

	// The dimmed state, for when the filter input is initially activated.
	// DimmedTitle lipgloss.Style
	// DimmedDesc  lipgloss.Style

	// // Characters matching the current filter, if any.
	// FilterMatch lipgloss.Style
}

// NewDefaultItemStyles returns style definitions for a default item. See
// DefaultItemView for when these come into play.
func NewDefaultItemStyles() (s DefaultItemStyles) {
	s.NormalTitle = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#dddddd"}).
		Padding(0, 0, 0, 2) //nolint:mnd

	s.NormalDesc = s.NormalTitle.
		Foreground(lipgloss.AdaptiveColor{Light: "#A49FA5", Dark: "#777777"})

	s.SelectedTitle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#990000"}).
		Foreground(lipgloss.AdaptiveColor{Light: "#EE6FF8", Dark: "#990000"}).
		Padding(0, 0, 0, 1)

	s.SelectedDesc = s.SelectedTitle.
		Foreground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#444444"})

	// s.DimmedTitle = lipgloss.NewStyle().
	// 	Foreground(lipgloss.AdaptiveColor{Light: "#A49FA5", Dark: "#777777"}).
	// 	Padding(0, 0, 0, 2) //nolint:mnd

	// s.DimmedDesc = s.DimmedTitle.
	// 	Foreground(lipgloss.AdaptiveColor{Light: "#C2B8C2", Dark: "#4D4D4D"})

	// s.FilterMatch = lipgloss.NewStyle().Underline(true)

	return s
}
