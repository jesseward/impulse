package ui

import "github.com/charmbracelet/lipgloss"

var (
	// Global styles
	TitleStyle       lipgloss.Style
	BorderColorStyle lipgloss.Style

	// Header styles
	HeaderStyle    lipgloss.Style
	LabelStyle     lipgloss.Style
	ValueStyle     lipgloss.Style
	SeparatorStyle lipgloss.Style

	// Tracker styles
	NoteStyle       lipgloss.Style
	InstrumentStyle lipgloss.Style
	EffectStyle     lipgloss.Style
)

func init() {
	updateStyles()
}

func updateStyles() {
	// Global
	TitleStyle = lipgloss.NewStyle().
		Background(CurrentTheme.ActiveRowBg).
		Foreground(CurrentTheme.ActiveRowFg).
		Bold(true).
		Padding(0, 1)

	BorderColorStyle = lipgloss.NewStyle().
		BorderForeground(CurrentTheme.Border).
		BorderStyle(CurrentTheme.AppBorder)

	// Header
	HeaderStyle = lipgloss.NewStyle().
		Border(CurrentTheme.HeaderBorder, true).
		Inherit(BorderColorStyle).
		Padding(0, 1)

	LabelStyle = lipgloss.NewStyle().Bold(true).Foreground(CurrentTheme.Label)
	ValueStyle = lipgloss.NewStyle().Foreground(CurrentTheme.Value)
	SeparatorStyle = lipgloss.NewStyle().Foreground(CurrentTheme.Separator)

	// Tracker
	NoteStyle = lipgloss.NewStyle().Foreground(CurrentTheme.Note)
	InstrumentStyle = lipgloss.NewStyle().Foreground(CurrentTheme.Instrument)
	EffectStyle = lipgloss.NewStyle().Foreground(CurrentTheme.Effect)
}
