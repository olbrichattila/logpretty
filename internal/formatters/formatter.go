// Package formatter contains different log formatting algorithms
package formatter

const (
	reset = "\033[0m"
	green = "\033[32m"
	blue  = "\033[34m"
)

type formatter interface {
	isValid(line string) bool
	format() string
}

// Run runs formatter for a specific line
func Run(line string) string {
	for _, f := range formatterMapping {
		if f.isValid(line) {
			return f.format()
		}
	}

	// If could not be formatted, return original
	return line
}
