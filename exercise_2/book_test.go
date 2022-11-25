package main

import (
	"testing"
	"untitled/orderbook"

	"github.com/stretchr/testify/assert"
)

func Benchmark_Group(b *testing.B) {
	b.Run("group", func(b *testing.B) {
		var (
			asksSrc = "../data/asks.gob"
			bidSrc  = "../data/bids.gob"
			book    = orderbook.NewFromFile(asksSrc, bidSrc)
		)
		// create book that will hold aggregated data
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			book.Aggregate(100)
			// aggregate all orders in this loop, not just one.
		}
	})
}

func Test_Group(t *testing.T) {
	var (
		asksSrc = "../data/asks.gob"
		bidSrc  = "../data/bids.gob"
		book    = orderbook.NewFromFile(asksSrc, bidSrc)
	)
	book.Aggregate(100)

	// Make a quick map here to test effeciently here
	orders := make(map[int]*orderbook.Order, len(book.Asks))
	for i := 0; i < len(book.Asks); i++ {
		orders[int(book.Asks[i].Price)] = book.Asks[i]
	}

	assert.Equal(t, 1010.518, orders[23000].Size)
	assert.Equal(t, 1010.518, orders[23000].Size)
	assert.Equal(t, 1010.518, orders[23000].Size)
	assert.Equal(t, 11.049, orders[29500].Size)
	assert.Equal(t, 17.996, orders[27300].Size)
	assert.Equal(t, 306.706, orders[23800].Size)
	assert.Equal(t, 2216.883, orders[21600].Size)
	assert.Equal(t, 2392.173, orders[24100].Size)
	assert.Equal(t, 20.946, orders[17000].Size)
	assert.Equal(t, 17.222, orders[26800].Size)
	assert.Equal(t, 1022.681, orders[23300].Size)
	assert.Equal(t, 401.26, orders[22200].Size)
	assert.Equal(t, 640.809, orders[22400].Size)
	assert.Equal(t, 22.272, orders[27600].Size)
	assert.Equal(t, 981.298, orders[23200].Size)
	assert.Equal(t, 18.051, orders[28400].Size)
	assert.Equal(t, 54.861, orders[24900].Size)
	assert.Equal(t, 2744.176, orders[21400].Size)
	assert.Equal(t, 19.512, orders[28600].Size)
	assert.Equal(t, 109.883, orders[26300].Size)
	assert.Equal(t, 127.471, orders[24000].Size)
	assert.Equal(t, 10.186, orders[29000].Size)
	assert.Equal(t, 87.069, orders[25900].Size)
	assert.Equal(t, 2267.394, orders[21300].Size)
	assert.Equal(t, 16.348, orders[27400].Size)
	assert.Equal(t, 270.195, orders[25200].Size)
	assert.Equal(t, 1036.14, orders[22900].Size)
	assert.Equal(t, 442.207, orders[25600].Size)
	assert.Equal(t, 270.46, orders[24700].Size)
	assert.Equal(t, 668.869, orders[22500].Size)
	assert.Equal(t, 137.444, orders[16800].Size)
	assert.Equal(t, 1491.506, orders[21800].Size)
	assert.Equal(t, 748.323, orders[24300].Size)
	assert.Equal(t, 735.7, orders[22000].Size)
	assert.Equal(t, 98.03, orders[28100].Size)
	assert.Equal(t, 17.163, orders[27000].Size)
	assert.Equal(t, 13.391, orders[27200].Size)
	assert.Equal(t, 642.461, orders[22700].Size)
	assert.Equal(t, 1084.352, orders[25400].Size)
	assert.Equal(t, 10.761, orders[27900].Size)
	assert.Equal(t, 645.224, orders[24500].Size)
	assert.Equal(t, 532.86, orders[23600].Size)
	assert.Equal(t, 163.7, orders[29700].Size)
	assert.Equal(t, 541.354, orders[16600].Size)
	assert.Equal(t, 2242.342, orders[21500].Size)
	assert.Equal(t, 12.395, orders[28800].Size)
	assert.Equal(t, 118.17, orders[16900].Size)
	assert.Equal(t, 431.942, orders[24200].Size)
	assert.Equal(t, 24.826, orders[29600].Size)
	assert.Equal(t, 1107.865, orders[26100].Size)
	assert.Equal(t, 54.68, orders[25000].Size)
	assert.Equal(t, 18.292, orders[29900].Size)
	assert.Equal(t, 16.047, orders[27700].Size)
	assert.Equal(t, 19.454, orders[26900].Size)
}
