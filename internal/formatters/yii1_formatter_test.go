package formatter

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type yii1TestSuite struct {
	customSuite
	formatter formatter
}

func TestYii1Runner(t *testing.T) {
	suite.Run(t, new(yii1TestSuite))
}

func (t *yii1TestSuite) SetupTest() {
	t.formatter = newYii1(
		newJSON(),
		newExtractor(),
	)
}

func (t *yii1TestSuite) TearDownTest() {
	t.formatter = nil
}

func (t *yii1TestSuite) TestisValid() {
	valid := t.formatter.isValid(invalidLine)
	t.False(valid)

	valid = t.formatter.isValid(yii1Line)
	t.True(valid)
}

func (t *yii1TestSuite) TestFormat() {
	if t.formatter.isValid(yii1Line) {
		t.SubstringCount(5, "\n", t.formatter.format())
	}
}
