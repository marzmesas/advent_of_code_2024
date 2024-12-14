package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/dominikbraun/graph"
)

type order struct {
	before int
	after  int
}

type update struct {
	pages  []int
	graph  graph.Graph[int, int]
	sorted []int
}

func main() {
	orders, updates := readInput("day_5/input.txt")

	for i := range updates {
		updates[i].graph = createGraph(orders, updates[i])
		updates[i].sorted, _ = graph.TopologicalSort(updates[i].graph)
	}

	part1 := solvePart1(updates)
	fmt.Println("Part 1:", part1)

	part2 := solvePart2(updates)
	fmt.Println("Part 2:", part2)
}

func solvePart1(updates []update) int {
	sum := 0
	for _, u := range updates {
		if u.isValidUpdate() {
			sum += u.middlePage()
		}
	}
	return sum
}

func solvePart2(updates []update) int {
	sum := 0
	for _, u := range updates {
		if !u.isValidUpdate() {
			sum += u.middlePage()
		}
	}
	return sum
}

func (u *update) isValidUpdate() bool {
	return slices.Equal(u.pages, u.sorted)
}

func (u *update) middlePage() int {
	if len(u.sorted) == 0 {
		return 0
	}
	return u.sorted[len(u.sorted)/2]
}

func createGraph(orders []order, update update) graph.Graph[int, int] {
	g := graph.New(graph.IntHash, graph.Directed(), graph.Acyclic(), graph.PreventCycles())

	for _, o := range orders {
		if slices.Contains(update.pages, o.before) && slices.Contains(update.pages, o.after) {
			_ = g.AddVertex(o.before)
			_ = g.AddVertex(o.after)
			_ = g.AddEdge(o.before, o.after)
		}
	}

	return g
}

func readInput(filepath string) ([]order, []update) {
	data, _ := os.ReadFile(filepath)
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	orders := []order{}
	updates := []update{}
	parsingOrders := true

	for _, line := range lines {
		if line == "" {
			parsingOrders = false
			continue
		}

		if parsingOrders {
			parts := strings.Split(line, "|")
			before, _ := strconv.Atoi(parts[0])
			after, _ := strconv.Atoi(parts[1])
			orders = append(orders, order{before, after})
		} else {
			pageStrings := strings.Split(line, ",")
			pages := []int{}
			for _, p := range pageStrings {
				page, _ := strconv.Atoi(p)
				pages = append(pages, page)
			}
			updates = append(updates, update{pages: pages})
		}
	}

	return orders, updates
}
