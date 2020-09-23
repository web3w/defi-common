package orderbook

import (
	"testing"

	"github.com/gisvr/defi-common/utils/utest"
)

type ExpiryQueueTest struct {
	utest.RequireSuite
}

// The hook of `go test`
func TestRun_ExpiryQueueTest(t *testing.T) {
	utest.Run(t, &ExpiryQueueTest{})
}

func (t *ExpiryQueueTest) TestExpiryQueue_ManyRounds() {
	que := NewExpiryQueue()
	for i := 0; i < 20; i++ {
		t.runSingleRound(que)
	}
}

// Takes an empty queue. Plays with it and then remove all the elements from it.
func (t *ExpiryQueueTest) runSingleRound(que *ExpiryQueue) {
	t.checkLen(que, 0)

	que.Add(100, 2000)
	t.Equal(orderExpiry{100, 2000}, que.Top())
	t.checkLen(que, 1)

	que.Add(101, 1900)
	que.Add(102, 2100)
	que.Add(103, 2200)
	que.Add(104, 1500)
	que.Add(105, 2000)
	que.Add(106, 2300)
	que.Add(107, 2150)

	t.Equal(orderExpiry{104, 1500}, que.Top())
	t.checkLen(que, 8)

	que.Add(108, 1500)
	t.Equal(1500, int(que.Top().expireTimeSec))
	t.checkLen(que, 9)

	t.True(que.Remove(que.Top().orderId))
	t.Equal(1500, int(que.Top().expireTimeSec))
	t.checkLen(que, 8)

	t.True(que.Remove(que.Top().orderId))
	t.Equal(orderExpiry{101, 1900}, que.Top())
	t.checkLen(que, 7)

	t.True(que.Remove(106)) // remove the last
	t.Equal(orderExpiry{101, 1900}, que.Top())
	t.checkLen(que, 6)

	t.True(que.Remove(105)) // remove a middle
	t.Equal(orderExpiry{101, 1900}, que.Top())
	t.checkLen(que, 5)

	// Left: 100, 101, 102, 103, 107

	// Removing deleted returns false
	t.False(que.Remove(104))
	t.False(que.Remove(105))
	t.False(que.Remove(106))
	t.False(que.Remove(108))

	// Remove from the top one by one.
	t.Equal(orderExpiry{101, 1900}, que.Top())
	t.checkLen(que, 5)
	t.True(que.Remove(101))

	t.Equal(orderExpiry{100, 2000}, que.Top())
	t.checkLen(que, 4)
	t.True(que.Remove(100))

	t.Equal(orderExpiry{102, 2100}, que.Top())
	t.checkLen(que, 3)
	t.True(que.Remove(102))

	t.Equal(orderExpiry{107, 2150}, que.Top())
	t.checkLen(que, 2)
	t.True(que.Remove(107))

	t.Equal(orderExpiry{103, 2200}, que.Top())
	t.checkLen(que, 1)
	t.True(que.Remove(103))

	t.checkLen(que, 0)
}

func (t *ExpiryQueueTest) checkLen(que *ExpiryQueue, expected int) {
	t.Equal(expected, que.Len())
	t.Equal(expected, len(que.hp.mapIndexByOrderId))
	t.Equal(expected, len(que.hp.slice))
}
