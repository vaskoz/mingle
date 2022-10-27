package main

import (
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
	return nil
}
