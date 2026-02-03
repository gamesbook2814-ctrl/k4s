package tui

import (
	"fmt"
	"sort"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/LywwKkA-aD/k4s/internal/domain"
)

// PodDetailsModel is the model for pod details view
type PodDetailsModel struct {
	pod      *domain.Pod
	events   []domain.PodEvent
	viewport viewport.Model
	styles   Styles
	width    int
	height   int
	ready    bool
}

// NewPodDetailsModel creates a new pod details model
func NewPodDetailsModel(styles Styles) PodDetailsModel {
	return PodDetailsModel{
		styles: styles,
	}
}

// SetPod sets the pod to display
func (m *PodDetailsModel) SetPod(pod *domain.Pod, events []domain.PodEvent) {
	m.pod = pod
	m.events = events
	if m.ready {
		m.viewport.SetContent(m.renderContent())
		m.viewport.GotoTop()
	}
}

// SetSize sets the viewport size
func (m *PodDetailsModel) SetSize(width, height int) {
	m.width = width
	m.height = height
	m.viewport = viewport.New(width, height)
	m.viewport.Style = lipgloss.NewStyle()
	m.ready = true
	if m.pod != nil {
		m.viewport.SetContent(m.renderContent())
	}
}

// Update handles messages
func (m PodDetailsModel) Update(msg tea.Msg) (PodDetailsModel, tea.Cmd) {
	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

// View renders the pod details
func (m PodDetailsModel) View() string {
	if !m.ready || m.pod == nil {
		return "Loading..."
	}
	return m.viewport.View()
}

func (m *PodDetailsModel) renderContent() string {
	if m.pod == nil {
		return "No pod selected"
	}

	var sb strings.Builder
	pod := m.pod

	// Styles
	sectionStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(colorPrimary).
		MarginTop(1)

	labelStyle := lipgloss.NewStyle().
		Foreground(colorMuted).
		Width(16)

	valueStyle := lipgloss.NewStyle()

	// Status color
	var statusStyle lipgloss.Style
	switch pod.Status {
	case domain.PodStatusRunning:
		statusStyle = lipgloss.NewStyle().Foreground(colorSuccess).Bold(true)
	case domain.PodStatusPending:
		statusStyle = lipgloss.NewStyle().Foreground(colorWarning).Bold(true)
	case domain.PodStatusFailed:
		statusStyle = lipgloss.NewStyle().Foreground(colorError).Bold(true)
	default:
		statusStyle = lipgloss.NewStyle().Foreground(colorMuted)
	}

	// === Metadata Section ===
	sb.WriteString(sectionStyle.Render("METADATA"))
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("%s %s\n", labelStyle.Render("Name:"), valueStyle.Render(pod.Name)))
	sb.WriteString(fmt.Sprintf("%s %s\n", labelStyle.Render("Namespace:"), valueStyle.Render(pod.Namespace)))
	sb.WriteString(fmt.Sprintf("%s %s\n", labelStyle.Render("Status:"), statusStyle.Render(pod.Status)))
	sb.WriteString(fmt.Sprintf("%s %s\n", labelStyle.Render("Ready:"), valueStyle.Render(pod.Ready)))
	sb.WriteString(fmt.Sprintf("%s %d\n", labelStyle.Render("Restarts:"), pod.Restarts))
	sb.WriteString(fmt.Sprintf("%s %s\n", labelStyle.Render("Age:"), valueStyle.Render(pod.Age)))
	sb.WriteString(fmt.Sprintf("%s %s\n", labelStyle.Render("Node:"), valueStyle.Render(pod.Node)))
	if pod.IP != "" {
		sb.WriteString(fmt.Sprintf("%s %s\n", labelStyle.Render("IP:"), valueStyle.Render(pod.IP)))
	}

	// === Labels Section ===
	if len(pod.Labels) > 0 {
		sb.WriteString("\n")
		sb.WriteString(sectionStyle.Render("LABELS"))
		sb.WriteString("\n")

		// Sort labels for consistent display
		keys := make([]string, 0, len(pod.Labels))
		for k := range pod.Labels {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			v := pod.Labels[k]
			keyStr := truncateString(k, 30)
			valStr := truncateString(v, 40)
			sb.WriteString(fmt.Sprintf("  %s=%s\n", keyStr, valStr))
		}
	}

	// === Containers Section ===
	sb.WriteString("\n")
	sb.WriteString(sectionStyle.Render(fmt.Sprintf("CONTAINERS (%d)", len(pod.Containers))))
	sb.WriteString("\n")

	for _, c := range pod.Containers {
		// Container header
		containerHeader := lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFFFFF")).
			Render(c.Name)

		var stateStyle lipgloss.Style
		switch c.State {
		case "Running":
			stateStyle = lipgloss.NewStyle().Foreground(colorSuccess)
		case "Waiting":
			stateStyle = lipgloss.NewStyle().Foreground(colorWarning)
		case "Terminated":
			stateStyle = lipgloss.NewStyle().Foreground(colorError)
		default:
			stateStyle = lipgloss.NewStyle().Foreground(colorMuted)
		}

		sb.WriteString(fmt.Sprintf("\n  %s\n", containerHeader))
		sb.WriteString(fmt.Sprintf("    %s %s\n", labelStyle.Render("Image:"), truncateString(c.Image, 50)))
		sb.WriteString(fmt.Sprintf("    %s %s\n", labelStyle.Render("State:"), stateStyle.Render(c.State)))
		if c.StateReason != "" {
			sb.WriteString(fmt.Sprintf("    %s %s\n", labelStyle.Render("Reason:"), c.StateReason))
		}
		if c.Started != "" {
			sb.WriteString(fmt.Sprintf("    %s %s ago\n", labelStyle.Render("Started:"), c.Started))
		}

		readyStr := "false"
		if c.Ready {
			readyStr = "true"
		}
		sb.WriteString(fmt.Sprintf("    %s %s\n", labelStyle.Render("Ready:"), readyStr))
		sb.WriteString(fmt.Sprintf("    %s %d\n", labelStyle.Render("Restarts:"), c.RestartCount))

		// Resources
		if c.Resources.CPURequest != "0" || c.Resources.MemoryRequest != "0" {
			sb.WriteString(fmt.Sprintf("    %s cpu: %s/%s, memory: %s/%s\n",
				labelStyle.Render("Resources:"),
				c.Resources.CPURequest, c.Resources.CPULimit,
				c.Resources.MemoryRequest, c.Resources.MemoryLimit))
		}
	}

	// === Conditions Section ===
	if len(pod.Conditions) > 0 {
		sb.WriteString("\n")
		sb.WriteString(sectionStyle.Render("CONDITIONS"))
		sb.WriteString("\n")

		// Header
		condHeader := lipgloss.NewStyle().
			Foreground(colorMuted).
			Bold(true).
			Render(fmt.Sprintf("  %-20s %-8s %-20s %s", "TYPE", "STATUS", "REASON", "AGE"))
		sb.WriteString(condHeader)
		sb.WriteString("\n")

		for _, cond := range pod.Conditions {
			var statusColor lipgloss.Style
			if cond.Status == "True" {
				statusColor = lipgloss.NewStyle().Foreground(colorSuccess)
			} else {
				statusColor = lipgloss.NewStyle().Foreground(colorError)
			}

			sb.WriteString(fmt.Sprintf("  %-20s %s %-20s %s\n",
				cond.Type,
				statusColor.Render(fmt.Sprintf("%-8s", cond.Status)),
				truncateString(cond.Reason, 20),
				cond.LastTransition))
		}
	}

	// === Events Section ===
	if len(m.events) > 0 {
		sb.WriteString("\n")
		sb.WriteString(sectionStyle.Render(fmt.Sprintf("EVENTS (%d)", len(m.events))))
		sb.WriteString("\n")

		// Show last 10 events
		eventsToShow := m.events
		if len(eventsToShow) > 10 {
			eventsToShow = eventsToShow[len(eventsToShow)-10:]
		}

		for _, e := range eventsToShow {
			var typeStyle lipgloss.Style
			if e.Type == "Warning" {
				typeStyle = lipgloss.NewStyle().Foreground(colorWarning)
			} else {
				typeStyle = lipgloss.NewStyle().Foreground(colorMuted)
			}

			countStr := ""
			if e.Count > 1 {
				countStr = fmt.Sprintf(" (x%d)", e.Count)
			}

			sb.WriteString(fmt.Sprintf("  %s %s%s\n",
				typeStyle.Render(fmt.Sprintf("%-8s", e.Type)),
				e.Reason,
				countStr))
			sb.WriteString(fmt.Sprintf("    %s\n", truncateString(e.Message, m.width-10)))
			sb.WriteString(fmt.Sprintf("    Last seen: %s ago\n", e.LastSeen))
		}
	}

	return sb.String()
}

// ScrollPercent returns the scroll percentage
func (m *PodDetailsModel) ScrollPercent() float64 {
	return m.viewport.ScrollPercent()
}
