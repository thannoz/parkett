package engine

import "testing"

func TestSideString(t *testing.T) {
	tests := []struct {
		name string
		side Side
		want string
	}{
		{name: "buy", side: Buy, want: "buy"},
		{name: "sell", side: Sell, want: "sell"},
		{name: "unknown", side: Side(42), want: "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.side.String(); got != tt.want {
				t.Errorf("Side.String() = %q, want %q", got, tt.want)
			}
		})
	}
}
