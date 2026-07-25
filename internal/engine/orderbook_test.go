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

func TestAddLimit_SortedInsert_Bids(t *testing.T) {
	b := NewBook("XOM")
	orders := []Order{
		{ID: 1, Side: Buy, Price: 10100, Qty: 5},
		{ID: 2, Side: Buy, Price: 10050, Qty: 5},
		{ID: 3, Side: Buy, Price: 10000, Qty: 5},
	}

	for _, odr := range orders {
		b.AddLimit(odr)
	}

	gotBidLen := len(b.bid)
	if gotBidLen != 3 {
		t.Fatalf("len(b.bid) = %d, want %d", gotBidLen, 3)
	}

	wantPrices := []int64{10100, 10050, 10000}
	for i, want := range wantPrices {
		if got := b.bid[i].Price; got != want {
			t.Errorf("b.bid[%d].Price = %d, want %d", i, got, want)
		}
	}
}

func TestAddLimit_SortedInsert_Asks(t *testing.T) {
	b := NewBook("XOM")
	orders := []Order{
		{ID: 1, Side: Sell, Price: 100, Qty: 5},
		{ID: 2, Side: Sell, Price: 110, Qty: 5},
		{ID: 3, Side: Sell, Price: 105, Qty: 5},
	}

	for _, odr := range orders {
		b.AddLimit(odr)
	}

	gotAskLen := len(b.ask)
	if gotAskLen != 3 {
		t.Fatalf("len(b.ask) = %d, want %d", gotAskLen, 3)
	}

	wantPrices := []int64{100, 105, 110}
	for i, want := range wantPrices {
		if got := b.ask[i].Price; got != want {
			t.Errorf("b.ask[%d].Price = %d, want %d", i, got, want)
		}
	}
}

func TestBestBid_Empty(t *testing.T) {
	b := NewBook("XOM")

	gotPrice, gotOk := b.BestBid()
	if gotOk != false {
		t.Errorf("ok = %v, want false", gotOk)
	}

	if gotPrice != 0 {
		t.Fatalf("price = %d, want %d", gotPrice, 0)
	}
}

func TestBestBid_NotEmpty(t *testing.T) {
	b := NewBook("APPL")

	orders := []Order{
		{Side: Buy, Price: 10100},
	}
	for _, odr := range orders {
		b.AddLimit(odr)
	}

	gotPrice, gotOk := b.BestBid()
	if gotOk != true {
		t.Errorf("ok = %v, want true", gotOk)
	}

	if gotPrice != 10100 {
		t.Errorf("price %d, want %d", gotPrice, 10100)
	}
}

func TestBestBid_MoreBids(t *testing.T) {
	b := NewBook("APPL")

	orders := []Order{
		{ID: 1, Side: Buy, Price: 10100},
		{ID: 2, Side: Buy, Price: 20100},
	}
	for _, odr := range orders {
		b.AddLimit(odr)
	}

	gotBestBid, gotOk := b.BestBid()
	if gotOk != true {
		t.Errorf("ok = %v, want true", gotOk)
	}
	if gotBestBid != 20100 {
		t.Errorf("best bid = %d, want %d", gotBestBid, 20100)
	}
}

func TestBestAsk_MoreAsks(t *testing.T) {
	b := NewBook("APPL")

	orders := []Order{
		{ID: 1, Side: Sell, Price: 120},
		{ID: 2, Side: Sell, Price: 105},
	}
	for _, odr := range orders {
		b.AddLimit(odr)
	}

	gotBestAsk, gotOk := b.BestAsk()
	if gotOk != true {
		t.Errorf("ok = %v, want false", gotOk)
	}

	if gotBestAsk != 105 {
		t.Errorf("best ask = %d, want %d", gotBestAsk, 105)
	}
}
