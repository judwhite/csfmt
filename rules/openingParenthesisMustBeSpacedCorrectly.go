package rules

import "regexp"

var openingParenthesisMustBeSpacedCorrectly = &Rule{
	Name:        "Opening parenthesis must be spaced correctly",
	Enabled:     true,
	Apply:       applyOpeningParenthesisMustBeSpacedCorrectly,
	Description: ``,
}

func applyOpeningParenthesisMustBeSpacedCorrectly(source []byte) []byte {
	spaceBetween := `(if|while|for|switch|foreach|using|\+|\-|\*|/|&|\||\^|=)`

	// Remove leading spaces
	re := regexp.MustCompile(`([\S])(\t| )([\(])`)
	for re.Match(source) {
		source = re.ReplaceAll(source, []byte("$1$3"))
	}

	// Remove trailing spaces
	re = regexp.MustCompile(`([\(])(\t| )([\S])`)
	for re.Match(source) {
		source = re.ReplaceAll(source, []byte("$1$3"))
	}

	// Add space between operators and keywords
	re = regexp.MustCompile(spaceBetween + `([\(])`)
	for re.Match(source) {
		source = re.ReplaceAll(source, []byte("$1 $2"))
	}
	return source
}
