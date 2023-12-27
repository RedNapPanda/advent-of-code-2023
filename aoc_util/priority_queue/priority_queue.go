package priority_queue

import (
    "cmp"
    "sort"
)

/*
PriorityQueue
Keys slice is priority ordered on Push
Store data in a map
Must implement Weight, so we can do a valid comparison
*/
type PriorityQueue[K comparable, W cmp.Ordered] struct {
    Keys  []K
    Nodes map[K]W
}

// Len is part of sort.Interface
func (q *PriorityQueue[K, W]) Len() int {
    return len(q.Keys)
}

// Swap is part of sort.Interface
func (q *PriorityQueue[K, W]) Swap(x, y int) {
    q.Keys[x], q.Keys[y] = q.Keys[y], q.Keys[x]
}

// Less is part of sort.Interface
func (q *PriorityQueue[K, W]) Less(x, y int) bool {
    return q.Nodes[q.Keys[x]] < q.Nodes[q.Keys[y]]
}

// Push updates or inserts a new key in the priority priority_queue
func (q *PriorityQueue[K, W]) Push(key K, data W) {
    if _, ok := q.Nodes[key]; !ok {
        q.Keys = append(q.Keys, key)
    }
    q.Nodes[key] = data
    sort.Sort(q)
}

// Pop returns the first (key, priority)
func (q *PriorityQueue[K, W]) Pop() (K, W) {
    if len(q.Keys) == 0 {
        return *new(K), *new(W)
    }
    var key K
    var weight W
    var ok bool
    q.Keys, key = q.Keys[1:], q.Keys[0]
    if weight, ok = q.Nodes[key]; ok {
        delete(q.Nodes, key)
    }

    return key, weight
}

// Get returns the cost of key
func (q *PriorityQueue[K, W]) Get(key K) (weight W, ok bool) {
    weight, ok = q.Nodes[key]
    return
}

// NewPriorityQueue creates a new instance
func NewPriorityQueue[K comparable, W cmp.Ordered]() *PriorityQueue[K, W] {
    var q PriorityQueue[K, W]
    q.Nodes = make(map[K]W)
    return &q
}
