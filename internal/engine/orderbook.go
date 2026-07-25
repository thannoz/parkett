package engine

// Order is a limit order resting in or entering the book.
// Prices are integer cents to avoid floating-point rounding on money.
type Order struct {
	ID     uint64
	Side   Side
	Price  int64 // in cents
	Qty    int64 // remaining quantity — shrinks with partial fills
	Seq    uint64
	Symbol string
}

// Trade is the result of two orders matching. Price is always the resting
// (maker) order's price.
type Trade struct {
	TakerOrderID uint64
	MakerOrderID uint64
	Price        int64
	Qty          int64
}

// Book is a limit order book for a single symbol. Orders are matched by
// price-time priority: best price first, oldest order first within a price.
type Book struct {
	bid    []*level
	ask    []*level
	seq    uint64
	symbol string
	orders map[uint64]*Order
}

type level struct {
	Price int64
	Order []*Order
}

func NewBook(symbol string) *Book {
	orders := make(map[uint64]*Order)
	return &Book{
		symbol: symbol,
		orders: orders,
	}
}

// AddLimit inserts a limit order into the book, matching it against the
// opposite side first. It returns the trades that were executed (empty if
// the order did not cross) — any remainder rests in the book.
func (b *Book) AddLimit(odr Order) []Trade {
	b.seq++
	odr.Seq = b.seq
	b.orders[odr.ID] = &odr

	switch odr.Side {
	case Buy:
		b.bid = b.insertLevel(b.bid, &odr, false)
	case Sell:
		b.ask = b.insertLevel(b.ask, &odr, true)
	}

	return nil
}

// insertLevel places odr into levels: if a level with odr.Price already
// exists, the order is appended to that level's FIFO queue; otherwise a new
// level is created and appended. Returns the (possibly grown) slice.
func (b *Book) insertLevel(levels []*level, odr *Order, ascending bool) []*level {
	found := false
	for _, l := range levels {
		if l.Price == odr.Price {
			l.Order = append(l.Order, odr)
			found = true
			break
		}
	}
	if !found {
		lvl := &level{
			Order: []*Order{odr},
			Price: odr.Price,
		}
		idx := len(levels)
		for i, l := range levels {
			belongBefore := false
			if ascending {
				belongBefore = l.Price > odr.Price
			} else {
				belongBefore = l.Price < odr.Price
			}
			if belongBefore {
				idx = i
				break
			}
		}
		levels = append(levels, nil)
		copy(levels[idx+1:], levels[idx:])
		levels[idx] = lvl
	}
	return levels
}

// BestBid returns the highest buy price and true, or 0 and false if empty.
func (b *Book) BestBid() (price int64, ok bool) {
	if len(b.bid) == 0 {
		return 0, false
	}
	return b.bid[0].Price, true

}

// BestAsk returns the lowest sell price and true, or 0 and false if empty.
func (b *Book) BestAsk() (price int64, ok bool) {
	if len(b.ask) == 0 {
		return 0, false
	}
	return b.ask[0].Price, true
}

// Cancel removes a resting order from the book. It returns false if the
// order is unknown (already filled, canceled, or never existed).
func (b *Book) Cancel(orderID uint64) bool {
	panic("not implemented")
}
