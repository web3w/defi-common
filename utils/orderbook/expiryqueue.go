package orderbook

import "container/heap"

type ExpiryQueue struct {
	hp *expiryHeap
}

func NewExpiryQueue() *ExpiryQueue {
	return &ExpiryQueue{hp: newExpiryHeap()}
}

type orderExpiry struct {
	orderId       int64
	expireTimeSec int64
}

func (eq *ExpiryQueue) Len() int {
	return eq.hp.Len()
}

// Must check Len() != 0 before calling this.
func (eq *ExpiryQueue) Top() orderExpiry {
	return eq.hp.slice[0]
}

// Precondition: `orderId` must be NOT in the queue yet.
func (eq *ExpiryQueue) Add(orderId int64, expireTimeSec int64) {
	item := orderExpiry{
		expireTimeSec: expireTimeSec,
		orderId:       orderId,
	}
	heap.Push(eq.hp, item)
}

// Returns false if not found.
func (eq *ExpiryQueue) Remove(orderId int64) bool {
	index, ok := eq.hp.mapIndexByOrderId[orderId]
	if !ok {
		return false
	}
	heap.Remove(eq.hp, index)
	return true
}

//--------------------- Internal type `expiryHeap` ------------------------------

// Implements interface `heap.Interface`
type expiryHeap struct {
	slice             []orderExpiry
	mapIndexByOrderId map[int64]int
}

func newExpiryHeap() *expiryHeap {
	return &expiryHeap{
		slice:             nil,
		mapIndexByOrderId: make(map[int64]int),
	}
}

func (hp *expiryHeap) Len() int {
	return len(hp.slice)
}

func (hp *expiryHeap) Less(i, j int) bool {
	return hp.slice[i].expireTimeSec < hp.slice[j].expireTimeSec
}

func (hp *expiryHeap) Swap(i, j int) {
	hp.slice[i], hp.slice[j] = hp.slice[j], hp.slice[i]
	hp.mapIndexByOrderId[hp.slice[i].orderId] = i
	hp.mapIndexByOrderId[hp.slice[j].orderId] = j
}

func (hp *expiryHeap) Push(x interface{}) {
	item := x.(orderExpiry)
	hp.mapIndexByOrderId[item.orderId] = len(hp.slice)
	hp.slice = append(hp.slice, item)
}

func (hp *expiryHeap) Pop() interface{} {
	n := len(hp.slice)
	item := hp.slice[n-1]
	hp.slice = hp.slice[0 : n-1]
	delete(hp.mapIndexByOrderId, item.orderId)
	return item
}
