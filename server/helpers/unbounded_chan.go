package helpers

// Returns an "unbounded" channel, which is really an 'in' and 'out' channel with an infinite
// queue in between.  To close the channel, close the 'in' channel.
func NewUnboundedChan[T any](bfr int) (ch_in chan<- T, ch_out <-chan T) {
	// Based on https://medium.com/capital-one-tech/building-an-unbounded-channel-in-go-789e175cd2cd
	_in, _out := make(chan T, bfr/2), make(chan T, bfr/2)

	go func() {
		defer close(_out)
		mid := make([]T, 16)
		val_or_nil := func() T {
			if len(mid) == 0 {
				return *new(T)
			} else {
				return mid[0]
			}
		}
		out_or_nil := func() chan T {
			if len(mid) == 0 {
				return nil
			} else {
				return _out
			}
		}
		for len(mid) > 0 || _in != nil {
			select {
			case item, ok := <-_in:
				if ok {
					mid = append(mid, item)
				} else {
					_in = nil
				}
			case out_or_nil() <- val_or_nil():
				// This case may have a nil "out" channel for this turn through the loop, which
				// will not allow the case to execute, if mid is 0. That's what "out_or_nil" does,
				// it chooses between the out channel or a nil channel based on whether the 'mid'
				// queue is empty.  Likewise, val_or_nil will either be a zero value or the first
				// item on the queue, depending on if there is an item there or not.
				mid = mid[1:]
			}
		}
	}()
	return _in, _out
}
