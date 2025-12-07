package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/jesseward/impulse/internal/ui"
	"github.com/jesseward/impulse/pkg/module"
	"github.com/urfave/cli/v2"
)

func infoAction(c *cli.Context) error {
	if c.NArg() == 0 {
		return cli.Exit(errors.New("no file specified"), 1)
	}
	filePath := c.Args().Get(0)
	// Apply the theme if specified, otherwise default (OceanSky) will be used
	themeName := c.String("theme")
	if themeName != "" {
		if err := ui.SetTheme(themeName); err != nil {
			return cli.Exit(fmt.Sprintf("Invalid theme: %v", err), 1)
		}
	}

	module, err := loadModule(filePath)
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}

	printModuleInfo(module)
	return nil
}

func printModuleInfo(module module.Module) {
	const keyWidth = 15

	// Helper to render a key-value pair
	renderKV := func(key, value string) string {
		k := ui.LabelStyle.Copy().Width(keyWidth).Render(key)
		sep := ui.SeparatorStyle.Render(" │ ")
		v := ui.ValueStyle.Render(value)
		return lipgloss.JoinHorizontal(lipgloss.Left, k, sep, v)
	}

	title := ui.TitleStyle.Render(" Module Information ")

	metadata := []string{
		renderKV("Filename", module.Name()),
		renderKV("Type", module.Type()),
		renderKV("Song Length", strconv.Itoa(module.SongLength())),
		renderKV("BPM", strconv.Itoa(module.DefaultBPM())),
		renderKV("Speed", strconv.Itoa(module.DefaultSpeed())),
		renderKV("Channels", strconv.Itoa(module.NumChannels())),
		renderKV("Patterns", strconv.Itoa(module.NumPatterns())),
	}
	metaBlock := lipgloss.JoinVertical(lipgloss.Left, metadata...)
	metaBox := lipgloss.NewStyle().
		Border(ui.CurrentTheme.AppBorder, true).
		Inherit(ui.BorderColorStyle).
		Padding(0, 1).
		Render(metaBlock)

	samples := module.Samples()
	numSamples := len(samples)
	midpoint := (numSamples + 1) / 2

	var leftColRows, rightColRows []string

	for i, sample := range samples {
		// Format: "01: SampleName"
		// Distinct coloring: Index (LabelStyle) : (SeparatorStyle) Name (ValueStyle)
		indexStr := ui.LabelStyle.Render(fmt.Sprintf("%02d", i+1))
		sepStr := ui.SeparatorStyle.Render(":")
		nameStr := ui.ValueStyle.Render(" " + sample.Name())

		renderedRow := lipgloss.JoinHorizontal(lipgloss.Left, indexStr, sepStr, nameStr)

		if i < midpoint {
			leftColRows = append(leftColRows, renderedRow)
		} else {
			rightColRows = append(rightColRows, renderedRow)
		}
	}

	leftBlock := lipgloss.JoinVertical(lipgloss.Left, leftColRows...)
	rightBlock := lipgloss.JoinVertical(lipgloss.Left, rightColRows...)

	// Join columns with a gap (e.g., 4 spaces)
	sampleBlock := lipgloss.JoinHorizontal(lipgloss.Top,
		leftBlock,
		lipgloss.NewStyle().Width(4).Render(""),
		rightBlock,
	)

	sampleHeader := ui.LabelStyle.Render("Samples:")
	sampleBox := lipgloss.NewStyle().
		Border(ui.CurrentTheme.AppBorder, true).
		Inherit(ui.BorderColorStyle).
		Padding(0, 1).
		Render(lipgloss.JoinVertical(lipgloss.Left, sampleHeader, "", sampleBlock))

	output := lipgloss.JoinVertical(lipgloss.Left,
		title,
		metaBox,
		sampleBox,
	)

	fmt.Println(output)
}
