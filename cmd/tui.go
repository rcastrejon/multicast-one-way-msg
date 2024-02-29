package cmd

import (
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/rcastrejon/multicast-channels/pkg/multicast"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	srv       *multicast.Server
	table     table.Model
	textInput textinput.Model
	chosen    bool
	choice    string
}

func initialModel(srv *multicast.Server) model {
	cols, rows := srv.BuildRoomsTable()

	t := table.New(
		table.WithColumns(cols),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(5),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	ti := textinput.New()
	ti.Placeholder = "Type your message..."
	ti.Focus()
	ti.CharLimit = 256
	ti.Width = 30

	m := model{srv, t, ti, false, ""}
	return m
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
	}

	if !m.chosen {
		return updateChoices(msg, m)
	}
	return updateChosen(msg, m)
}

func (m model) View() string {
	if !m.chosen {
		return choicesView(m)
	} else {
		return chosenView(m)
	}
}

func updateChoices(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			m.chosen = true
			m.choice = m.table.SelectedRow()[1]
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func updateChosen(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			m.chosen = false
		case tea.KeyEnter:
			m.srv.SendTo(m.choice, m.textInput.Value())
			m.textInput.Reset()
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func choicesView(m model) string {
	return fmt.Sprintf("%s\n\n(esc to quit)", baseStyle.Render(m.table.View())) + "\n"
}

func chosenView(m model) string {
	return fmt.Sprintf(
		"Connected to %s\n\n%s\n\n%s",
		m.choice,
		m.textInput.View(),
		"(esc to go back)",
	) + "\n"
}
