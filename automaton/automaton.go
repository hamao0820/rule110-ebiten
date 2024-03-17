package automaton

type Automaton struct {
	State    []uint
	n        int
	stateNum int
}

func NewAutomaton(initial []uint) *Automaton {
	return &Automaton{
		State: initial,
		n:     len(initial),
	}
}

func (a *Automaton) Update() int {
	next := make([]uint, len(a.State))
	for i := 0; i < a.n; i++ {
		if i == 0 {
			next[i] = rule110(0, a.State[i], a.State[i+1])
			continue
		}
		if i == a.n-1 {
			next[i] = rule110(a.State[i-1], a.State[i], 0)
			continue
		}
		next[i] = rule110(a.State[i-1], a.State[i], a.State[i+1])
	}
	a.State = next
	a.stateNum++
	return a.stateNum
}

func rule110(a, b, c uint) uint {
	if a == 1 && b == 1 && c == 1 {
		return 0
	}
	if a == 1 && b == 1 && c == 0 {
		return 1
	}
	if a == 1 && b == 0 && c == 1 {
		return 1
	}
	if a == 1 && b == 0 && c == 0 {
		return 0
	}
	if a == 0 && b == 1 && c == 1 {
		return 1
	}
	if a == 0 && b == 1 && c == 0 {
		return 1
	}
	if a == 0 && b == 0 && c == 1 {
		return 1
	}
	if a == 0 && b == 0 && c == 0 {
		return 0
	}
	return 0
}
