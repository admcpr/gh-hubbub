package models

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHalf(t *testing.T) {
	tests := []struct {
		arg  int
		want int
	}{
		{4, 2},
		{5, 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Half of %v", tt.arg), func(t *testing.T) {
			assert.Equal(t, tt.want, half(tt.arg))
		})
	}
}

func TestQuarter(t *testing.T) {
	tests := []struct {
		arg  int
		want int
	}{
		{4, 1},
		{5, 1},
		{8, 2},
		{11, 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Quarter of %v", tt.arg), func(t *testing.T) {
			assert.Equal(t, tt.want, quarter(tt.arg))
		})
	}
}

func TestMax(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Max", args{1, 2}, 2},
		{"Max", args{2, 1}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Min", args{1, 2}, 1},
		{"Min", args{2, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}
