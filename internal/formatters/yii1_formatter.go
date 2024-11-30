package formatter

import (
	"fmt"
	"regexp"
)

func newYii1(jsonFormatter formatter, extractor jsonExtractor) formatter {
	return &fYii1{
		jsonFormatter: jsonFormatter,
		extractor:     extractor,
	}
}

type fYii1 struct {
	line          string
	jsonFormatter formatter
	extractor     jsonExtractor
}

func (f *fYii1) isValid(line string) bool {
	f.line = line
	re := regexp.MustCompile(`^\d{4}/\d{2}/\d{2} \d{2}:\d{2}:\d{2} \[\w+\] \[.*?\] .*`)

	return re.MatchString(f.line)
}

func (f *fYii1) format() string {
	re := regexp.MustCompile(`(\d{4}/\d{2}/\d{2} \d{2}:\d{2}:\d{2}) \[(\w+)\] \[(.*?)\] (.*)`)
	matches := re.FindStringSubmatch(f.line)
	if len(matches) > 0 {
		timestamp := matches[1]
		logLevel := matches[2]
		category := matches[3]
		message := matches[4]
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
			"---Yii1---\n%sTimestamp:%s %s\n%sLog Level:%s %s\n%sCategory:%s %s\n%sMessage:%s %s\n",
			green, reset,
			timestamp,
			green, reset,
			logLevel,
			green, reset,
			category,
			green, reset,
			message,
		)
	}

	return f.line
}
