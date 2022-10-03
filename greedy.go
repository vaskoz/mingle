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
