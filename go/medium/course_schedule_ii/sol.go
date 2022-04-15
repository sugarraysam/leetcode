package course_schedule_ii

func findOrder(numCourses int, prerequisites [][]int) []int {
	sched := NewSchedule(numCourses, prerequisites)
	return sched.solve()
}

type State string

const (
	VISITED   State = "VISITED"
	VISITING  State = "VISITING"
	UNVISITED State = "UNVISITED"
)

type Schedule struct {
	res        []int
	hasCycle   bool
	prereq     map[int][]int
	numCourses int
	state      map[int]State
}

func NewSchedule(numCourses int, prerequisites [][]int) *Schedule {
	prereq := make(map[int][]int)
	state := make(map[int]State)

	// init map with empty slice for every course
	for c := 0; c < numCourses; c++ {
		prereq[c] = []int{}
		state[c] = UNVISITED
	}
	// populate prereq map with prerequisites
	for _, p := range prerequisites {
		prereq[p[0]] = append(prereq[p[0]], p[1])
	}

	return &Schedule{
		res:        make([]int, 0),
		prereq:     prereq,
		numCourses: numCourses,
		state:      state,
	}
}

func (s *Schedule) solve() []int {
	// call dfs on every course to populate my schedule
	// need to loop through every course because might have a disconnected
	// graph
	for c := 0; c < s.numCourses; c++ {
		if s.state[c] == UNVISITED {
			s.dfs(c)
		}

		// cycle was detected
		if s.hasCycle {
			return []int{}
		}
	}
	return s.res
}

func (s *Schedule) dfs(c int) {
	// base case - detect cycle, if curr node is already in VISITING mode
	if s.hasCycle || s.state[c] == VISITING {
		s.hasCycle = true
		return
	}

	// base case - visited
	if s.state[c] == VISITED {
		return
	}

	s.state[c] = VISITING
	for _, prereq := range s.prereq[c] {
		// want to follow both UNVISITED + VISITING
		if s.state[prereq] != VISITED {
			s.dfs(prereq)
		}
	}

	// visisted all prereqs successfully, can append to my schedule
	if !s.hasCycle {
		s.state[c] = VISITED
		s.res = append(s.res, c)
	}
}
