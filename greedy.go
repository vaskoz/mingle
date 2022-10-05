package main

func Greedy(person Person, maxSize int) []Mingle {
	var mingles []Mingle

	m := Mingle{nil, maxSize}
	for i, p := range person.Prefs {
		if i != 0 && i%(maxSize-1) == 0 {
			m.People = append(m.People, person.Name)
			mingles = append(mingles, m)
			m = Mingle{nil, maxSize}
		}
		m.People = append(m.People, p)
	}

	if len(m.People) != 0 {
		m.People = append(m.People, person.Name)
		mingles = append(mingles, m)
	}

	return mingles
}

func GreedyPeople(people []*Person, maxSize int) []Mingle {
	var (
		mingles []Mingle
		group   map[string]struct{} = make(map[string]struct{})
		pi      int
	)

	matches := make(map[string]map[string]bool)

	for len(people) != 0 {
		p := people[pi]
		p2 := p.Prefs[0]
		if !matches[p.Name][p2] {
			if m := matches[p.Name]; m == nil {
				matches[p.Name] = make(map[string]bool)
			}
			if m := matches[p2]; m == nil {
				matches[p2] = make(map[string]bool)
			}
			group[p.Name] = struct{}{}
			group[p2] = struct{}{}
		}

		p.Prefs = p.Prefs[1:]
		if len(p.Prefs) == 0 {
			people = append(people[:pi], people[pi+1:]...)
		}

		if len(group) >= maxSize {
			res := make([]string, 0, len(group))
			for k := range group {
				res = append(res, k)
			}
			mingles = append(mingles, Mingle{res, maxSize})
			group = make(map[string]struct{})
		}

		pi++
		if pi >= len(people) {
			pi = 0
		}
	}

	if len(group) != 0 {
		res := make([]string, 0, len(group))
		for k := range group {
			res = append(res, k)
		}
		mingles = append(mingles, Mingle{res, maxSize})
	}

	return mingles
}
