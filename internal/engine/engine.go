// Package engine implements the in-memory matching engine: one order book
// per symbol, price-time-priority matching, single writer goroutine per book.
package engine

// Side is the side of an order: buy (bid) or sell (ask).
type Side uint8

const (
	Buy Side = iota
	Sell
)

func (s Side) String() string {
	switch s {
	case Buy:
		return "buy"
	case Sell:
		return "sell"
	default:
		return "unknown"
	}
}
