package formatter

import (
	"fmt"
	"regexp"
)

func newPHP(jsonFormatter formatter, extractor jsonExtractor) formatter {
	return &fPhp{
		jsonFormatter: jsonFormatter,
		extractor:     extractor,
	}
}

type fPhp struct {
	line          string
	jsonFormatter formatter
	extractor     jsonExtractor
}

func (f *fPhp) isValid(line string) bool {
	f.line = line
	re := regexp.MustCompile(`^\[\d{2}-\w{3}-\d{4} \d{2}:\d{2}:\d{2} [A-Z]+\] PHP \w+: .*`)

	return re.MatchString(f.line)
}

func (f *fPhp) format() string {
	re := regexp.MustCompile(`\[(.*?)\] PHP (\w+): (.*)`)
	matches := re.FindStringSubmatch(f.line)
	if len(matches) > 0 {
		timestamp := matches[1]
		logLevel := matches[2]
		message := matches[3]
		if f.jsonFormatter.isValid(message) {
			message = f.jsonFormatter.format()
		}

		extractedJsons, remaining := f.extractor.extractJSON(message)
		message = remaining
		for _, jsonStr := range extractedJsons {
			if f.jsonFormatter.isValid(jsonStr) {
				message += "\n" + blue + f.jsonFormatter.format() + reset
			}
		}

		return fmt.Sprintf(
			"---PHP----\n%sTimestamp:%s %s\n%sLog Level:%s %s\n%sMessage:%s %s\n",
			green, reset,
			timestamp,
			green, reset,
			logLevel,
			green, reset,
			message,
		)
	}

	return f.line
}
