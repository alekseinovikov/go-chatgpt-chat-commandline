package tui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg error
)

type apiKeyModel struct {
	textInput textinput.Model
	err       error
	result    string
}

func initialModel() *apiKeyModel {
	ti := textinput.New()
	ti.Placeholder = "***"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return &apiKeyModel{
		textInput: ti,
		err:       nil,
	}
}
func (m apiKeyModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m apiKeyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			m.result = m.textInput.Value()
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m apiKeyModel) View() string {
	return fmt.Sprintf(
		"Enter your OpenAI API Key:\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}

func GetApiKey() (string, error) {
	p := tea.NewProgram(initialModel())
	if result, err := p.Run(); err != nil {
		return "", err
	} else {
		apiModel := result.(apiKeyModel)
		return apiModel.result, nil
	}
}
