package tui

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ScaleDialog is a dialog for scaling deployments
type ScaleDialog struct {
	deployName string
	current    int32
	target     int32
	input      textinput.Model
	visible    bool
	width      int
	err        string
}

// NewScaleDialog creates a new scale dialog
func NewScaleDialog() ScaleDialog {
	ti := textinput.New()
	ti.Placeholder = "0"
	ti.CharLimit = 5
	ti.Width = 10
	ti.Prompt = ""

	return ScaleDialog{
		input: ti,
	}
}

// Show displays the scale dialog
func (d *ScaleDialog) Show(deployName string, currentReplicas int32) {
	d.deployName = deployName
	d.current = currentReplicas
	d.target = currentReplicas
	d.visible = true
	d.err = ""
	d.input.SetValue(fmt.Sprintf("%d", currentReplicas))
	d.input.Focus()
	d.input.CursorEnd()
}

// Hide hides the scale dialog
func (d *ScaleDialog) Hide() {
	d.visible = false
	d.deployName = ""
	d.err = ""
	d.input.Blur()
}

// IsVisible returns whether the dialog is visible
func (d *ScaleDialog) IsVisible() bool {
	return d.visible
}

// DeploymentName returns the deployment name
func (d *ScaleDialog) DeploymentName() string {
	return d.deployName
}

// TargetReplicas returns the target replica count
func (d *ScaleDialog) TargetReplicas() int32 {
	return d.target
}

// CurrentReplicas returns the current replica count
func (d *ScaleDialog) CurrentReplicas() int32 {
	return d.current
}

// SetWidth sets the dialog width
func (d *ScaleDialog) SetWidth(width int) {
	d.width = width
}

// Update handles key messages for the dialog
func (d *ScaleDialog) Update(msg tea.Msg) (confirmed bool, cancelled bool, cmd tea.Cmd) {
	if !d.visible {
		return false, false, nil
	}

	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "enter":
			// Validate and confirm
			val, err := strconv.ParseInt(d.input.Value(), 10, 32)
			if err != nil {
				d.err = "Invalid number"
				return false, false, nil
			}
			if val < 0 {
				d.err = "Must be >= 0"
				return false, false, nil
			}
			if val > 1000 {
				d.err = "Max is 1000"
				return false, false, nil
			}
			d.target = int32(val)
			d.err = ""
			return true, false, nil

		case "esc":
			return false, true, nil
		}
	}

	// Pass other messages to the text input
	var inputCmd tea.Cmd
	d.input, inputCmd = d.input.Update(msg)
	return false, false, inputCmd
}

// View renders the scale dialog
func (d *ScaleDialog) View() string {
	if !d.visible {
		return ""
	}

	dialogWidth := 45
	if d.width > 0 && d.width < 55 {
		dialogWidth = d.width - 10
	}

	// Title style
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(colorPrimary).
		MarginBottom(1)

	// Deployment name style
	nameStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Bold(true)

	// Info style
	infoStyle := lipgloss.NewStyle().
		Foreground(colorMuted).
		MarginBottom(1)

	// Input label style
	labelStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF"))

	// Error style
	errorStyle := lipgloss.NewStyle().
		Foreground(colorError).
		MarginTop(1)

	// Hint style
	hintStyle := lipgloss.NewStyle().
		Foreground(colorMuted).
		MarginTop(1)

	// Build content
	title := titleStyle.Render("Scale Deployment")
	name := nameStyle.Render(truncateString(d.deployName, dialogWidth-4))
	info := infoStyle.Render(fmt.Sprintf("Current replicas: %d", d.current))

	inputLine := lipgloss.JoinHorizontal(
		lipgloss.Center,
		labelStyle.Render("New replicas: "),
		d.input.View(),
	)

	var content string
	if d.err != "" {
		errMsg := errorStyle.Render(d.err)
		content = lipgloss.JoinVertical(
			lipgloss.Center,
			title,
			name,
			info,
			inputLine,
			errMsg,
		)
	} else {
		hint := hintStyle.Render("Enter to confirm, Esc to cancel")
		content = lipgloss.JoinVertical(
			lipgloss.Center,
			title,
			name,
			info,
			inputLine,
			hint,
		)
	}

	// Dialog box style
	dialogStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colorPrimary).
		Padding(1, 2).
		Width(dialogWidth).
		Align(lipgloss.Center)

	return dialogStyle.Render(content)
}
