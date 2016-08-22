package util

import "testing"

func TestEncode(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		src string
		// Expected results.
		wantDst string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if gotDst := Encode(tt.src); gotDst != tt.wantDst {
			t.Errorf("%q. Encode() = %v, want %v", tt.name, gotDst, tt.wantDst)
		}
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		src string
		// Expected results.
		wantDst string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if gotDst := Decode(tt.src); gotDst != tt.wantDst {
			t.Errorf("%q. Decode() = %v, want %v", tt.name, gotDst, tt.wantDst)
		}
	}
}
