package day4

import (
	"log"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/wincus/adventofcode/internal/common"
)

type passport struct {
	BYR string `validate:"required"`
	IYR string `validate:"required"`
	EYR string `validate:"required"`
	HGT string `validate:"required"`
	HCL string `validate:"required"`
	ECL string `validate:"required"`
	PID string `validate:"required"`
	CID string `validate:""`
}

var validate *validator.Validate

// Solve returns the solution for day3 problem
func Solve(s []string, p common.Part) int {

	var counter int

	validate = validator.New()

	passports := parse(s, p)

	for _, pass := range passports {
		if pass.validate() {
			counter++
		}
	}

	return counter

}

func parse(s []string, p common.Part) []passport {

	var passports []passport
	var tokens []string

	// we append a new line to the input to ensure last
	// passport is processed
	s = append(s, "")

	for _, line := range s {

		switch len(line) {
		case 0:
			pass := parseLine(strings.Join(tokens, " "), p)
			passports = append(passports, pass)
			tokens = nil // empty slice

		default:
			tokens = append(tokens, line)
		}
	}

	return passports
}

func parseLine(s string, p common.Part) passport {

	var pass passport

	tokens := strings.Split(s, " ")

	for _, t := range tokens {

		if len(t) == 0 {
			continue
		}

		kv := strings.Split(t, ":")

		switch kv[0] {
		case "byr", "BYR":
			pass.BYR = validateBirthYear(kv[1], p)
		case "iyr", "IYR":
			pass.IYR = validateIssueYear(kv[1], p)
		case "eyr", "EYR":
			pass.EYR = validateExpirationYear(kv[1], p)
		case "hgt", "HGT":
			pass.HGT = validateHeight(kv[1], p)
		case "hcl", "HCL":
			pass.HCL = validateHairColor(kv[1], p)
		case "ecl", "ECL":
			pass.ECL = validateEyeColor(kv[1], p)
		case "pid", "PID":
			pass.PID = validatePassportID(kv[1], p)
		case "cid", "CID":
			pass.CID = validateCountryID(kv[1], p)
		default:
			log.Printf("unknown token found: %v", kv[0])
		}
	}

	return pass
}

func (p passport) validate() bool {

	err := validate.Struct(p)

	if err != nil {
		return false
	}

	return true
}

func validateBirthYear(s string, p common.Part) string {

	if p == common.Part1 {
		return s
	}

	// extra validation for Part2
	if p == common.Part2 {
		// check this is a valida number
		v, err := strconv.Atoi(s)

		if err != nil {
			return ""
		}

		err = validate.Var(v, "required,min=1920,max=2002")

		if err != nil {
			return ""
		}
	}

	return s

}

func validateIssueYear(s string, p common.Part) string {

	if p == common.Part1 {
		return s
	}

	// extra validation for Part2
	if p == common.Part2 {
		// check this is a valida number
		v, err := strconv.Atoi(s)

		if err != nil {
			return ""
		}

		err = validate.Var(v, "required,min=2010,max=2020")

		if err != nil {
			return ""
		}
	}

	return s

}

func validateExpirationYear(s string, p common.Part) string {
	if p == common.Part1 {
		return s
	}

	// extra validation for Part2
	if p == common.Part2 {
		// check this is a valida number
		v, err := strconv.Atoi(s)

		if err != nil {
			return ""
		}

		err = validate.Var(v, "required,min=2020,max=2030")

		if err != nil {
			return ""
		}
	}

	return s
}

func validateHeight(s string, p common.Part) string {

	if p == common.Part1 {
		return s
	}

	if p == common.Part2 {
		if strings.HasSuffix(s, "cm") {

			v, err := strconv.Atoi(strings.TrimSuffix(s, "cm"))

			if err != nil {
				return ""
			}

			err = validate.Var(v, "required,min=150,max=193")

			if err != nil {
				return ""
			}

			return s
		}

		if strings.HasSuffix(s, "in") {

			v, err := strconv.Atoi(strings.TrimSuffix(s, "in"))

			if err != nil {
				return ""
			}

			err = validate.Var(v, "required,min=59,max=76")

			if err != nil {
				return ""
			}

			return s
		}
	}

	return ""
}

func validateHairColor(s string, p common.Part) string {

	if p == common.Part1 {
		return s
	}

	if p == common.Part2 {
		err := validate.Var(s, "hexcolor")

		if err != nil {
			return ""
		}
	}

	return s
}

func validateEyeColor(s string, p common.Part) string {
	if p == common.Part1 {
		return s
	}

	if p == common.Part2 {
		err := validate.Var(s, "oneof=amb blu brn gry grn hzl oth")

		if err != nil {
			return ""
		}
	}

	return s
}

func validatePassportID(s string, p common.Part) string {
	if p == common.Part1 {
		return s
	}

	if p == common.Part2 {
		err := validate.Var(s, "numeric,len=9")

		if err != nil {
			return ""
		}
	}

	return s

}

func validateCountryID(s string, p common.Part) string {
	return s
}
