package main

import (
	"fmt"
	"math/rand"
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

func MingleTeams(teams []Team, groupSize []int, matches map[string]map[string]struct{}) []Mingle {
	seated := make(map[string]struct{})
	mingles := make([]Mingle, len(groupSize))

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	r.Shuffle(len(teams), func(i, j int) {
		teams[i], teams[j] = teams[j], teams[i]
	})

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
						// alreadySeated := len(mingles[i].People)
						key := fmt.Sprintf("%s from %s", person.PersonName, person.TeamName)
						count := 0
						for _, k := range mingles[i].People {
							if _, found := matches[key][k]; found {
								count++
							}
						}
						if count > gs/2 {
							continue
						}
						seated[person.PersonName] = struct{}{}
						total++
						mingles[i].People = append(mingles[i].People, key)
						if total == gs {
							break outer
						}
						break
					}
				}
			}
		}
	}

	for _, mingle := range mingles {
		for i := 0; i < len(mingle.People)-1; i++ {
			for j := i + 1; j < len(mingle.People); j++ {
				first := mingle.People[i]
				second := mingle.People[j]

				if matches[first] == nil {
					matches[first] = make(map[string]struct{})
				}
				if matches[second] == nil {
					matches[second] = make(map[string]struct{})
				}
				matches[first][second] = struct{}{}
				matches[second][first] = struct{}{}
			}
		}
	}

	return mingles
}
