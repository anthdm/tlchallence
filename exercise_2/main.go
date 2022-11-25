package main

import (
	"fmt"
	"untitled/orderbook"
)

// var asks = _data.LoadOrdersFromFile("./_data/asks.gob")
// var bids = _data.LoadOrdersFromFile("./_data/bids.gob")

const (
	MaxFactor = 1e8
)

func main() {
	asksSrc := "../data/asks.gob"
	bidSrc := "../data/bids.gob"
	book := orderbook.NewFromFile(asksSrc, bidSrc)

	book.Aggregate(100)

	for _, ask := range book.Asks {
		fmt.Println(ask)
	}

	// orders := make(Orders, len(asks))
	// it := 0
	// for price, size := range asks {
	// 	orders[it] = &Order{
	// 		price: price,
	// 		size:  size,
	// 	}
	// 	it++
	// }

	// sort.Sort(ByBestAsk{orders})

	// data := make(map[int]*Order)
	// for _, order := range orders {
	// 	r := RR(order.price)
	// 	a, ok := data[r]
	// 	if !ok {
	// 		data[r] = order
	// 	} else {
	// 		a.size += order.size
	// 	}
	// }

	// aggOrders := make(Orders, len(data))
	// i := 0
	// for price, order := range data {
	// 	aggOrders[i] = &Order{
	// 		price: float64(price),
	// 		size:  order.size,
	// 	}
	// 	i++
	// }

	// sort.Sort(ByBestAsk{aggOrders})

	// for _, order := range aggOrders {
	// 	fmt.Println(order)
	// }

	// The two sides of the book are IN of sync
	// Find a fast way to aggregate the orders on both sides of the book in price increments of 100.
	// ...16400, 16500, 16600, 16700, 16800, 16900, 17000...
	// Do not use the original maps, convert them to a different data structure.

	// Note: Some loss of precision is acceptable as long as its <= MaxFactor
	// Note: Initial data structure creation is not part of the benchmark. Only the aggregation operation is.
}
