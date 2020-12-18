package day6

import (
	"log"

	"github.com/wincus/adventofcode/internal/common"
)

type question bool

type form [26]question

type group []form

// Solve returns the solution for day6 problem
func Solve(s []string, p common.Part) int {

	var counter int
	var forms []form

	groups, err := parseGroups(s)

	if err != nil {
		log.Printf("could not parse forms: %v", err)
	}

	for _, group := range groups {
		forms = append(forms, mergeGroup(group, p))
	}

	for _, f := range forms {
		for i := 0; i < len(f); i++ {
			if f[i] {
				counter++
			}
		}
	}

	return counter
}

// merge group will get slice of forms belonging to a
// group, and merge it in a single form using for Part1
// an OR operation ( an answer is counted if anyone has answered )
// and an AND operation for Part2 ( an answer is counted only if everyone
// has answered it )
func mergeGroup(g group, p common.Part) form {

	var mergedForm form

	if p == common.Part1 {

		for i := 0; i < len(mergedForm); i++ {

			mergedForm[i] = false // not really needed as false is the
			// boolean default value. Still I liked
			// the symmetry so Ill keep it :)

			for _, f := range g {
				if f[i] {
					mergedForm[i] = true
				}
			}
		}
	}

	if p == common.Part2 {

		for i := 0; i < len(mergedForm); i++ {

			mergedForm[i] = true

			for _, f := range g {
				if !f[i] {
					mergedForm[i] = false
				}
			}
		}
	}

	return mergedForm
}

func parseGroups(s []string) ([]group, error) {

	var groups []group
	var forms []form

	// ensure last group is parsed
	s = append(s, "")

	for _, line := range s {

		switch len(line) {
		case 0:

			// ensures we process 1 or more
			// people's forms
			if len(forms) > 0 {
				groups = append(groups, forms)
			}

			forms = nil

		default:
			form, err := parseForm(line)

			if err != nil {
				log.Printf("Could not parse form: %v", err)
			}

			forms = append(forms, form)
		}
	}

	return groups, nil
}

func parseForm(s string) (form, error) {

	var f form

	for _, c := range s {
		f[int(c)-97] = true
	}

	return f, nil

}
