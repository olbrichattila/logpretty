package formatter

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type genericTestSuite struct {
	customSuite
	formatter formatter
}

func TestGenericRunner(t *testing.T) {
	suite.Run(t, new(genericTestSuite))
}

func (t *genericTestSuite) SetupTest() {
	t.formatter = newGeneric(
		newJSON(),
	)
}

func (t *genericTestSuite) TearDownTest() {
	t.formatter = nil
}

func (t *genericTestSuite) TestisValid() {
	// it always returns true
	valid := t.formatter.isValid("")
	t.True(valid)
}

func (t *genericTestSuite) TestFormat() {
	if t.formatter.isValid(genericLine) {
		t.SubstringCount(8, "\n", t.formatter.format())
	}
}
