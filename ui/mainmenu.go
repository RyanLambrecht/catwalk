package ui

import (
	"github.com/RyanLambrecht/catwalk/ui/codex"
	"github.com/RyanLambrecht/catwalk/ui/components"
	tea "github.com/charmbracelet/bubbletea"
)

type mainMenuModel struct {
	menu *components.ListMenuModel
}

func NewMainMenuModel() *mainMenuModel {
	return &mainMenuModel{
		menu: components.NewListMenuModel("Main Menu", []string{"Codex"}),
	}
}

func (m *mainMenuModel) Init() tea.Cmd {
	return m.menu.Init()
}

func (m *mainMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	newMenu, cmd := m.menu.Update(msg)
	m.menu = newMenu.(*components.ListMenuModel)

	switch msg := msg.(type) {
	case components.MenuSelectedMsg:
		switch msg.Index {
		case 0:
			return m, func() tea.Msg { return PushModelMsg{codex.NewCodexModel()} }
		}
	}

	return m, cmd
}

func (m *mainMenuModel) View() string {
	return m.menu.View()
}
