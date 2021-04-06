package infra

import "regexp"

func MatchString(expression, letter string) bool {
	ok, err := regexp.MatchString(expression, letter)
	if err != nil {
		// TODO REFACTOR ERROR TREATMENT
		panic(err)
	}
	return ok
}
