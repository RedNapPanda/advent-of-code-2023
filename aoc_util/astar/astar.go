package astar

import (
	"aoc/aoc_util/priority_queue"
	"container/heap"
)

type VisitedNode[T comparable] struct {
	Node      *T
	Parent    *T
	Score     int
	Heuristic int
}

/*
AStar A* pathfinding from src to dst
https://en.wikipedia.org/wiki/A*_search_algorithm#Pseudocode

neighbors func - calculate next possible locations from current
Score func -
Heuristic func - calculate g(n T) for n to dst diff for [current, neighbor]

	think Manhattan dist, etc. to ensure path always moves in direction (src,dist)

Need to carry forward the Score sum when pushing neighbors to queue
*/
func AStar[T comparable](src *T, destinationFn func(*T) bool, neighborsFn func(*T) []T, scoreFn func(*T) int, heuristicFn func(*T) int) (VisitedNode[T], []VisitedNode[T], map[T]VisitedNode[T]) {
	var dest VisitedNode[T]
	var pq priority_queue.PriorityQueue[VisitedNode[T]]
	visitedMap := make(map[T]VisitedNode[T])
	srcH := heuristicFn(src)
	srcVN := VisitedNode[T]{src, nil, 0, srcH}
	visitedMap[*src] = srcVN
	item := buildItem(srcVN, srcH)

	heap.Init(&pq)
	heap.Push(&pq, item)

	for pq.Len() > 0 {
		item = heap.Pop(&pq).(*priority_queue.Item[VisitedNode[T]])
		current, score := item.Data, item.Priority
		if destinationFn(current.Node) {
			dest = current
			break
		}

		neighbors := neighborsFn(current.Node)
		for i, _ := range neighbors {
			neighbor := neighbors[i]
			newScore := score + scoreFn(&neighbor)
			visitedNeighbor, visited := visitedMap[neighbor]
			if !visited || newScore < visitedNeighbor.Score {
				vn := VisitedNode[T]{&neighbor, current.Node, newScore, heuristicFn(&neighbor)}
				visitedMap[neighbor] = vn
				newItem := buildItem(vn, newScore+heuristicFn(&neighbor))
				heap.Push(&pq, newItem)
			}
		}
	}

	if dest.Node == nil || dest.Parent == nil {
		return dest, nil, nil
	}

	var l []VisitedNode[T]
	n := visitedMap[*dest.Node]
	for n.Parent != nil {
		l = append(l, n)
		n = visitedMap[*n.Parent]
	}

	return visitedMap[*dest.Node], l, visitedMap
}

func buildItem[T comparable](node VisitedNode[T], weight int) *priority_queue.Item[VisitedNode[T]] {
	item := priority_queue.Item[VisitedNode[T]]{Data: node, Priority: weight}
	return &item
}
