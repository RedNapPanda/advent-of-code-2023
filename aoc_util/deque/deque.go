package deque

// Deque need to write some unit tests for this
type Deque[T any] struct {
	buffer            []T
	head, tail, count int
}

func NewDeque[T any](initCap int) *Deque[T] {
	d := Deque[T]{}
	if initCap > 0 {
		// find the next power of 2 higher than the requested capacity
		if initCap%2 != 0 {
			initCap = initCap << 1
		}
		d.buffer = make([]T, initCap)
	} else {
		// always make a buffer of size 2, that was 2 things at minimum can be push/popped
		d.buffer = make([]T, 2)
	}
	return &d
}

func (d *Deque[T]) Cap() int {
	if d == nil {
		return 0
	}
	return len(d.buffer)
}

func (d *Deque[T]) Len() int {
	if d == nil {
		return 0
	}
	return d.count
}

/*
PushFront
[_ _ _ _] 0,0
[A _ _ _] 3,1
[A _ _ B] 2,1
[A _ _ _] 3,1
*/
func (d *Deque[T]) PushFront(t T) {
	if d == nil {
		return
	}
	if d.count == len(d.buffer) {
		d.expand()
	}
	d.buffer[d.head] = t
	d.shiftHead(true)
	d.count++
}

/*
PushBack
[_ _ _ _] 0,0
[A _ _ _] 3,1
[A B _ _] 3,2
[A _ _ _] 3,1
*/
func (d *Deque[T]) PushBack(t T) {
	if d == nil {
		return
	}
	if d.count == len(d.buffer) {
		d.expand()
	}
	d.buffer[d.tail] = t
	d.shiftTail(false)
	d.count++
}

func (d *Deque[T]) PopFront() T {
	if d == nil {
		return *new(T)
	}
	v := d.buffer[d.head]
	d.shiftHead(false)
	d.buffer[d.head] = *new(T)
	d.count--
	return v
}

func (d *Deque[T]) PopBack() T {
	if d == nil {
		return *new(T)
	}
	v := d.buffer[d.tail]
	d.shiftTail(true)
	d.buffer[d.tail] = *new(T)
	d.count--
	return v
}

func (d *Deque[T]) shiftHead(neg bool) {
	if neg {
		if d.head == d.tail {
			d.tail = (d.tail + 1) % len(d.buffer)
		}
		d.head--
		if d.head < 0 {
			d.head = len(d.buffer) - 1
		}
	} else {
		d.head = (d.head + 1) % len(d.buffer)
	}
}

func (d *Deque[T]) shiftTail(neg bool) {
	if neg {
		if d.tail == d.head {
			d.head = (d.head + 1) % len(d.buffer)
		}
		d.tail--
		if d.tail < 0 {
			d.tail = len(d.buffer) - 1
		}
	} else {
		d.tail = (d.tail + 1) % len(d.buffer)
	}
}

func (d *Deque[T]) expand() {
	newBuffer := make([]T, d.count<<1)
	if d.tail > d.head {
		copy(newBuffer, d.buffer[d.head:d.tail])
	} else {
		n := copy(newBuffer, d.buffer[d.head:])
		copy(newBuffer[n:], d.buffer[:d.tail])
	}
	d.head = 0
	d.tail = d.count
	d.buffer = newBuffer
}
