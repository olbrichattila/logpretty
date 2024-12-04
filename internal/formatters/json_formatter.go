package formatter

import (
	"encoding/json"
	"strings"
)

func newJSON() formatter {
	return &fJSON{}
}

type fJSON struct {
	line string
	json interface{}
}

func (f *fJSON) isValid(line string) bool {
	f.line = line
	err := json.Unmarshal([]byte(line), &f.json)
	f.json = f.processAndReplaceJSONStrings(f.json)
	if err != nil {
		f.json = nil
	}

	return err == nil
}

func (f *fJSON) format() string {
	if f.json == nil {
		return f.line
	}

	// Use a custom encoder to prevent escaping
	var sb strings.Builder
	encoder := json.NewEncoder(&sb)
	encoder.SetIndent("", " ")
	encoder.SetEscapeHTML(false) // Disable HTML escaping

	err := encoder.Encode(f.json)
	if err != nil {
		return f.line
	}

	// Format the result and add color
	formatted := sb.String()
	return green + "---Json---\n" + reset + strings.ReplaceAll(f.colorize(formatted), "\\n", "\n")
}

// This function looks for JSON in the structure encoded as a string and further decode it recursively for more readability
func (f *fJSON) processAndReplaceJSONStrings(data interface{}) interface{} {
	switch v := data.(type) {
	case map[string]interface{}:
		for key, value := range v {
			v[key] = f.processAndReplaceJSONStrings(value)
		}
	case []interface{}:
		for i, value := range v {
			v[i] = f.processAndReplaceJSONStrings(value)
		}
	case string:
		var parsed interface{}
		if err := json.Unmarshal([]byte(v), &parsed); err == nil {
			return f.processAndReplaceJSONStrings(parsed)
		} else {
		}
	}
	return data
}

func (f *fJSON) colorize(s string) string {
	sb := &strings.Builder{}

	for _, char := range s {
		sb.WriteString(f.getReplacement(char))
	}

	return sb.String()
}

func (f *fJSON) getReplacement(char rune) string {
	switch char {
	case '{', '}', '(', ')':
		return blue + string(char) + reset
	case '[', ']':
		return red + string(char) + reset
	case '"', '\'', '`':
		return green + string(char) + reset
	}

	return string(char)
}
