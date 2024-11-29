package formatter

import (
	"fmt"
	"regexp"
)

func newLaravel(jsonFormatter formatter, extractor jsonExtractor) formatter {
	return &fLaravel{
		jsonFormatter: jsonFormatter,
		extractor:     extractor,
	}
}

type fLaravel struct {
	line          string
	jsonFormatter formatter
	extractor     jsonExtractor
}

func (f *fLaravel) isValid(line string) bool {
	f.line = line
	re := regexp.MustCompile(`^\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\] [a-zA-Z0-9\._]+: .*`)

	return re.MatchString(f.line)
}

func (f *fLaravel) format() string {
	re := regexp.MustCompile(`\[(.*?)\] (\w+)\.(\w+): (.*)`)
	matches := re.FindStringSubmatch(f.line)
	if len(matches) > 0 {
		timestamp := matches[1]
		environment := matches[2]
		level := matches[3]
		message := matches[4]
		if f.jsonFormatter.isValid(message) {
			message = f.jsonFormatter.format()
		}

		extractedJsons, remaining := f.extractor.extractJSON(message)
		message = remaining
		for _, jsonStr := range extractedJsons {
			if f.jsonFormatter.isValid(jsonStr) {
				message += "\n" + Blue + f.jsonFormatter.format() + Reset
			}
		}

		return fmt.Sprintf(
			"%sTimestamp:%s %s\n%sEnvironment:%s %s\n%sLevel:%s %s\n%sMessage:%s %s\n",
			Green, Reset,
			timestamp,
			Green, Reset,
			environment,
			Green, Reset,
			level,
			Green, Reset,
			message,
		)
	}

	return f.line
}
