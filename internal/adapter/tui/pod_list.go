package tui

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/LywwKkA-aD/k4s/internal/domain"
)

// podItem implements list.Item for pods
type podItem struct {
	pod domain.Pod
}

func (i podItem) FilterValue() string { return i.pod.Name }

// podDelegate renders pod list items
type podDelegate struct {
	styles Styles
}

func (d podDelegate) Height() int                             { return 1 }
func (d podDelegate) Spacing() int                            { return 0 }
func (d podDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }

func (d podDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	item, ok := listItem.(podItem)
	if !ok {
		return
	}

	pod := item.pod

	// Status color based on pod state
	var statusStyle lipgloss.Style
	switch pod.Status {
	case domain.PodStatusRunning:
		statusStyle = lipgloss.NewStyle().Foreground(colorSuccess)
	case domain.PodStatusPending, "ContainerCreating", "PodInitializing":
		statusStyle = lipgloss.NewStyle().Foreground(colorWarning)
	case domain.PodStatusSucceeded, "Completed":
		statusStyle = lipgloss.NewStyle().Foreground(colorMuted)
	case domain.PodStatusFailed, "Error", "CrashLoopBackOff", "ImagePullBackOff", "ErrImagePull":
		statusStyle = lipgloss.NewStyle().Foreground(colorError)
	case "Terminating":
		statusStyle = lipgloss.NewStyle().Foreground(colorWarning)
	default:
		if len(pod.Status) > 5 && pod.Status[:5] == "Init:" {
			statusStyle = lipgloss.NewStyle().Foreground(colorWarning)
		} else {
			statusStyle = lipgloss.NewStyle().Foreground(colorMuted)
		}
	}

	// Restarts color
	var restartsStyle lipgloss.Style
	if pod.Restarts > 10 {
		restartsStyle = lipgloss.NewStyle().Foreground(colorError)
	} else if pod.Restarts > 0 {
		restartsStyle = lipgloss.NewStyle().Foreground(colorWarning)
	} else {
		restartsStyle = lipgloss.NewStyle().Foreground(colorMuted)
	}

	// Pad plain text FIRST, then apply styling (ANSI codes break fmt width calculation)
	namePadded := fmt.Sprintf("%-45s", truncateString(pod.Name, 45))
	readyPadded := fmt.Sprintf("%-7s", pod.Ready)
	statusPadded := fmt.Sprintf("%-12s", truncateString(pod.Status, 12))
	restartsPadded := fmt.Sprintf("%-8s", fmt.Sprintf("%d", pod.Restarts))
	agePadded := pod.Age

	// Apply colors after padding
	statusStyled := statusStyle.Render(statusPadded)
	restartsStyled := restartsStyle.Render(restartsPadded)
	ageStyled := lipgloss.NewStyle().Foreground(colorMuted).Render(agePadded)

	// Cursor and selection styling
	var line string
	if index == m.Index() {
		nameStyle := lipgloss.NewStyle().Bold(true).Foreground(colorPrimary)
		line = fmt.Sprintf("▸ %s %s %s %s %s",
			nameStyle.Render(namePadded), readyPadded, statusStyled, restartsStyled, ageStyled)
	} else {
		line = fmt.Sprintf("  %s %s %s %s %s",
			namePadded, readyPadded, statusStyled, restartsStyled, ageStyled)
	}

	fmt.Fprint(w, line)
}

// newPodList creates a list model for pods
func newPodList(pods []domain.Pod, width, height int, styles Styles) list.Model {
	items := make([]list.Item, len(pods))
	for i, pod := range pods {
		items[i] = podItem{pod: pod}
	}

	delegate := podDelegate{styles: styles}
	l := list.New(items, delegate, width, height)
	l.SetShowTitle(false)      // We render our own title
	l.SetShowStatusBar(false)  // We render our own status
	l.SetFilteringEnabled(true)
	l.SetShowHelp(false)
	l.Styles.FilterPrompt = lipgloss.NewStyle().Foreground(colorPrimary)
	l.Styles.FilterCursor = lipgloss.NewStyle().Foreground(colorPrimary)

	return l
}

// updatePodList updates the pod list items while preserving selection
func updatePodList(l *list.Model, pods []domain.Pod) {
	// Preserve current selection
	currentIndex := l.Index()
	var currentName string
	if item, ok := l.SelectedItem().(podItem); ok {
		currentName = item.pod.Name
	}

	items := make([]list.Item, len(pods))
	newIndex := 0
	for i, pod := range pods {
		items[i] = podItem{pod: pod}
		if pod.Name == currentName {
			newIndex = i
		}
	}

	l.SetItems(items)

	// Try to restore selection by name, otherwise by index
	if currentName != "" {
		l.Select(newIndex)
	} else if currentIndex < len(items) {
		l.Select(currentIndex)
	}
}

func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	if maxLen <= 3 {
		return s[:maxLen]
	}
	return s[:maxLen-3] + "..."
}
