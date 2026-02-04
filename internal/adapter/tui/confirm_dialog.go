package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ConfirmAction represents the type of action being confirmed
type ConfirmAction int

const (
	ConfirmActionNone ConfirmAction = iota
	ConfirmActionDeletePod
	ConfirmActionRestartPod
	ConfirmActionDeleteDeployment
	ConfirmActionRestartDeployment
)

// ConfirmDialog is a confirmation dialog model
type ConfirmDialog struct {
	action      ConfirmAction
	title       string
	message     string
	targetName  string
	visible     bool
	width       int
	yesSelected bool
}

// NewConfirmDialog creates a new confirmation dialog
func NewConfirmDialog() ConfirmDialog {
	return ConfirmDialog{
		yesSelected: false,
	}
}

// Show displays the confirmation dialog
func (d *ConfirmDialog) Show(action ConfirmAction, targetName string) {
	d.action = action
	d.targetName = targetName
	d.visible = true
	d.yesSelected = false

	switch action {
	case ConfirmActionDeletePod:
		d.title = "Delete Pod"
		d.message = fmt.Sprintf("Are you sure you want to delete pod '%s'?", targetName)
	case ConfirmActionRestartPod:
		d.title = "Restart Pod"
		d.message = fmt.Sprintf("Are you sure you want to restart pod '%s'?\n(This will delete the pod; the controller will recreate it)", targetName)
	case ConfirmActionDeleteDeployment:
		d.title = "Delete Deployment"
		d.message = fmt.Sprintf("Are you sure you want to delete deployment '%s'?\n(All associated pods will be terminated)", targetName)
	case ConfirmActionRestartDeployment:
		d.title = "Restart Deployment"
		d.message = fmt.Sprintf("Are you sure you want to restart deployment '%s'?\n(This triggers a rolling restart of all pods)", targetName)
	default:
		d.title = "Confirm"
		d.message = "Are you sure?"
	}
}

// Hide hides the confirmation dialog
func (d *ConfirmDialog) Hide() {
	d.visible = false
	d.action = ConfirmActionNone
	d.targetName = ""
}

// IsVisible returns whether the dialog is visible
func (d *ConfirmDialog) IsVisible() bool {
	return d.visible
}

// Action returns the current action being confirmed
func (d *ConfirmDialog) Action() ConfirmAction {
	return d.action
}

// TargetName returns the target name for the action
func (d *ConfirmDialog) TargetName() string {
	return d.targetName
}

// SetWidth sets the dialog width
func (d *ConfirmDialog) SetWidth(width int) {
	d.width = width
}

// Update handles key messages for the dialog
func (d *ConfirmDialog) Update(msg tea.Msg) (confirmed bool, cancelled bool) {
	if !d.visible {
		return false, false
	}

	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "y", "Y":
			return true, false
		case "n", "N", "esc":
			return false, true
		case "left", "h":
			d.yesSelected = true
		case "right", "l":
			d.yesSelected = false
		case "enter":
			if d.yesSelected {
				return true, false
			}
			return false, true
		case "tab":
			d.yesSelected = !d.yesSelected
		}
	}

	return false, false
}

// View renders the confirmation dialog
func (d *ConfirmDialog) View() string {
	if !d.visible {
		return ""
	}

	dialogWidth := 50
	if d.width > 0 && d.width < 60 {
		dialogWidth = d.width - 10
	}

	// Title style
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(colorWarning).
		MarginBottom(1)

	// Message style
	messageStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Width(dialogWidth - 4).
		MarginBottom(1)

	// Button styles
	yesButtonStyle := lipgloss.NewStyle().
		Padding(0, 2).
		MarginRight(2)

	noButtonStyle := lipgloss.NewStyle().
		Padding(0, 2)

	if d.yesSelected {
		yesButtonStyle = yesButtonStyle.
			Background(colorError).
			Foreground(lipgloss.Color("#FFFFFF")).
			Bold(true)
		noButtonStyle = noButtonStyle.
			Background(lipgloss.Color("#444444")).
			Foreground(lipgloss.Color("#AAAAAA"))
	} else {
		yesButtonStyle = yesButtonStyle.
			Background(lipgloss.Color("#444444")).
			Foreground(lipgloss.Color("#AAAAAA"))
		noButtonStyle = noButtonStyle.
			Background(colorPrimary).
			Foreground(lipgloss.Color("#FFFFFF")).
			Bold(true)
	}

	// Build buttons
	yesButton := yesButtonStyle.Render("[Y]es")
	noButton := noButtonStyle.Render("[N]o")
	buttons := lipgloss.JoinHorizontal(lipgloss.Center, yesButton, noButton)

	// Hint text
	hintStyle := lipgloss.NewStyle().
		Foreground(colorMuted).
		MarginTop(1)
	hint := hintStyle.Render("Press Y to confirm, N or Esc to cancel")

	// Build dialog content
	content := lipgloss.JoinVertical(
		lipgloss.Center,
		titleStyle.Render(d.title),
		messageStyle.Render(d.message),
		buttons,
		hint,
	)

	// Dialog box style
	dialogStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colorWarning).
		Padding(1, 2).
		Width(dialogWidth).
		Align(lipgloss.Center)

	return dialogStyle.Render(content)
}
