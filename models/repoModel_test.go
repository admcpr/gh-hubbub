package models

import "testing"

func TestRepoModel_NewRepoModel(t *testing.T) {
	type args struct {
		width  int
		height int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Test NewRepoModel", args: args{width: 1024, height: 768}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRepoModel(tt.args.width, tt.args.height)
			if got.width != tt.args.width {
				t.Errorf("NewRepoModel width got = %v, want %v", got.width, tt.args.width)
			}
			if got.height != tt.args.height {
				t.Errorf("NewRepoModel height got = %v, want %v", got.height, tt.args.height)
			}
			if got.keys.Esc.Enabled() != true {
				t.Errorf("NewRepoModel keys.Esc.Enabled got = %v, want %v", got.keys.Esc.Enabled(), true)
			}
		})
	}
}
