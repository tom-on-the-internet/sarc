package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
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

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.formats)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			// TODO
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// The header
	str := "sarc TUI\n\n"

	// Iterate over our choices
	for i, format := range m.formats {
		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Render the row
		str += fmt.Sprintf("%s %s\n", cursor, format)
	}

	str += "\n" + formats[m.formats[m.cursor]](m.text) + "\n"

	// The footer
	str += "\nPress q or ESC to quit.\n"

	// Send the UI for rendering
	return str
}
