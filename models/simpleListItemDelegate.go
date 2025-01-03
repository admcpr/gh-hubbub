package models

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/v2/list"
	tea "github.com/charmbracelet/bubbletea/v2"
)

type simpleItem string

func (i simpleItem) FilterValue() string { return "" }

type simpleItemDelegate struct{}

func (d simpleItemDelegate) Height() int                             { return 1 }
func (d simpleItemDelegate) Spacing() int                            { return 1 }
func (d simpleItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d simpleItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(simpleItem)
	if !ok {
		fmt.Fprintf(w, "invalid item type: %T", listItem)
		return
	}

	str := string(i)

	renderFunction := itemStyle.Render
	if index == m.Index() {
		renderFunction = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, renderFunction(str))
}
