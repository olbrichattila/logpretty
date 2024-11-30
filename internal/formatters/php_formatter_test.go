package formatter

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type phpTestSuite struct {
	customSuite
	formatter formatter
}

func TestPhpRunner(t *testing.T) {
	suite.Run(t, new(phpTestSuite))
}

func (t *phpTestSuite) SetupTest() {
	t.formatter = newPHP(
		newJSON(),
		newExtractor(),
	)
}

func (t *phpTestSuite) TearDownTest() {
	t.formatter = nil
}

func (t *phpTestSuite) TestisValid() {
	valid := t.formatter.isValid(invalidLine)
	t.False(valid)

	valid = t.formatter.isValid(phpLine)
	t.True(valid)
}

func (t *phpTestSuite) TestFormat() {
	if t.formatter.isValid(phpLine) {
		t.SubstringCount(4, "\n", t.formatter.format())
	}
}
