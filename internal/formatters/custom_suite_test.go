package formatter

import (
	"strings"

	"github.com/stretchr/testify/suite"
)

type customSuite struct {
	suite.Suite
}

func (c *customSuite) SubstringCount(expectedCount int, needle, haystack string) {
	count := strings.Count(haystack, needle)
	c.Equal(expectedCount, count, "Expected %d occurrences of '%s' in '%s', but found %d", expectedCount, needle, haystack, count)
}
