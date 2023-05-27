package validators

import "testing"

func TestIsCredit(t *testing.T) {
	testCases := []struct {
		amount   string
		expected bool
	}{
		{"+100", true},
		{"-100", false},
		{"0", false},
		{"", false},
	}

	for _, tc := range testCases {
		result := IsCredit(tc.amount)
		if result != tc.expected {
			t.Errorf("IsCredit(%s) = %v, want %v", tc.amount, result, tc.expected)
		}
	}
}
