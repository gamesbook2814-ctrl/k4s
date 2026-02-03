package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// HelpScreen displays keyboard shortcuts
type HelpScreen struct {
	visible bool
	width   int
	height  int
}

// NewHelpScreen creates a new help screen
func NewHelpScreen() HelpScreen {
	return HelpScreen{}
}

// Show displays the help screen
func (h *HelpScreen) Show() {
	h.visible = true
}

// Hide hides the help screen
func (h *HelpScreen) Hide() {
	h.visible = false
}

// Toggle toggles the help screen visibility
func (h *HelpScreen) Toggle() {
	h.visible = !h.visible
}

// IsVisible returns true if the help screen is visible
func (h *HelpScreen) IsVisible() bool {
	return h.visible
}

// SetSize sets the help screen dimensions
func (h *HelpScreen) SetSize(width, height int) {
	h.width = width
	h.height = height
}

// View renders the help screen
func (h *HelpScreen) View() string {
	if !h.visible {
		return ""
	}

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(colorPrimary).
		MarginBottom(1)

	sectionStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(colorSecondary).
		MarginTop(1)

	keyStyle := lipgloss.NewStyle().
		Foreground(colorWarning).
		Width(12)

	descStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF"))

	mutedStyle := lipgloss.NewStyle().
		Foreground(colorMuted).
		Italic(true)

	var content strings.Builder

	content.WriteString(titleStyle.Render("k4s Keyboard Shortcuts"))
	content.WriteString("\n")

	// Global shortcuts
	content.WriteString(sectionStyle.Render("Global"))
	content.WriteString("\n")
	content.WriteString(renderShortcut(keyStyle, descStyle, "?", "Toggle this help screen"))
	content.WriteString(renderShortcut(keyStyle, descStyle, "q", "Quit application"))
	content.WriteString(renderShortcut(keyStyle, descStyle, "Ctrl+C", "Force quit"))
	content.WriteString(renderShortcut(keyStyle, descStyle, "Esc", "Go back / Cancel"))
	content.WriteString(renderShortcut(keyStyle, descStyle, "r", "Refresh current view"))

	// Navigation
	content.WriteString(sectionStyle.Render("Navigation"))
	content.WriteString("\n")
	content.WriteString(renderShortcut(keyStyle, descStyle, "↑/k", "Move up"))
	content.WriteString(renderShortcut(keyStyle, descStyle, "↓/j", "Move down"))
	content.WriteString(renderShortcut(keyStyle, descStyle, "Enter", "Select / Open"))
	content.WriteString(renderShortcut(keyStyle, descStyle, "/", "Filter list"))
	content.WriteString(renderShortcut(keyStyle, descStyle, "0", "Go to Namespaces"))
	content.WriteString(renderShortcut(keyStyle, descStyle, "1", "Go to Pods"))
	content.WriteString(renderShortcut(keyStyle, descStyle, "9", "Go to SSH Hosts"))

	// Pod actions
	content.WriteString(sectionStyle.Render("Pod Actions"))
	content.WriteString("\n")
	content.WriteString(renderShortcut(keyStyle, descStyle, "l", "View logs"))
	content.WriteString(renderShortcut(keyStyle, descStyle, "d", "Delete pod"))
	content.WriteString(renderShortcut(keyStyle, descStyle, "R", "Restart pod (Shift+R)"))

	// Log viewer
	content.WriteString(sectionStyle.Render("Log Viewer"))
	content.WriteString("\n")
	content.WriteString(renderShortcut(keyStyle, descStyle, "f", "Toggle follow mode"))
	content.WriteString(renderShortcut(keyStyle, descStyle, "t", "Toggle timestamps"))
	content.WriteString(renderShortcut(keyStyle, descStyle, "/", "Search in logs"))
	content.WriteString(renderShortcut(keyStyle, descStyle, "n", "Next search match"))
	content.WriteString(renderShortcut(keyStyle, descStyle, "N", "Previous search match"))
	content.WriteString(renderShortcut(keyStyle, descStyle, "g", "Go to top"))
	content.WriteString(renderShortcut(keyStyle, descStyle, "G", "Go to bottom"))
	content.WriteString(renderShortcut(keyStyle, descStyle, "c", "Change container"))

	// Scrolling
	content.WriteString(sectionStyle.Render("Scrolling"))
	content.WriteString("\n")
	content.WriteString(renderShortcut(keyStyle, descStyle, "↑/↓", "Scroll line by line"))
	content.WriteString(renderShortcut(keyStyle, descStyle, "PgUp/PgDn", "Scroll page"))
	content.WriteString(renderShortcut(keyStyle, descStyle, "Home/g", "Go to top"))
	content.WriteString(renderShortcut(keyStyle, descStyle, "End/G", "Go to bottom"))

	content.WriteString("\n")
	content.WriteString(mutedStyle.Render("Press ? or Esc to close"))

	// Create the box
	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colorPrimary).
		Padding(1, 2).
		Width(50)

	return boxStyle.Render(content.String())
}

func renderShortcut(keyStyle, descStyle lipgloss.Style, key, desc string) string {
	return keyStyle.Render(key) + descStyle.Render(desc) + "\n"
}
