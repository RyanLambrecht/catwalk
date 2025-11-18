package codex

import (
	"github.com/RyanLambrecht/catwalk/ui/components"
	tea "github.com/charmbracelet/bubbletea"
)

type CodexModel struct {
	menu *components.ListMenuModel
}

func NewCodexModel() *CodexModel {
	return &CodexModel{
		menu: components.NewListMenuModel("Codex", []string{"Items", "Recipes", "Buildings"}),
	}
}

func (m *CodexModel) Init() tea.Cmd {
	return m.menu.Init()
}

func (m *CodexModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_, cmd := m.menu.Update(msg)
	//m.menu = newMenu.(*listMenuModel)

	//switch msg := msg.(type) {
	//case MenuSelectedMsg:
	//	switch msg.index {
	//	case 0:
	//		return m, func() tea.Msg { return PushModelMsg{NewMenuAModel()} }
	//	case 1:
	//		return m, func() tea.Msg { return PushModelMsg{NewMenuBModel()} }
	//	}
	//}

	return m, cmd
}

func (m *CodexModel) View() string {
	return m.menu.View()
}
