package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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

	matches := make(map[string]map[string]struct{})

	matchesFile := os.Getenv("MINGLE_MATCHES_FILE")
	if matchesFile != "" {
		matchesData, err := os.ReadFile(matchesFile)
		if err != nil {
			fmt.Fprintln(stderr, "failed to open matches file")
			exit(1)
		}
		matchesS := string(matchesData)
		matchesS = strings.TrimSpace(matchesS)
		if matchesS != "" {
			lines := strings.Split(matchesS, "\n")
			for _, line := range lines {
				parts := strings.Split(line, ",")
				if matches[parts[0]] == nil {
					matches[parts[0]] = make(map[string]struct{})
				}
				if matches[parts[1]] == nil {
					matches[parts[1]] = make(map[string]struct{})
				}
				matches[parts[0]][parts[1]] = struct{}{}
				matches[parts[1]][parts[0]] = struct{}{}
			}
		}
	}

	totalPeople := 0
	for _, team := range teams {
		totalPeople += len(team.Mates)
	}

	totalSeating := 0
	for _, size := range groupSizes {
		totalSeating += size
	}

	if totalPeople != totalSeating {
		fmt.Fprintln(stderr, "People and Seats must match:", totalPeople, "people, ", totalSeating, "seats.")
		exit(1)
	}

	seating := MingleTeams(teams, groupSizes, matches)

	for i, table := range seating {
		fmt.Fprintln(stdout, "=====================================================")
		fmt.Fprintf(stdout, "[GROUP %d] Seating Arrangement for %d people.\n", i+1, table.MaxSize)
		fmt.Fprintln(stdout, "=====================================================")
		for _, person := range table.People {
			fmt.Fprintln(stdout, person)
		}
	}

	writeS := ""
	for first, val := range matches {
		for second := range val {
			writeS += fmt.Sprintf("%s,%s\n", first, second)
		}
	}
	os.Truncate(matchesFile, 0)
	os.WriteFile(matchesFile, []byte(writeS), 0666)
}
