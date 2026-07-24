package engine

import (
	"testing"
)

func TestAddLimit_EmptyBook(t *testing.T) {
	b := NewBook("XOM")
	order := Order{ID: 1, Side: Buy, Price: 10000, Qty: 5, Seq: 1, Symbol: ""}

	trades := b.AddLimit(order)

	gotTrades := len(trades)
	if gotTrades != 0 {
		t.Errorf("len(trades) = %d, want %d", gotTrades, 0)
	}
	gotBids := len(b.bid)
	if gotBids != 1 {
		t.Fatalf("len(b.bid) = %d, want %d", gotBids, 1)
	}
	gotAsks := len(b.ask)
	if gotAsks != 0 {
		t.Errorf("len(b.ask) = %d, want %d", gotAsks, 0)
	}
	gotPrice := b.bid[0].Price
	if gotPrice != 10000 {
		t.Errorf("Price = %d, want %d", gotPrice, 10000)
	}
}
