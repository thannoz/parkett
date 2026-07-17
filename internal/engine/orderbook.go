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

// NewBook creates an empty order book.
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
func (b *Book) AddLimit(o Order) []Trade {
	panic("not implemented")
}

// BestBid returns the highest buy price and true, or 0 and false if empty.
func (b *Book) BestBid() (price int64, ok bool) {
	panic("not implemented")
}

// BestAsk returns the lowest sell price and true, or 0 and false if empty.
func (b *Book) BestAsk() (price int64, ok bool) {
	panic("not implemented")
}

// Cancel removes a resting order from the book. It returns false if the
// order is unknown (already filled, cancelled, or never existed).
func (b *Book) Cancel(orderID uint64) bool {
	panic("not implemented")
}
