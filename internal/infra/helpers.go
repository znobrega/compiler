package infra

import "regexp"

func MatchString(pattern, letter string) bool {
	ok, err := regexp.MatchString(pattern, letter)
	if err != nil {
		// TODO REFACTOR ERROR TREATMENT
		panic(err)
	}
	return ok
}
