package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/jesseward/impulse/internal/player"
	"github.com/jesseward/impulse/pkg/module"
)

const (
	trackerHeight      = 24
	playheadDisplayRow = trackerHeight/2 - 1
	channelWidth       = 14 // Approximate width for one channel column
	rowNumWidth        = 3  // Width for the row number column
)

type trackerModel struct {
	module  module.Module
	width   int
	height  int
	row     int
	pattern int
}

func newTrackerModel(m module.Module) trackerModel {
	return trackerModel{
		module: m,
	}
}

func (m *trackerModel) update(state player.PlayerStateUpdate) {
	m.row = state.Row
	m.pattern = state.Pattern
}

func (m trackerModel) View() string {
	if m.module == nil || m.pattern >= m.module.NumPatterns() {
		return ""
	}

	var b strings.Builder

	// Calculate how many channels can be displayed
	availableWidth := m.width - 2 - 2 // subtract border and padding
	maxVisibleChannels := max((availableWidth-rowNumWidth)/channelWidth, 0)
	numChannelsToDisplay := min(m.module.NumChannels(), maxVisibleChannels)

	title := TitleStyle.Render(m.module.Name())
	b.WriteString(title + "\n")

	// Header
	headerStyle := lipgloss.NewStyle().Foreground(CurrentTheme.Secondary).Bold(true)
	header := " "
	for ch := range numChannelsToDisplay {
		header += fmt.Sprintf("    Chan %-4d", ch+1)
	}
	if m.module.NumChannels() > numChannelsToDisplay {
		header += "..."
	}
	b.WriteString(headerStyle.Render(header) + "\n")

	// availableHeight is the total height of the component, minus border, padding and header row
	availableHeight := m.height - 4 - 1
	for displayRow := 1; displayRow <= availableHeight; displayRow++ {
		patternRow := m.row + (displayRow - playheadDisplayRow)

		rowStyle := lipgloss.NewStyle()
		// Default row number style (dimmed)
		rowNumFg := CurrentTheme.Fade

		if displayRow == playheadDisplayRow {
			rowStyle = rowStyle.Background(CurrentTheme.ActiveRowBg).Foreground(CurrentTheme.ActiveRowFg)
			// On active row, use the active foreground color (Black) for high contrast
			rowNumFg = CurrentTheme.ActiveRowFg
		}

		if patternRow >= 0 && patternRow < m.module.NumRows(patternRow) {
			rowNumStr := fmt.Sprintf("%02d", patternRow+1)
			// Apply the calculated foreground to the row number
			b.WriteString(rowStyle.Copy().Foreground(rowNumFg).Render(rowNumStr))

			for ch := range numChannelsToDisplay {
				cellData := m.module.PatternCell(m.pattern, patternRow, ch)
				noteStr := NoteStyle.Copy().Inherit(rowStyle).Render(cellData.HumanNote)
				instrumentStr := InstrumentStyle.Copy().Inherit(rowStyle).Render(fmt.Sprintf("%02X", cellData.Instrument))
				effectStr := EffectStyle.Copy().Inherit(rowStyle).Render(fmt.Sprintf("%X%02X", cellData.Effect, cellData.EffectParam))
				cellStr := fmt.Sprintf(" %s %s %s │", noteStr, instrumentStr, effectStr)
				b.WriteString(rowStyle.Render(cellStr))
			}
			b.WriteString("\n")
		} else {
			b.WriteString("\n")
		}
	}

	style := lipgloss.NewStyle().
		Border(CurrentTheme.AppBorder, true).
		Inherit(BorderColorStyle).
		Width(m.width - 2).
		Height(m.height - 2)
	return style.Render(b.String())
}
