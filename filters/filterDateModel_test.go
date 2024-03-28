package filters

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// func TestNewFilterDateModel(t *testing.T) {
// 	const title = "Title"
// 	fromString := "2022-01-01"
// 	toString := "2022-12-31"
// 	from, _ := time.Parse("2006-01-02", fromString)
// 	to, _ := time.Parse("2006-01-02", toString)

// 	t.Run("NewFilterDateModel", func(t *testing.T) {
// 		m := NewFilterDateModel(title, from, to)
// 		assert.Equal(t, m.Title, title)
// 		assert.Equal(t, m.fromInput.Placeholder, fromString)
// 		assert.Equal(t, m.toInput.Placeholder, toString)

// 		gotFrom, gotTo := m.GetValue()

// 		assert.Equal(t, gotFrom, from)
// 		assert.Equal(t, gotTo, to)
// 	})
// }

func TestNewFilterDateModel(t *testing.T) {
	const tab = "Tab"
	const title = "Title"
	fromString := "2022-01-01"
	toString := "2022-12-31"
	from, _ := time.Parse("2006-01-02", fromString)
	to, _ := time.Parse("2006-01-02", toString)

	t.Run("NewFilterDateModel", func(t *testing.T) {
		m := NewDateModel(tab, title, from, to)
		assert.Equal(t, m.Tab, tab)
		assert.Equal(t, m.Title, title)
		assert.Equal(t, m.fromInput.Placeholder, fromString)
		assert.Equal(t, m.toInput.Placeholder, toString)
	})
}

func TestDateValidator(t *testing.T) {
	errorMessage := fmt.Errorf("please enter a YYYY-MM-DD date for `from`")

	tests := []struct {
		name   string
		input  string
		prompt string
		want   error
	}{
		{name: "Valid date", input: "2022-01-01", prompt: "from", want: nil},
		{name: "Invalid date format", input: "01-01-2022", prompt: "from", want: errorMessage},
		{name: "Invalid date value", input: "2022-13-01", prompt: "from", want: errorMessage},
		{name: "Too long input", input: "2022-01-01T00:00:00Z", prompt: "from", want: errorMessage},
		{name: "Invalid characters", input: "2022-01-0a", prompt: "from", want: errorMessage},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dateValidator(tt.input, tt.prompt)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestFilterDateModel_GetValue(t *testing.T) {
	tests := []struct {
		name      string
		fromValue string
		toValue   string
		fromError error
		toError   error
		wantFrom  time.Time
		wantTo    time.Time
		wantErr   error
	}{
		{
			name:      "Valid input",
			fromValue: "2022-01-01",
			toValue:   "2022-12-31",
			wantFrom:  time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			wantTo:    time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
			wantErr:   nil,
		},
		{
			name:      "Invalid from input",
			fromValue: "01-01-2022",
			toValue:   "2022-12-31",
			wantErr:   fmt.Errorf("please enter a YYYY-MM-DD date for `From:`"),
		},
		{
			name:      "Invalid to input",
			fromValue: "2022-01-01",
			toValue:   "31-12-2022",
			wantErr:   fmt.Errorf("please enter a YYYY-MM-DD date for `To:`"),
		},
		{
			name:      "Invalid from and to input",
			fromValue: "01-01-2022",
			toValue:   "31-12-2022",
			wantErr:   fmt.Errorf("please enter a YYYY-MM-DD date for `From:`"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			m := NewDateModel("Tab", tt.name, time.Time{}, time.Time{})

			m.fromInput.SetValue(tt.fromValue)
			m.toInput.SetValue(tt.toValue)

			gotFrom, gotTo, gotErr := m.GetValue()

			assert.Equal(t, gotFrom, tt.wantFrom)
			assert.Equal(t, gotTo, tt.wantTo)
			assert.Equal(t, gotErr, tt.wantErr)
		})
	}
}

// func TestFilterDateModel_View(t *testing.T) {
// 	// 	type args struct {
// 	// 		title string
// 	// 		from  time.Time
// 	// 		to    time.Time
// 	// 	}
// 	// 	tests := []struct {
// 	// 		name string
// 	// 		args args
// 	// 		want FilterDateModel
// 	// 	}{
// 	// 		{name: "1", args: {"Title", time.Now(), time.Now() }, FilterDateModel{title: "Title", from: time.Now(), to: time.Now()}}
// 	// 	}
// 	// 	for _, tt := range tests {
// 	// 		t.Run(tt.name, func(t *testing.T) {
// 	// 			if got := NewFilterDateModel(tt.args.title, tt.args.from, tt.args.to); !reflect.DeepEqual(got, tt.want) {
// 	// 				t.Errorf("NewFilterDateModel() = %v, want %v", got, tt.want)
// 	// 			}
// 	// 		})
// 	// 	}

// 	t.Run("View", func(t *testing.T) {
// 		m := NewFilterDateModel(title, from, to)
// 		m.View()
// 	})

// 	t.Run("Focus", func(t *testing.T) {
// 		m := NewFilterDateModel(title, from, to)
// 		m.Focus()
// 	})
// }
