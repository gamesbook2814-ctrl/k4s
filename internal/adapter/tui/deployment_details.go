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

// DeploymentDetailsModel is the model for deployment details view
type DeploymentDetailsModel struct {
	deployment *domain.Deployment
	viewport   viewport.Model
	styles     Styles
	width      int
	height     int
	ready      bool
}

// NewDeploymentDetailsModel creates a new deployment details model
func NewDeploymentDetailsModel(styles Styles) DeploymentDetailsModel {
	return DeploymentDetailsModel{
		styles: styles,
	}
}

// SetDeployment sets the deployment to display
func (m *DeploymentDetailsModel) SetDeployment(dep *domain.Deployment) {
	m.deployment = dep
	if m.ready {
		m.viewport.SetContent(m.renderContent())
		m.viewport.GotoTop()
	}
}

// SetSize sets the viewport size
func (m *DeploymentDetailsModel) SetSize(width, height int) {
	m.width = width
	m.height = height
	m.viewport = viewport.New(width, height)
	m.viewport.Style = lipgloss.NewStyle()
	m.ready = true
	if m.deployment != nil {
		m.viewport.SetContent(m.renderContent())
	}
}

// Update handles messages
func (m DeploymentDetailsModel) Update(msg tea.Msg) (DeploymentDetailsModel, tea.Cmd) {
	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

// View renders the deployment details
func (m DeploymentDetailsModel) View() string {
	if !m.ready || m.deployment == nil {
		return "Loading..."
	}
	return m.viewport.View()
}

func (m *DeploymentDetailsModel) renderContent() string {
	if m.deployment == nil {
		return "No deployment selected"
	}

	var sb strings.Builder
	dep := m.deployment

	// Styles
	sectionStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(colorPrimary).
		MarginTop(1)

	labelStyle := lipgloss.NewStyle().
		Foreground(colorMuted).
		Width(16)

	valueStyle := lipgloss.NewStyle()

	// Status color based on ready state
	var statusStyle lipgloss.Style
	if dep.ReadyReplicas == dep.Replicas && dep.Replicas > 0 {
		statusStyle = lipgloss.NewStyle().Foreground(colorSuccess).Bold(true)
	} else if dep.ReadyReplicas == 0 && dep.Replicas > 0 {
		statusStyle = lipgloss.NewStyle().Foreground(colorError).Bold(true)
	} else {
		statusStyle = lipgloss.NewStyle().Foreground(colorWarning).Bold(true)
	}

	// === Metadata Section ===
	sb.WriteString(sectionStyle.Render("METADATA"))
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("%s %s\n", labelStyle.Render("Name:"), valueStyle.Render(dep.Name)))
	sb.WriteString(fmt.Sprintf("%s %s\n", labelStyle.Render("Namespace:"), valueStyle.Render(dep.Namespace)))
	sb.WriteString(fmt.Sprintf("%s %s\n", labelStyle.Render("Age:"), valueStyle.Render(dep.Age)))

	// === Status Section ===
	sb.WriteString("\n")
	sb.WriteString(sectionStyle.Render("STATUS"))
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("%s %s\n", labelStyle.Render("Ready:"), statusStyle.Render(dep.Ready)))
	sb.WriteString(fmt.Sprintf("%s %d\n", labelStyle.Render("Up-to-date:"), dep.UpToDate))
	sb.WriteString(fmt.Sprintf("%s %d\n", labelStyle.Render("Available:"), dep.Available))
	sb.WriteString(fmt.Sprintf("%s %d\n", labelStyle.Render("Replicas:"), dep.Replicas))

	// === Strategy Section ===
	if dep.Strategy != "" {
		sb.WriteString("\n")
		sb.WriteString(sectionStyle.Render("STRATEGY"))
		sb.WriteString("\n")
		sb.WriteString(fmt.Sprintf("%s %s\n", labelStyle.Render("Type:"), valueStyle.Render(dep.Strategy)))
	}

	// === Images Section ===
	if len(dep.Images) > 0 {
		sb.WriteString("\n")
		sb.WriteString(sectionStyle.Render(fmt.Sprintf("IMAGES (%d)", len(dep.Images))))
		sb.WriteString("\n")
		for _, img := range dep.Images {
			sb.WriteString(fmt.Sprintf("  %s\n", truncateString(img, m.width-6)))
		}
	}

	// === Labels Section ===
	if len(dep.Labels) > 0 {
		sb.WriteString("\n")
		sb.WriteString(sectionStyle.Render("LABELS"))
		sb.WriteString("\n")

		keys := make([]string, 0, len(dep.Labels))
		for k := range dep.Labels {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			v := dep.Labels[k]
			keyStr := truncateString(k, 30)
			valStr := truncateString(v, 40)
			sb.WriteString(fmt.Sprintf("  %s=%s\n", keyStr, valStr))
		}
	}

	// === Selector Section ===
	if len(dep.Selector) > 0 {
		sb.WriteString("\n")
		sb.WriteString(sectionStyle.Render("SELECTOR"))
		sb.WriteString("\n")

		keys := make([]string, 0, len(dep.Selector))
		for k := range dep.Selector {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			v := dep.Selector[k]
			sb.WriteString(fmt.Sprintf("  %s=%s\n", k, v))
		}
	}

	// === Conditions Section ===
	if len(dep.Conditions) > 0 {
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

		for _, cond := range dep.Conditions {
			var statusColor lipgloss.Style
			if cond.Status == "True" {
				statusColor = lipgloss.NewStyle().Foreground(colorSuccess)
			} else {
				statusColor = lipgloss.NewStyle().Foreground(colorError)
			}

			sb.WriteString(fmt.Sprintf("  %-20s %s %-20s %s\n",
				truncateString(cond.Type, 20),
				statusColor.Render(fmt.Sprintf("%-8s", cond.Status)),
				truncateString(cond.Reason, 20),
				cond.LastTransition))

			if cond.Message != "" {
				sb.WriteString(fmt.Sprintf("    %s\n", truncateString(cond.Message, m.width-10)))
			}
		}
	}

	return sb.String()
}

// ScrollPercent returns the scroll percentage
func (m *DeploymentDetailsModel) ScrollPercent() float64 {
	return m.viewport.ScrollPercent()
}

// Deployment returns the current deployment
func (m *DeploymentDetailsModel) Deployment() *domain.Deployment {
	return m.deployment
}
