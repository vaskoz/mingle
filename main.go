package main

import (
	"fmt"
	"os"
	"strconv"
)

var (
	stderr = os.Stderr
	stdout = os.Stdout
	exit   = os.Exit
)

func main() {
	size := os.Getenv("MINGLE_SIZE")
	if size == "" {
		fmt.Fprintln(stderr, "must specify a positive non-zero integer for group size")
		exit(1)
	}

	sizeI, err := strconv.Atoi(size)
	if err != nil {
		fmt.Fprintln(stderr, "invalid group size", size)
		exit(1)
	}

	peopleDir := os.Getenv("MINGLE_DIR")
	if peopleDir == "" {
		fmt.Fprintln(stderr, "must set env var MINGLE_DIR")
		exit(1)
	}

	files, err := os.ReadDir(peopleDir)
	if err != nil {
		fmt.Fprintln(stderr, "failed to read directory given")
		exit(1)
	}

	var people []Person

	for _, file := range files {
		fullPath := fmt.Sprintf("%s/%s", peopleDir, file.Name())
		b, _ := os.ReadFile(fullPath)
		fileS := string(b)
		p := ExtractPerson(file.Name(), fileS)
		people = append(people, p)
	}

	mingles := Greedy(people[0], sizeI)

	fmt.Fprintln(stdout, mingles)
}
