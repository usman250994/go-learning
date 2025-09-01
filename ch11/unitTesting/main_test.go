package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		a      int
		b      int
		expect int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{2, 3, 5},
	}

	for _, tt := range tests {
		result := add(tt.a, tt.b)
		if result != tt.expect {
			t.Errorf("Expected %d, got %d", tt.expect, result)
		}
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		a      int
		b      int
		expect int
	}{
		{2, 3, 6},
		{4, 5, 20},
		{6, 7, 42},
	}

	for _, tt := range tests {
		result := multiply(tt.a, tt.b)
		if result != tt.expect {
			t.Errorf("Expected %d, got %d", tt.expect, result)
		}
	}
}

// Benchmark test
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {

		result := add(1, 4)
		if result != 5 {
			b.Errorf("Expected %d, got %d", 5, result)
		}
	}
}
