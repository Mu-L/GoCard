// File: internal/ui/menu.go

package ui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/DavidMiserak/GoCard/internal/data"
)

// Define key mappings
type keyMap struct {
	Up    key.Binding
	Down  key.Binding
	Enter key.Binding
	Quit  key.Binding
}

var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"), // "k" for Vim users
		key.WithHelp("↑/k", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"), // "j" for Vim users
		key.WithHelp("↓/j", "down"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

// MainMenu represents the main menu model
type MainMenu struct {
	items    []string
	cursor   int
	selected int
	width    int
	height   int
	store    *data.Store // Added store field
}

// NewMainMenu creates a new main menu
func NewMainMenu(store *data.Store) *MainMenu {
	// If no store provided, create a new one with dummy data
	if store == nil {
		store = data.NewStore()
	}

	return &MainMenu{
		items:    []string{"Study", "Browse Decks", "Statistics", "Quit"},
		cursor:   0,
		selected: -1,
		store:    store,
	}
}

// Init initializes the main menu
func (m MainMenu) Init() tea.Cmd {
	return nil
}

// Update handles user input and updates the model
func (m MainMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, keys.Up):
			if m.cursor > 0 {
				m.cursor--
			}

		case key.Matches(msg, keys.Down):
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}

		case key.Matches(msg, keys.Enter):
			m.selected = m.cursor

			// Handle menu selection
			switch m.cursor {
			case 0: // Study
				// Navigate to study screen
				return NewBrowseScreen(m.store), nil

			case 1: // Browse Decks
				// Navigate to browse decks screen
				return NewBrowseScreen(m.store), nil

			case 2: // Statistics
				// Navigate to statistics screen
				return NewStatisticsScreen(m.store), nil

			case 3: // Quit
				return m, tea.Quit
			}
		}

	case tea.WindowSizeMsg:
		m.width = 120 // Default width
		m.height = msg.Height
	}

	return m, nil
}

// View renders the main menu
func (m MainMenu) View() string {
	// Title and subtitle
	s := titleStyle.Render("GoCard")
	s += "\n" + subtitleStyle.Render("Terminal Flashcards")
	s += "\n\n"

	// Menu items
	for i, item := range m.items {
		if i == m.cursor {
			s += selectedItemStyle.Render("> " + item)
		} else {
			s += normalItemStyle.Render("  " + item)
		}
		s += "\n"
	}

	// Help
	s += "\n" + helpStyle.Render("\t↑/↓: Navigate"+"\tEnter: Select"+"\tq: Quit")

	return s
}
