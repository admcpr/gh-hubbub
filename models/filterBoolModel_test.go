package models

import (
	"gh-hubbub/structs"
	"testing"

	tea "github.com/charmbracelet/bubbletea/v2"
)

func TestNewFilterBoolModel(t *testing.T) {
	tests := []struct {
		name  string
		value bool
	}{
		{"Test 1", true},
		{"Test 2", false},
		{"Test 3", true},
		{"Test 4", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewBoolModel(tt.name, tt.value)

			if m.Name != tt.name {
				t.Errorf("got %q, want %q", m.Name, tt.name)
			}

			if m.input.Value() != structs.YesNo(tt.value) {
				t.Errorf("got %q, want %q", m.input.Value(), structs.YesNo(tt.value))
			}
		})
	}
}

func TestFilterBoolModel_Update(t *testing.T) {
	trueModel := NewBoolModel("True", true)
	falseModel := NewBoolModel("False", false)

	tests := []struct {
		model BoolModel
		name  string
		_rune rune
		want  bool
	}{
		{name: "'n' should set value to false", model: trueModel, _rune: 'n', want: false},
		{name: "'N' should set value to false", model: trueModel, _rune: 'N', want: false},
		{name: "'y' should set value to true", model: falseModel, _rune: 'y', want: true},
		{name: "'Y' should set value to true", model: falseModel, _rune: 'Y', want: true},
		{name: "'x' shouldn't change false value", model: falseModel, _rune: 'x', want: false},
		{name: "'x' shouldn't change true value", model: trueModel, _rune: 'x', want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keyReleaseMsg := tea.KeyPressMsg{Code: tt._rune}
			m, _ := tt.model.Update(keyReleaseMsg)

			filterBooleanModel, _ := m.(BoolModel)
			got := filterBooleanModel.GetValue()

			if got != tt.want {
				t.Errorf("FilterBooleanModel.GetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
