package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for idx, line := range lines {
		q := line[0]
		a := line[1]
		p := problem{q, a}
		ret[idx] = p
	}
	return ret
}

func main() {
	filename := "problems.csv"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error reading %s", filename)
	}
	defer file.Close()
	csv := csv.NewReader(file)
	lines, err := csv.ReadAll()
	problems := parseLines(lines)
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s\n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			fmt.Println("right!")
			correct++
		} else {
			fmt.Println("WRONG!")
		}
	}
	fmt.Printf("You got %d out of %d correct\n", correct, len(lines))

}
