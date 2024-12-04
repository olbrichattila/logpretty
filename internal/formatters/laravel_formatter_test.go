package formatter

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type laravelTestSuite struct {
	customSuite
	formatter formatter
}

func TestLaravelRunner(t *testing.T) {
	suite.Run(t, new(laravelTestSuite))
}

func (t *laravelTestSuite) SetupTest() {
	t.formatter = newLaravel(
		newJSON(),
		newExtractor(),
	)
}

func (t *laravelTestSuite) TearDownTest() {
	t.formatter = nil
}

func (t *laravelTestSuite) TestisValid() {
	valid := t.formatter.isValid(invalidLine)
	t.False(valid)

	valid = t.formatter.isValid(laravelLine)
	t.True(valid)
}

func (t *laravelTestSuite) TestFormat() {
	if t.formatter.isValid(laravelLine) {
		t.SubstringCount(14, "\n", t.formatter.format())
	}
}
