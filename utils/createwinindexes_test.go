package utils

import (
	"testing"
)

func TestCreateWinIndexes(t *testing.T) {
	tests := []struct {
		difficulty int
	}{
		{difficulty: 1},
		{difficulty: 2},
		{difficulty: 3},
		{difficulty: 4},
	}
	for i := 0; i < 50; i++ {
		for _, tt := range tests {
		got := CreateWinIndexes(tt.difficulty)
		if len(got) != tt.difficulty {
			t.Fatalf("diff=%d len=%d", tt.difficulty, len(got))
		}

		for _, value := range got {
			if value < 1 || value > 10 {
				t.Fatalf("out of range: %d", value)
			}
		}

		seen := make(map[int]struct{})
		for _, v := range got {
    		if _, ok := seen[v]; ok {
        		t.Fatalf("duplicate value: %d", v)
    		}
    		seen[v] = struct{}{}
		}
	}
	}


}
