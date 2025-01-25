package shared

import tea "github.com/charmbracelet/bubbletea/v2"

type NextMessage struct{ ModelData interface{} }
type PreviousMessage struct{ Message tea.Msg }
