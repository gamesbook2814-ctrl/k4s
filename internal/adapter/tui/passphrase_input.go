package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// PassphraseInput is a modal dialog for entering SSH passphrase
type PassphraseInput struct {
	visible   bool
	input     textinput.Model
	hostName  string
	width     int
}

// NewPassphraseInput creates a new passphrase input dialog
func NewPassphraseInput() PassphraseInput {
	ti := textinput.New()
	ti.Placeholder = "Enter passphrase..."
	ti.EchoMode = textinput.EchoPassword
	ti.EchoCharacter = '•'
	ti.CharLimit = 256
	ti.Width = 40

	return PassphraseInput{
		input: ti,
	}
}

// Show displays the passphrase input for the given host
func (p *PassphraseInput) Show(hostName string) {
	p.visible = true
	p.hostName = hostName
	p.input.Reset()
	p.input.Focus()
}

// Hide hides the passphrase input
func (p *PassphraseInput) Hide() {
	p.visible = false
	p.input.Blur()
	p.input.Reset()
}

// IsVisible returns true if the dialog is visible
func (p *PassphraseInput) IsVisible() bool {
	return p.visible
}

// SetWidth sets the dialog width
func (p *PassphraseInput) SetWidth(width int) {
	p.width = width
}

// Update handles input messages, returns (passphrase, submitted, cancelled, cmd)
func (p *PassphraseInput) Update(msg tea.Msg) (string, bool, bool, tea.Cmd) {
	// Handle key messages for submit/cancel
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "enter":
			passphrase := p.input.Value()
			return passphrase, true, false, nil
		case "esc":
			return "", false, true, nil
		}
	}

	// Pass all messages to the textinput
	var cmd tea.Cmd
	p.input, cmd = p.input.Update(msg)
	return "", false, false, cmd
}

// View renders the passphrase input dialog
func (p *PassphraseInput) View() string {
	if !p.visible {
		return ""
	}

	// Dialog box style
	dialogStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colorPrimary).
		Padding(1, 2).
		Width(50)

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(colorPrimary)

	subtitleStyle := lipgloss.NewStyle().
		Foreground(colorMuted).
		Italic(true)

	helpStyle := lipgloss.NewStyle().
		Foreground(colorMuted)

	// Build content
	var sb strings.Builder
	sb.WriteString(titleStyle.Render("SSH Passphrase Required"))
	sb.WriteString("\n\n")
	sb.WriteString(subtitleStyle.Render("Host: " + p.hostName))
	sb.WriteString("\n\n")
	sb.WriteString(p.input.View())
	sb.WriteString("\n\n")
	sb.WriteString(helpStyle.Render("Enter: submit • Esc: cancel"))

	return dialogStyle.Render(sb.String())
}
