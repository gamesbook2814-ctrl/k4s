package tui

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/LywwKkA-aD/k4s/internal/domain"
)

// kubeConfigItem implements list.Item for kubeconfigs
type kubeConfigItem struct {
	kubeConfig domain.KubeConfig
}

func (i kubeConfigItem) FilterValue() string { return i.kubeConfig.Name }

// kubeConfigDelegate renders kubeconfig list items
type kubeConfigDelegate struct {
	styles Styles
}

func (d kubeConfigDelegate) Height() int                             { return 2 }
func (d kubeConfigDelegate) Spacing() int                            { return 1 }
func (d kubeConfigDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }

func (d kubeConfigDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	item, ok := listItem.(kubeConfigItem)
	if !ok {
		return
	}

	name := item.kubeConfig.Name
	path := item.kubeConfig.Path

	var nameStyle, pathStyle lipgloss.Style
	if index == m.Index() {
		nameStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(colorPrimary)
		pathStyle = lipgloss.NewStyle().
			Foreground(colorMuted)
		fmt.Fprintf(w, "▸ %s\n  %s", nameStyle.Render(name), pathStyle.Render(path))
	} else {
		nameStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF"))
		pathStyle = lipgloss.NewStyle().
			Foreground(colorMuted)
		fmt.Fprintf(w, "  %s\n  %s", nameStyle.Render(name), pathStyle.Render(path))
	}

	if item.kubeConfig.Default {
		fmt.Fprint(w, " (default)")
	}
}

// newKubeConfigList creates a list model for kubeconfigs
func newKubeConfigList(configs []domain.KubeConfig, width, height int, styles Styles) list.Model {
	items := make([]list.Item, len(configs))
	for i, cfg := range configs {
		items[i] = kubeConfigItem{kubeConfig: cfg}
	}

	delegate := kubeConfigDelegate{styles: styles}
	l := list.New(items, delegate, width, height)
	l.Title = "Select Kubeconfig"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.SetShowHelp(false)
	l.Styles.Title = styles.Title
	l.Styles.FilterPrompt = lipgloss.NewStyle().Foreground(colorPrimary)
	l.Styles.FilterCursor = lipgloss.NewStyle().Foreground(colorPrimary)

	return l
}
