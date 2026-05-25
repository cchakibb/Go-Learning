package math

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct{
		a,b 		int
		expected 	int
	}{
		{1, 3, 4},
		{6, 3, 9},
		{-2, 2, 0},
	}

	for _, tt := range tests {
		result := Add(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Add(%d,%d) = %d, want %d", tt.a, tt.b, result, tt.expected)
		}
	}
}

func TestDivide(t *testing.T) {
	tests:= []struct{
		a, b		float64
		expected	float64
		wantErr		bool
	}{
		{12, 4, 3, false},
		{10, 2, 5, false},
		{21, 10, 2.1, false},
		{0, 3, 0, false},
		{10, 0, 0, true},
	}

	for _, tt := range tests {
		if tt.wantErr {
			_, errErr := Divide(tt.a, tt.b)
			if errErr == nil {
    			t.Errorf("Divide(%f, %f) should have returned an error", tt.a, tt.b)
			}

		} else {
			res, err := Divide(tt.a, tt.b)
			if err != nil {
				t.Errorf("Divide(%f, %f) unexpected error: %s", tt.a, tt.b, err)
			}
			if res != tt.expected {
				t.Errorf("Divide(%f,%f) = %f, want %f", tt.a, tt.b, res, tt.expected)
			}
		}
	}
}