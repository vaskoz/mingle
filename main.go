package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	stderr = os.Stderr
	stdout = os.Stdout
	exit   = os.Exit
	args   = os.Args
)

func main() {
	peopleDir := os.Getenv("MINGLE_TEAM_DIR")
	if peopleDir == "" {
		fmt.Fprintln(stderr, "must set env var MINGLE_TEAM_DIR")
		exit(1)
	}

	groupSizes := make([]int, 0, len(args))

	for i := 1; i < len(args); i++ {
		if size, err := strconv.Atoi(args[i]); err != nil {
			fmt.Fprintln(stderr, "argument is not a integer for group size", args[i])
			exit(1)
		} else {
			groupSizes = append(groupSizes, size)
		}
	}

	files, err := os.ReadDir(peopleDir)
	if err != nil {
		fmt.Fprintln(stderr, "failed to read directory given")
		exit(1)
	}

	teams := make([]Team, 0, len(files))

	for _, file := range files {
		fullPath := fmt.Sprintf("%s/%s", peopleDir, file.Name())
		b, _ := os.ReadFile(fullPath)
		fileS := string(b)
		team := ExtractTeam(file.Name(), fileS)
		teams = append(teams, team)
	}

	totalPeople := 0

	for _, team := range teams {
		totalPeople += len(team.Mates)
	}

	log.Println("Total People", totalPeople)

	totalSeating := 0
	for _, size := range groupSizes {
		totalSeating += size
	}

	if totalPeople > totalSeating {
		fmt.Fprintln(stderr, "not enough seats for the people", totalPeople, totalSeating)
		exit(1)
	}

	seating := MingleTeams(teams, groupSizes)

	fmt.Fprintln(stdout, seating)
}
