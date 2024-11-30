package formatter

import (
	"regexp"
	"strings"
)

func newGeneric(jsonFormatter formatter) formatter {
	return &fGeneric{
		jsonFormatter: jsonFormatter,
	}
}

type fGeneric struct {
	line          string
	jsonFormatter formatter
}

func (f *fGeneric) isValid(line string) bool {
	f.line = line

	return true
}

func (f *fGeneric) format() string {
	return f.universalSplitWithEscapes()
}

func (f *fGeneric) universalSplitWithEscapes() string {
	sb := &strings.Builder{}
	sb.WriteString("Generic\n")
	re := regexp.MustCompile(`"([^"\\]*(\\.)?)*"|\[[^\]]*\]|\\.|[^\s]+`)
	matches := re.FindAllString(f.line, -1)

	for i, block := range matches {
		unescapedBlock := f.unescape(block)

		if i > 0 {
			sb.WriteRune('\n')
		}

		if f.jsonFormatter.isValid(unescapedBlock) {
			sb.WriteString(green)
			sb.WriteString(f.jsonFormatter.format())
			sb.WriteString(reset)
		} else {
			sb.WriteString(unescapedBlock)
		}
	}

	return sb.String()
}

func (f *fGeneric) unescape(input string) string {
	replacer := strings.NewReplacer(
		`\"`, `"`,
		`\\`, `\`,
		`\t`, "\t",
		`\n`, "\n",
		`\r`, "\r",
		`\ `, " ", // Escaped space
	)
	return replacer.Replace(input)
}
