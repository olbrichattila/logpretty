package formatter

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type jsonTestSuite struct {
	customSuite
	formatter formatter
}

func TestJSONRunner(t *testing.T) {
	suite.Run(t, new(jsonTestSuite))
}

func (t *jsonTestSuite) SetupTest() {

	t.formatter = newJSON()
}

func (t *jsonTestSuite) TearDownTest() {
	t.formatter = nil
}

func (t *jsonTestSuite) TestisValid() {
	valid := t.formatter.isValid("Invalid Json")
	t.False(valid)

	valid = t.formatter.isValid("{\"key\": 100}")
	t.True(valid)
}

func (t *jsonTestSuite) TestFormat() {
	t.formatter.isValid("{\"key\": 100, \"sub\": [{\"a\": 5, \"b\": 6}]}")
	formatted := t.formatter.format()

	t.SubstringCount(9, "\n", formatted)
}
