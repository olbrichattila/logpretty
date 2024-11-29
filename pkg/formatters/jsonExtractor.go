package formatter

import (
	"encoding/json"
	"strings"
)

func newExtractor() jsonExtractor {
	return &jExtractor{}
}

type jsonExtractor interface {
	extractJSON(input string) ([]string, string)
}

type jExtractor struct {
}

func (*jExtractor) extractJSON(input string) ([]string, string) {
	var jsonBlocks []string
	var current strings.Builder
	var remaining strings.Builder
	depth := 0
	inEscape := false // Track if the current character is escaped
	start := -1       // Tracks the start index of the JSON block

	for i, char := range input {
		if inEscape {
			// Add escaped character to the current block and skip depth tracking
			current.WriteRune(char)
			inEscape = false
			continue
		}

		// Check for escape character
		if char == '\\' && i+1 < len(input) {
			nextChar := input[i+1]
			if nextChar == '{' || nextChar == '}' || nextChar == '\\' {
				current.WriteRune(char) // Write the escape character
				inEscape = true
				continue
			}
		}

		switch char {
		case '{':
			if depth == 0 {
				current.Reset()
				start = i // Mark the start of the JSON block
			}
			current.WriteRune(char)
			depth++
		case '}':
			if depth > 0 {
				current.WriteRune(char)
				depth--
				if depth == 0 {
					// Validate the JSON block
					block := current.String()
					var temp interface{}
					if json.Unmarshal([]byte(block), &temp) == nil {
						jsonBlocks = append(jsonBlocks, block)
						start = -1 // Reset start after successful extraction
					} else {
						remaining.WriteString(input[start : i+1]) // Add unvalidated block back to remaining string
					}
				}
			}
		default:
			if depth > 0 {
				current.WriteRune(char)
			} else {
				remaining.WriteRune(char)
			}
		}
	}

	// Add remaining unprocessed part of the string if depth > 0
	if depth > 0 && start >= 0 {
		remaining.WriteString(current.String())
	}

	return jsonBlocks, remaining.String()
}
