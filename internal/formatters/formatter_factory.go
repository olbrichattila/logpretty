package formatter

import (
	"errors"
	"os"
)

var (
	errFormatterNotFound = errors.New("formatter not found")
)

func getFormatter(line string) (formatter, error) {
	argLen := len(os.Args)
	if argLen == 1 {
		for _, f := range formatterMapping {
			if f.isValid(line) {
				return f, nil
			}
		}

		return nil, errFormatterNotFound
	}

	if formatterID, ok := envFormatterMap[os.Args[1]]; ok {
		formatterMapping[formatterID].isValid(line)
		return formatterMapping[formatterID], nil
	}

	return nil, errFormatterNotFound
}
