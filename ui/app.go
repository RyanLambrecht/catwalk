package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type modelStack []tea.Model

func (s *modelStack) Push(p tea.Model) {
	*s = append(*s, p)
}

func (s *modelStack) Pop() tea.Model {
	if len(*s) == 0 {
		return nil
	}
	p := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return p
}

func (s *modelStack) Top() tea.Model {
	if len(*s) == 0 {
		return nil
	}
	return (*s)[len(*s)-1]
}

type appModel struct {
	stack modelStack
}

func NewAppModel() *appModel {
	m := &appModel{}
	m.stack.Push(NewMainMenuModel())
	return m
}

func (m *appModel) Init() tea.Cmd {
	return nil
}

type PushModelMsg struct{ tea.Model }
type PopModelMsg struct{}

func (m *appModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	top := m.stack.Top()
	if top == nil {
		return m, tea.Quit
	}

	newTop, cmd := top.Update(msg)
	m.stack[len(m.stack)-1] = newTop

	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "backspace", "delete":
			if len(m.stack) > 1 {
				m.stack.Pop()
			}

		case "ctrl+c", "q":
			return m, tea.Quit
		}

	case PushModelMsg:
		m.stack.Push(msg.Model)
	case PopModelMsg:
		m.stack.Pop()
	}

	return m, cmd
}

func (m *appModel) View() string {
	top := m.stack.Top()
	if top == nil {
		return ""
	}
	return top.View()
}
