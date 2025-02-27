package lexer

import (
	"fmt"
	"regexp"
)

func validateModelName(name string) error {
	camelCaseRegex := regexp.MustCompile(`^[A-Z][a-z]*(?:[A-Z][a-z]*)*$`)

	if camelCaseRegex.MatchString(name) {
		return nil
	}

	return fmt.Errorf("lexer::error -> modelName %s is not in CamelCase", name)
}
