package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// NotificationType represents the type of notification
type NotificationType int

const (
	NotificationSuccess NotificationType = iota
	NotificationError
	NotificationInfo
)

const notificationDuration = 3 * time.Second

// notificationExpiredMsg is sent when a notification should be hidden
type notificationExpiredMsg struct{}

// Notification represents a notification message
type Notification struct {
	message  string
	notifType NotificationType
	visible  bool
	width    int
}

// NewNotification creates a new notification
func NewNotification() Notification {
	return Notification{}
}

// Show displays a notification
func (n *Notification) Show(message string, notifType NotificationType) tea.Cmd {
	n.message = message
	n.notifType = notifType
	n.visible = true

	// Return a command that will hide the notification after duration
	return tea.Tick(notificationDuration, func(t time.Time) tea.Msg {
		return notificationExpiredMsg{}
	})
}

// Hide hides the notification
func (n *Notification) Hide() {
	n.visible = false
	n.message = ""
}

// IsVisible returns whether the notification is visible
func (n *Notification) IsVisible() bool {
	return n.visible
}

// SetWidth sets the notification width
func (n *Notification) SetWidth(width int) {
	n.width = width
}

// View renders the notification
func (n *Notification) View() string {
	if !n.visible || n.message == "" {
		return ""
	}

	var bgColor lipgloss.Color
	var icon string

	switch n.notifType {
	case NotificationSuccess:
		bgColor = colorSuccess
		icon = "✓"
	case NotificationError:
		bgColor = colorError
		icon = "✗"
	case NotificationInfo:
		bgColor = colorPrimary
		icon = "ℹ"
	}

	style := lipgloss.NewStyle().
		Background(bgColor).
		Foreground(lipgloss.Color("#FFFFFF")).
		Padding(0, 1).
		Bold(true)

	return style.Render(icon + " " + n.message)
}
