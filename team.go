package main

import (
	"log"
	"sort"
	"strings"
)

type Teammate struct {
	PersonName string
	TeamName   string
}

type Team struct {
	Mates []Teammate
}

func ExtractTeam(name, file string) Team {
	lines := strings.Split(strings.TrimSpace(file), "\n")
	team := Team{make([]Teammate, 0, len(lines))}
	for _, person := range lines {
		tMate := Teammate{person, name}
		team.Mates = append(team.Mates, tMate)
	}
	return team
}

func MingleTeams(teams []Team, groupSize []int) []Mingle {
	seated := make(map[string]struct{})
	mingles := make([]Mingle, len(groupSize))

	sort.Slice(teams, func(i, j int) bool { // largest teams first
		return len(teams[i].Mates) > len(teams[j].Mates)
	})

	for i, gs := range groupSize {
		mingles[i].MaxSize = groupSize[i]

		total := 0
	outer:
		for total != gs {
			for _, team := range teams {
				for _, person := range team.Mates {
					if _, done := seated[person.PersonName]; !done {
						seated[person.PersonName] = struct{}{}
						total++
						mingles[i].People = append(mingles[i].People, person.PersonName)
						if total == gs {
							break outer
						}
					}
				}
			}
		}
	}

	log.Println("seated size:", len(seated))
	return mingles
}
