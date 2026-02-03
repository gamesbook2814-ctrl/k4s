package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// SearchInput is a search input component for logs
type SearchInput struct {
	visible      bool
	input        textinput.Model
	matchCount   int
	currentMatch int
}

// NewSearchInput creates a new search input
func NewSearchInput() SearchInput {
	ti := textinput.New()
	ti.Placeholder = "Search..."
	ti.CharLimit = 100
	ti.Width = 30

	return SearchInput{
		input: ti,
	}
}

// Show displays the search input
func (s *SearchInput) Show() {
	s.visible = true
	s.input.Reset()
	s.input.Focus()
	s.matchCount = 0
	s.currentMatch = 0
}

// Hide hides the search input
func (s *SearchInput) Hide() {
	s.visible = false
	s.input.Blur()
}

// IsVisible returns true if the search input is visible
func (s *SearchInput) IsVisible() bool {
	return s.visible
}

// Query returns the current search query
func (s *SearchInput) Query() string {
	return s.input.Value()
}

// SetMatchCount sets the number of matches found
func (s *SearchInput) SetMatchCount(count int) {
	s.matchCount = count
	if count > 0 && s.currentMatch == 0 {
		s.currentMatch = 1
	}
	if count == 0 {
		s.currentMatch = 0
	}
}

// CurrentMatch returns the current match index (1-based)
func (s *SearchInput) CurrentMatch() int {
	return s.currentMatch
}

// NextMatch moves to the next match
func (s *SearchInput) NextMatch() int {
	if s.matchCount == 0 {
		return 0
	}
	s.currentMatch++
	if s.currentMatch > s.matchCount {
		s.currentMatch = 1
	}
	return s.currentMatch
}

// PrevMatch moves to the previous match
func (s *SearchInput) PrevMatch() int {
	if s.matchCount == 0 {
		return 0
	}
	s.currentMatch--
	if s.currentMatch < 1 {
		s.currentMatch = s.matchCount
	}
	return s.currentMatch
}

// Update handles input messages
// Returns (query, submitted, cancelled, cmd)
func (s *SearchInput) Update(msg tea.Msg) (string, bool, bool, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "enter":
			return s.input.Value(), true, false, nil
		case "esc":
			return "", false, true, nil
		}
	}

	var cmd tea.Cmd
	s.input, cmd = s.input.Update(msg)
	return s.input.Value(), false, false, cmd
}

// View renders the search input
func (s *SearchInput) View() string {
	if !s.visible {
		return ""
	}

	promptStyle := lipgloss.NewStyle().
		Foreground(colorPrimary).
		Bold(true)

	matchStyle := lipgloss.NewStyle().
		Foreground(colorMuted)

	var sb strings.Builder
	sb.WriteString(promptStyle.Render("/"))
	sb.WriteString(s.input.View())

	if s.matchCount > 0 {
		sb.WriteString(matchStyle.Render(" "))
		sb.WriteString(matchStyle.Render("["))
		sb.WriteString(matchStyle.Render(string(rune('0' + s.currentMatch))))
		sb.WriteString(matchStyle.Render("/"))
		sb.WriteString(matchStyle.Render(string(rune('0' + s.matchCount))))
		sb.WriteString(matchStyle.Render("]"))
	} else if s.input.Value() != "" {
		sb.WriteString(matchStyle.Render(" [no matches]"))
	}

	return sb.String()
}

// ViewInline renders the search input inline (for footer)
func (s *SearchInput) ViewInline() string {
	if !s.visible {
		return ""
	}

	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colorPrimary).
		Padding(0, 1)

	return boxStyle.Render(s.View())
}
