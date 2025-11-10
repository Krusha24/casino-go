package guessnumber

import (
	"math"
	"testing"
)

const eps = 1e-9

func TestCalculateNetWinnings(t *testing.T) {
	tests := []struct {
		difficulty int
		bet        float64
		want       float64
	}{
		{difficulty: 1, bet: 10, want: 90},
		{difficulty: 2, bet: 10, want: 35},
		{difficulty: 3, bet: 10, want: 15},
		{difficulty: 4, bet: 10, want: 7},
		{difficulty: 1, bet: 7.5, want: 67.5},
		{difficulty: 0, bet: 10, want: 0},
		{difficulty: 5, bet: 10, want: 0},
	}

	for _, tt := range tests {
		got := calculateNetWinnings(tt.difficulty, tt.bet)
		if math.Abs(got-tt.want) > eps {
			t.Fatalf("diff=%d bet=%.2f: got %.2f, want %.2f", tt.difficulty, tt.bet, got, tt.want)
		}
	}

}
