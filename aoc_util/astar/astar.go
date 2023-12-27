package astar

import (
    "aoc/aoc_util/priority_queue"
)

type VisitedNode[T comparable] struct {
    Parent T
    Score  int
}

type ScoredNode[T comparable] struct {
    Node      T
    Score     int
    Heuristic int
}

/*
AStar A* pathfinding from src to dst
This will implicitly

neighbors func - calculate next possible locations from current
Score func -
Heuristic func - calculate g(n T) for n to dst diff for [current, neighbor]

	think Manhattan dist, etc. to ensure path always moves in direction (src,dist)

Need to carry forward the Score sum when pushing neighbors to queue
*/
func AStar[T comparable](src T, maxNodes int, destinationFn func(*T) bool, neighborsFn func(*T) []T, scoreFn func(*T) int, heuristicFn func(*T) int) (int, VisitedNode[T], map[T]VisitedNode[T]) {
    var destination ScoredNode[T]
    visitedMap := make(map[T]VisitedNode[T], maxNodes)
    priorityQueue := priority_queue.NewPriorityQueue[ScoredNode[T], int]()

    srcHeuristic := heuristicFn(&src)
    priorityQueue.Push(ScoredNode[T]{src, 0, srcHeuristic}, srcHeuristic)
    visitedMap[src] = VisitedNode[T]{src, 0}

    var current ScoredNode[T]
    var score int
    for priorityQueue.Len() > 0 {
        current, score = priorityQueue.Pop()
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
                priorityQueue.Push(ScoredNode[T]{neighbor, newScore, heuristicFn(&neighbor)}, newScore+heuristicFn(&neighbor))
                visitedMap[neighbor] = VisitedNode[T]{current.Node, newScore}
            }
        }
    }

    return visitedMap[destination.Node].Score, visitedMap[destination.Node], visitedMap
}
