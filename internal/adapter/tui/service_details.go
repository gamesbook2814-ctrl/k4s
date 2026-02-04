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

// ServiceDetailsModel is the model for service details view
type ServiceDetailsModel struct {
	service  *domain.Service
	viewport viewport.Model
	styles   Styles
	width    int
	height   int
	ready    bool
}

// NewServiceDetailsModel creates a new service details model
func NewServiceDetailsModel(styles Styles) ServiceDetailsModel {
	return ServiceDetailsModel{
		styles: styles,
	}
}

// SetService sets the service to display
func (m *ServiceDetailsModel) SetService(svc *domain.Service) {
	m.service = svc
	if m.ready {
		m.viewport.SetContent(m.renderContent())
		m.viewport.GotoTop()
	}
}

// SetSize sets the viewport size
func (m *ServiceDetailsModel) SetSize(width, height int) {
	m.width = width
	m.height = height
	m.viewport = viewport.New(width, height)
	m.viewport.Style = lipgloss.NewStyle()
	m.ready = true
	if m.service != nil {
		m.viewport.SetContent(m.renderContent())
	}
}

// Update handles messages
func (m ServiceDetailsModel) Update(msg tea.Msg) (ServiceDetailsModel, tea.Cmd) {
	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

// View renders the service details
func (m ServiceDetailsModel) View() string {
	if !m.ready || m.service == nil {
		return "Loading..."
	}
	return m.viewport.View()
}

func (m *ServiceDetailsModel) renderContent() string {
	if m.service == nil {
		return "No service selected"
	}

	var sb strings.Builder
	svc := m.service

	// Styles
	sectionStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(colorPrimary).
		MarginTop(1)

	labelStyle := lipgloss.NewStyle().
		Foreground(colorMuted).
		Width(16)

	valueStyle := lipgloss.NewStyle()

	// Type color
	var typeStyle lipgloss.Style
	switch svc.Type {
	case domain.ServiceTypeLoadBalancer:
		if svc.ExternalIP != "<pending>" && svc.ExternalIP != "<none>" {
			typeStyle = lipgloss.NewStyle().Foreground(colorSuccess).Bold(true)
		} else {
			typeStyle = lipgloss.NewStyle().Foreground(colorWarning).Bold(true)
		}
	case domain.ServiceTypeNodePort:
		typeStyle = lipgloss.NewStyle().Foreground(colorSecondary).Bold(true)
	default:
		typeStyle = lipgloss.NewStyle().Foreground(colorMuted).Bold(true)
	}

	// === Metadata Section ===
	sb.WriteString(sectionStyle.Render("METADATA"))
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("%s %s\n", labelStyle.Render("Name:"), valueStyle.Render(svc.Name)))
	sb.WriteString(fmt.Sprintf("%s %s\n", labelStyle.Render("Namespace:"), valueStyle.Render(svc.Namespace)))
	sb.WriteString(fmt.Sprintf("%s %s\n", labelStyle.Render("Age:"), valueStyle.Render(svc.Age)))

	// === Service Info Section ===
	sb.WriteString("\n")
	sb.WriteString(sectionStyle.Render("SERVICE"))
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("%s %s\n", labelStyle.Render("Type:"), typeStyle.Render(svc.Type)))
	sb.WriteString(fmt.Sprintf("%s %s\n", labelStyle.Render("Cluster IP:"), valueStyle.Render(svc.ClusterIP)))
	sb.WriteString(fmt.Sprintf("%s %s\n", labelStyle.Render("External IP:"), valueStyle.Render(svc.ExternalIP)))

	// === Ports Section ===
	if len(svc.PortDetails) > 0 {
		sb.WriteString("\n")
		sb.WriteString(sectionStyle.Render(fmt.Sprintf("PORTS (%d)", len(svc.PortDetails))))
		sb.WriteString("\n")

		// Header
		portHeader := lipgloss.NewStyle().
			Foreground(colorMuted).
			Bold(true).
			Render(fmt.Sprintf("  %-15s %-8s %-12s %-12s %s", "NAME", "PORT", "TARGET", "NODEPORT", "PROTOCOL"))
		sb.WriteString(portHeader)
		sb.WriteString("\n")

		for _, p := range svc.PortDetails {
			name := p.Name
			if name == "" {
				name = "<unnamed>"
			}

			nodePort := ""
			if p.NodePort > 0 {
				nodePort = fmt.Sprintf("%d", p.NodePort)
			} else {
				nodePort = "-"
			}

			sb.WriteString(fmt.Sprintf("  %-15s %-8d %-12s %-12s %s\n",
				truncateString(name, 15),
				p.Port,
				p.TargetPort,
				nodePort,
				p.Protocol))
		}
	}

	// === Selector Section ===
	if len(svc.Selector) > 0 {
		sb.WriteString("\n")
		sb.WriteString(sectionStyle.Render("SELECTOR"))
		sb.WriteString("\n")

		keys := make([]string, 0, len(svc.Selector))
		for k := range svc.Selector {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			v := svc.Selector[k]
			sb.WriteString(fmt.Sprintf("  %s=%s\n", k, v))
		}
	}

	// === Labels Section ===
	if len(svc.Labels) > 0 {
		sb.WriteString("\n")
		sb.WriteString(sectionStyle.Render("LABELS"))
		sb.WriteString("\n")

		keys := make([]string, 0, len(svc.Labels))
		for k := range svc.Labels {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			v := svc.Labels[k]
			keyStr := truncateString(k, 30)
			valStr := truncateString(v, 40)
			sb.WriteString(fmt.Sprintf("  %s=%s\n", keyStr, valStr))
		}
	}

	return sb.String()
}

// ScrollPercent returns the scroll percentage
func (m *ServiceDetailsModel) ScrollPercent() float64 {
	return m.viewport.ScrollPercent()
}

// Service returns the current service
func (m *ServiceDetailsModel) Service() *domain.Service {
	return m.service
}
