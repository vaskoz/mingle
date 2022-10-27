package main

import (
	"fmt"
	"os"
)

var (
	stderr = os.Stderr
	stdout = os.Stdout
	exit   = os.Exit
)

func main() {
	peopleDir := os.Getenv("MINGLE_TEAM_DIR")
	if peopleDir == "" {
		fmt.Fprintln(stderr, "must set env var MINGLE_TEAM_DIR")
		exit(1)
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

	seating := MingleTeams(teams)

	fmt.Fprintln(stdout, seating)
}
