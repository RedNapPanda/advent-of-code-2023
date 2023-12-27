package astar

import (
	"aoc/aoc_util/priority_queue"
	"container/heap"
)

type VisitedNode[T comparable] struct {
	Parent T
	Score  int
}

type scoredNode[T comparable] struct {
	Node      T
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
func AStar[T comparable](src T, destinationFn func(*T) bool, neighborsFn func(*T) []T, scoreFn func(*T) int, heuristicFn func(*T) int) (int, VisitedNode[T], map[T]VisitedNode[T]) {
	var destination scoredNode[T]
	visitedMap := make(map[T]VisitedNode[T])
	var pq priority_queue.PriorityQueue[scoredNode[T]]
	heap.Init(&pq)

	srcHeuristic := heuristicFn(&src)
	item := buildItem(scoredNode[T]{src, 0, srcHeuristic}, srcHeuristic)

	heap.Push(&pq, item)
	visitedMap[src] = VisitedNode[T]{src, 0}

	var current scoredNode[T]
	var score int
	for pq.Len() > 0 {
		item = heap.Pop(&pq).(*priority_queue.Item[scoredNode[T]])
		current, score = item.Data, item.Priority
		if destinationFn(&current.Node) {
			destination = current
			break
		}

		for _, neighbor := range neighborsFn(&current.Node) {
			visitedMap[current.Node] = VisitedNode[T]{current.Node, score}
			if _, ok := visitedMap[neighbor]; ok {
				continue
			}
			newScore := score + scoreFn(&neighbor)
			visitedNeighbor, wasVisited := visitedMap[neighbor]
			if !wasVisited || newScore < visitedNeighbor.Score {
				newItem := buildItem(scoredNode[T]{neighbor, newScore, heuristicFn(&neighbor)}, newScore+heuristicFn(&neighbor))
				heap.Push(&pq, newItem)
				visitedMap[neighbor] = VisitedNode[T]{current.Node, newScore}
			}
		}
	}

	return visitedMap[destination.Node].Score, visitedMap[destination.Node], visitedMap
}

func buildItem[T comparable](node scoredNode[T], weight int) *priority_queue.Item[scoredNode[T]] {
	item := priority_queue.Item[scoredNode[T]]{Data: node, Priority: weight}
	return &item
}
