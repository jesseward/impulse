package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// Theme defines the color palette and styles for the application.
type Theme struct {
	Name         string
	Primary      lipgloss.Color // Main accent (e.g., Cyan)
	Secondary    lipgloss.Color // Secondary accent (e.g., White/Grey)
	Background   lipgloss.Color // Main background (e.g., Black)
	Border       lipgloss.Color // Border color
	HeaderBorder lipgloss.Border
	AppBorder    lipgloss.Border
	Label        lipgloss.Color // Field labels
	Value        lipgloss.Color // Field values
	Separator    lipgloss.Color
	ActiveRowBg  lipgloss.Color // Highlighted row background
	ActiveRowFg  lipgloss.Color // Highlighted row foreground
	Note         lipgloss.Color
	Instrument   lipgloss.Color
	Effect       lipgloss.Color
	Fade         lipgloss.Color // Faint/dimmed text
}

// SunsetFireTheme (Red to Yellow)

var SunsetFireTheme = Theme{
	Name:         "Sunset Fire",
	Primary:      lipgloss.Color("#FF0000"), // Bright Red
	Secondary:    lipgloss.Color("#FFFF00"), // Bright Yellow
	Background:   lipgloss.Color("#000000"), // Black
	Border:       lipgloss.Color("#FF0000"), // Red
	HeaderBorder: lipgloss.RoundedBorder(),
	AppBorder:    lipgloss.RoundedBorder(),
	Label:        lipgloss.Color("#FFFF00"), // Yellow
	Value:        lipgloss.Color("#FFFFFF"), // Bright White
	Separator:    lipgloss.Color("#800000"), // Dark Red
	ActiveRowBg:  lipgloss.Color("#FF0000"), // Red
	ActiveRowFg:  lipgloss.Color("#000000"), // Black
	Note:         lipgloss.Color("#FF0000"), // Bright Red
	Instrument:   lipgloss.Color("#FFFF00"), // Bright Yellow
	Effect:       lipgloss.Color("#FFFFFF"), // White
	Fade:         lipgloss.Color("#808080"), // Gray
}

// OceanSkyTheme (Blue to Cyan)
var OceanSkyTheme = Theme{
	Name:         "Ocean Sky",
	Primary:      lipgloss.Color("#00FFFF"), // Bright Cyan
	Secondary:    lipgloss.Color("#0000FF"), // Bright Blue
	Background:   lipgloss.Color("#000000"), // Black
	Border:       lipgloss.Color("#00FFFF"), // Cyan
	HeaderBorder: lipgloss.RoundedBorder(),
	AppBorder:    lipgloss.RoundedBorder(),
	Label:        lipgloss.Color("#00FFFF"), // Cyan
	Value:        lipgloss.Color("#FFFFFF"), // Bright White
	Separator:    lipgloss.Color("#000080"), // Dark Blue
	ActiveRowBg:  lipgloss.Color("#00FFFF"), // Cyan
	ActiveRowFg:  lipgloss.Color("#000000"), // Black
	Note:         lipgloss.Color("#00FFFF"), // Bright Cyan
	Instrument:   lipgloss.Color("#5F5FFF"), // Light Blue
	Effect:       lipgloss.Color("#FFFFFF"), // White
	Fade:         lipgloss.Color("#808080"), // Gray
}

// NatureFadeTheme (Green)
var NatureFadeTheme = Theme{
	Name:         "Nature Fade",
	Primary:      lipgloss.Color("#00FF00"), // Bright Green
	Secondary:    lipgloss.Color("#00FFFF"), // Cyan
	Background:   lipgloss.Color("#000000"), // Black
	Border:       lipgloss.Color("#00FF00"), // Green
	HeaderBorder: lipgloss.RoundedBorder(),
	AppBorder:    lipgloss.RoundedBorder(),
	Label:        lipgloss.Color("#00FF00"), // Bright Green
	Value:        lipgloss.Color("#FFFFFF"), // Bright White
	Separator:    lipgloss.Color("#008000"), // Dark Green
	ActiveRowBg:  lipgloss.Color("#008000"), // Green
	ActiveRowFg:  lipgloss.Color("#FFFFFF"), // Bright White
	Note:         lipgloss.Color("#00FF00"), // Bright Green
	Instrument:   lipgloss.Color("#00FFFF"), // Cyan
	Effect:       lipgloss.Color("#FFFF00"), // Yellow
	Fade:         lipgloss.Color("#808080"), // Gray
}

// GrayscaleFadeTheme
var GrayscaleFadeTheme = Theme{
	Name:         "Grayscale",
	Primary:      lipgloss.Color("#FFFFFF"), // Bright White
	Secondary:    lipgloss.Color("#808080"), // Gray
	Background:   lipgloss.Color("#000000"), // Black
	Border:       lipgloss.Color("#808080"), // Gray
	HeaderBorder: lipgloss.RoundedBorder(),
	AppBorder:    lipgloss.RoundedBorder(),
	Label:        lipgloss.Color("#C0C0C0"), // Silver
	Value:        lipgloss.Color("#FFFFFF"), // Bright White
	Separator:    lipgloss.Color("#404040"), // Dark Gray
	ActiveRowBg:  lipgloss.Color("#C0C0C0"), // Silver
	ActiveRowFg:  lipgloss.Color("#000000"), // Black
	Note:         lipgloss.Color("#FFFFFF"), // Bright White
	Instrument:   lipgloss.Color("#808080"), // Gray
	Effect:       lipgloss.Color("#606060"), // Dim Gray
	Fade:         lipgloss.Color("#404040"), // Dark Gray
}

// CurrentTheme holds the active theme configuration.
// This makes it easy to swap themes at runtime or compile time.
var CurrentTheme = OceanSkyTheme

// SetTheme sets the CurrentTheme based on the provided name.
func SetTheme(name string) error {
	switch name {
	case "sunset":
		CurrentTheme = SunsetFireTheme
	case "ocean":
		CurrentTheme = OceanSkyTheme
	case "nature":
		CurrentTheme = NatureFadeTheme
	case "gray", "grayscale":
		CurrentTheme = GrayscaleFadeTheme
	default:
		return fmt.Errorf("unknown theme: %s", name)
	}
	updateStyles()
	return nil
}

// ThemeNames returns a list of available theme names.
func ThemeNames() []string {
	return []string{"ocean", "sunset", "nature", "gray"}
}
