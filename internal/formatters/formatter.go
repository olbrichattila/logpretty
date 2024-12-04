// Package formatter contains different log formatting algorithms
package formatter

import (
	"fmt"
	"os"
	"strings"
)

const (
	reset = "\033[0m"
	green = "\033[32m"
	red   = "\033[31m"
	blue  = "\033[34m"
)

type formatter interface {
	isValid(line string) bool
	format() string
}

// Run runs formatter for a specific line
func Run(line string) string {
	formatter, err := getFormatter(line)
	if err != nil {
		return err.Error()
	}

	return formatter.format()
}

// Validate formatter
func Validate() error {
	if len(os.Args) == 1 {
		return nil
	}

	_, ok := envFormatterMap[os.Args[1]]
	if !ok {
		var types []string
		for formatterType := range envFormatterMap {
			types = append(types, formatterType)
		}

		return fmt.Errorf("cannot find formatter " + os.Args[1] + " valid types: " + strings.Join(types, ", "))
	}

	return nil
}
