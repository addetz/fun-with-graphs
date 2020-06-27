package courseschedule

type graph struct {
	nodes map[int][]int
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	g := &graph{
		nodes: make(map[int][]int, 0),
	}
	for i := 0; i < numCourses; i++ {
		g.nodes[i] = make([]int, 0)
	}
	for _, v := range prerequisites {
		g.nodes[v[0]] = append(g.nodes[v[0]], v[1])
	}
	return !contains_cycle(g, numCourses)
}

func contains_cycle(g *graph, n int) bool {
	visited := make([]int, n)
	for i := 0; i < n; i++ {
		if hasCycle(g, i, visited) {
			return true
		}
	}
	return false
}

func hasCycle(g *graph, i int, visited []int) bool {
	visited[i] = 1
	for _, u := range g.nodes[i] {
		if visited[u] == 0 {
			if hasCycle(g, u, visited) {
				return true
			}
		}
		if visited[u] == 1 {
			return true
		}
	}
	visited[i] = 2
	return false
}
