package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ContainerSelector is a component for selecting a container
type ContainerSelector struct {
	containers    []string
	selectedIndex int
	visible       bool
	width         int
}

// NewContainerSelector creates a new container selector
func NewContainerSelector() ContainerSelector {
	return ContainerSelector{}
}

// Show displays the container selector
func (c *ContainerSelector) Show(containers []string, currentContainer string) {
	c.containers = containers
	c.visible = true
	c.selectedIndex = 0

	// Find current container index
	for i, name := range containers {
		if name == currentContainer {
			c.selectedIndex = i
			break
		}
	}
}

// Hide hides the container selector
func (c *ContainerSelector) Hide() {
	c.visible = false
}

// IsVisible returns whether the selector is visible
func (c *ContainerSelector) IsVisible() bool {
	return c.visible
}

// SetWidth sets the selector width
func (c *ContainerSelector) SetWidth(width int) {
	c.width = width
}

// SelectedContainer returns the currently selected container name
func (c *ContainerSelector) SelectedContainer() string {
	if c.selectedIndex < len(c.containers) {
		return c.containers[c.selectedIndex]
	}
	return ""
}

// Update handles key messages for the selector
func (c *ContainerSelector) Update(msg tea.Msg) (selected bool, cancelled bool) {
	if !c.visible {
		return false, false
	}

	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "up", "k":
			if c.selectedIndex > 0 {
				c.selectedIndex--
			}
		case "down", "j":
			if c.selectedIndex < len(c.containers)-1 {
				c.selectedIndex++
			}
		case "enter":
			return true, false
		case "esc", "c":
			return false, true
		}
	}

	return false, false
}

// View renders the container selector
func (c *ContainerSelector) View() string {
	if !c.visible || len(c.containers) == 0 {
		return ""
	}

	selectorWidth := 40
	if c.width > 0 && c.width < 50 {
		selectorWidth = c.width - 10
	}

	// Title style
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(colorPrimary).
		MarginBottom(1)

	// Build container list
	var items string
	for i, container := range c.containers {
		itemStyle := lipgloss.NewStyle().
			Padding(0, 1).
			Width(selectorWidth - 4)

		if i == c.selectedIndex {
			itemStyle = itemStyle.
				Background(colorPrimary).
				Foreground(lipgloss.Color("#FFFFFF")).
				Bold(true)
			items += itemStyle.Render(fmt.Sprintf("▸ %s", container)) + "\n"
		} else {
			itemStyle = itemStyle.
				Foreground(lipgloss.Color("#FFFFFF"))
			items += itemStyle.Render(fmt.Sprintf("  %s", container)) + "\n"
		}
	}

	// Hint text
	hintStyle := lipgloss.NewStyle().
		Foreground(colorMuted).
		MarginTop(1)
	hint := hintStyle.Render("↑/↓: select • Enter: confirm • Esc: cancel")

	// Build selector content
	content := lipgloss.JoinVertical(
		lipgloss.Left,
		titleStyle.Render("Select Container"),
		items,
		hint,
	)

	// Selector box style
	selectorStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colorPrimary).
		Padding(1, 2).
		Width(selectorWidth)

	return selectorStyle.Render(content)
}
