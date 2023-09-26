package filters

import (
	"gh-hubbub/structs"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestNewFilterBoolModel(t *testing.T) {
	tests := []struct {
		tab   string
		title string
		value bool
	}{
		{"Tab1", "Test 1", true},
		{"Tab1", "Test 2", false},
		{"Tab1", "Test 3", true},
		{"Tab1", "Test 4", false},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			m := NewBoolModel(tt.tab, tt.title, tt.value)

			if m.Title != tt.title {
				t.Errorf("got %q, want %q", m.Title, tt.title)
			}

			if m.input.Value() != structs.YesNo(tt.value) {
				t.Errorf("got %q, want %q", m.input.Value(), structs.YesNo(tt.value))
			}
		})
	}
}

func TestFilterBoolModel_Update(t *testing.T) {
	trueModel := NewBoolModel("Tab 1", "True", true)
	falseModel := NewBoolModel("Tab 1", "False", false)

	tests := []struct {
		model  BoolModel
		title  string
		msgKey rune
		want   bool
	}{
		{title: "'n' should set value to false", model: trueModel, msgKey: 'n', want: false},
		{title: "'N' should set value to false", model: trueModel, msgKey: 'N', want: false},
		{title: "'y' should set value to true", model: falseModel, msgKey: 'y', want: true},
		{title: "'Y' should set value to true", model: falseModel, msgKey: 'Y', want: true},
		{title: "'x' shouldn't change false value", model: falseModel, msgKey: 'x', want: false},
		{title: "'x' shouldn't change true value", model: trueModel, msgKey: 'x', want: true},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{tt.msgKey}}
			m, _ := tt.model.Update(msg)

			filterBooleanModel, _ := m.(BoolModel)
			got := filterBooleanModel.GetValue()

			if got != tt.want {
				t.Errorf("FilterBooleanModel.GetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
