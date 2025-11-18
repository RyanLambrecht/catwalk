package components

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type ListMenuModel struct {
	choices []string
	cursor  int
	header  string
}

func NewListMenuModel(header string, choices []string) *ListMenuModel {
	return &ListMenuModel{
		header:  header,
		choices: choices,
	}
}

func (m *ListMenuModel) Init() tea.Cmd {
	return nil
}

type MenuSelectedMsg struct{ Index int }

type MenuBackMsg struct{}

func (m *ListMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {

		case "backspace", "delete":
			return m, func() tea.Msg { return MenuBackMsg{} }

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case " ", "enter":
			return m, func() tea.Msg { return MenuSelectedMsg{m.cursor} }
		}

	}
	return m, nil
}

func (m *ListMenuModel) View() string {
	s := fmt.Sprintf("%s\n\n", m.header)

	for i, choice := range m.choices {

		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice)

	}

	s += "\nq to quit\n"
	return s
}
