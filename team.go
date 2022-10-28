package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
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

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for _, team := range teams {
		r.Shuffle(len(team.Mates), func(i, j int) {
			team.Mates[i], team.Mates[j] = team.Mates[j], team.Mates[i]
		})
	}

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
						mingles[i].People = append(mingles[i].People, fmt.Sprintf("%s from %s", person.PersonName, person.TeamName))
						if total == gs {
							break outer
						}
						break
					}
				}
			}
		}
	}

	return mingles
}
