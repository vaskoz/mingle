package main

import "strings"

type Person struct {
	Name  string
	Prefs []string
}

func ExtractPerson(name, file string) Person {
	lines := strings.Split(file, "\n")
	p := Person{name, nil}
	p.Prefs = append(p.Prefs, lines...)
	return p
}
