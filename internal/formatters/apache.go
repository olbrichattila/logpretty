package formatter

import (
	"fmt"
	"regexp"
)

func newApache(jsonFormatter formatter) formatter {
	return &fApache{
		jsonFormatter: jsonFormatter,
	}
}

type fApache struct {
	line          string
	jsonFormatter formatter
}

func (f *fApache) isValid(line string) bool {
	f.line = line
	re := regexp.MustCompile(`^\S+ \S+ \S+ \[[^]]+\] ".*?" \d+ \d+`)

	return re.MatchString(f.line)
}

func (f *fApache) format() string {
	re := regexp.MustCompile(`([\d.]+) - (.*?) \[(.*?)\] "(.*?)" (\d{3}) (\d+|-)`)
	matches := re.FindStringSubmatch(f.line)
	if len(matches) > 0 {
		ip := matches[1]
		user := matches[2]
		timestamp := matches[3]
		request := matches[4]
		status := matches[5]
		bytes := matches[6]

		return fmt.Sprintf(
			"%sIP:%s %s\n%sUser:%s %s\n%sTimestamp:%s %s\n%sRequest:%s %s\n%sStatus:%s %s\n%sBytes:%s %s\n",
			Green, Reset,
			ip,
			Green, Reset,
			user,
			Green, Reset,
			timestamp,
			Green, Reset,
			request,
			Green, Reset,
			status,
			Green, Reset,
			bytes,
		)
	}

	return f.line
}
