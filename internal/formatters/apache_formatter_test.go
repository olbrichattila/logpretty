package formatter

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type apacheTestSuite struct {
	customSuite
	formatter formatter
}

func TestApacheRunner(t *testing.T) {
	suite.Run(t, new(apacheTestSuite))
}

func (t *apacheTestSuite) SetupTest() {
	t.formatter = newApache(
		newJSON(),
	)
}

func (t *apacheTestSuite) TearDownTest() {
	t.formatter = nil
}

func (t *apacheTestSuite) TestisValid() {
	valid := t.formatter.isValid(invalidLine)
	t.False(valid)

	valid = t.formatter.isValid(apacheLine)
	t.True(valid)
}

func (t *apacheTestSuite) TestFormat() {
	if t.formatter.isValid(apacheLine) {
		t.SubstringCount(7, "\n", t.formatter.format())
	}
}
