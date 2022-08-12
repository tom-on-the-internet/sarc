package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	formats []string
	cursor  int
	text    string
}

func initialModel(format, text string) model {
	cursor := 0

	for i, f := range getFormats() {
		if format == f {
			cursor = i
		}
	}

	return model{formats: getFormats(), cursor: cursor, text: text}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.formats)-1 {
				m.cursor++
			}

		case "enter", " ":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	menu := m.RenderMenu()
	result := formats[m.formats[m.cursor]](m.text)
	result = lipgloss.NewStyle().Padding(1, 2, 1, 2).Margin(1, 0, 1, 0).Render(result)

	return lipgloss.JoinHorizontal(lipgloss.Center, menu, result)
}

func (m model) RenderMenu() string {
	str := ""

	for i, format := range m.formats {
		if m.cursor == i {
			format = lipgloss.NewStyle().Foreground(lipgloss.Color("5")).Render(format)
		}

		if i != 0 {
			str += "\n"
		}

		str += format
	}

	return lipgloss.NewStyle().
		Padding(1, 2, 1, 2).
		Margin(1, 0, 1, 0).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("5")).
		Render(str)
}
