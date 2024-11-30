package formatter

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type yii2TestSuite struct {
	customSuite
	formatter formatter
}

func TestYii2Runner(t *testing.T) {
	suite.Run(t, new(yii2TestSuite))
}

func (t *yii2TestSuite) SetupTest() {
	t.formatter = newYii2(
		newJSON(),
		newExtractor(),
	)
}

func (t *yii2TestSuite) TearDownTest() {
	t.formatter = nil
}

func (t *yii2TestSuite) TestisValid() {
	valid := t.formatter.isValid(invalidLine)
	t.False(valid)

	valid = t.formatter.isValid(yii2Line)
	t.True(valid)
}

func (t *yii2TestSuite) TestFormat() {
	if t.formatter.isValid(yii2Line) {
		t.SubstringCount(5, "\n", t.formatter.format())
	}
}
