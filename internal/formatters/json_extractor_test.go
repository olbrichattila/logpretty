package formatter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	extractJSON1      = "{\"key\": 1}"
	extractJSON2      = "{\"key\": 1}"
	extractNestedJSON = "{\"key\": \"subKey\": [{\"v1\": 1},{\"v2\": 2}]}"
)

type jsonExtractorTestSuite struct {
	customSuite
	extractor jsonExtractor
}

func TestJSONExtractorRunner(t *testing.T) {
	suite.Run(t, new(jsonExtractorTestSuite))
}

func (t *jsonExtractorTestSuite) SetupTest() {
	t.extractor = newExtractor()
}

func (t *jsonExtractorTestSuite) TearDownTest() {
	t.extractor = nil
}

func (t *jsonExtractorTestSuite) TestNothingToExtract() {
	jsonText := "there is no json"
	blocks, remaining := t.extractor.extractJSON(jsonText)

	t.Len(blocks, 0)
	t.Equal(jsonText, remaining)
}

func (t *jsonExtractorTestSuite) TestJsonOnly() {
	blocks, remaining := t.extractor.extractJSON(extractJSON1)

	t.Len(blocks, 1)
	t.Equal("", remaining)
	t.Equal(extractJSON1, blocks[0])
}

func (t *jsonExtractorTestSuite) TestJsonAtTheBeginning() {
	jsonText := fmt.Sprintf("%s after", extractJSON1)
	blocks, remaining := t.extractor.extractJSON(jsonText)

	t.Len(blocks, 1)
	t.Equal(" after", remaining)
	t.Equal(extractJSON1, blocks[0])
}

func (t *jsonExtractorTestSuite) TestJsonAtTheEnd() {
	jsonText := fmt.Sprintf("before %s", extractJSON1)
	blocks, remaining := t.extractor.extractJSON(jsonText)

	t.Len(blocks, 1)
	t.Equal("before ", remaining)
	t.Equal(extractJSON1, blocks[0])
}

func (t *jsonExtractorTestSuite) TestOneJsonToExtractInTheMiddle() {
	jsonText := fmt.Sprintf("before %s after", extractJSON1)
	blocks, remaining := t.extractor.extractJSON(jsonText)

	t.Len(blocks, 1)
	t.Equal("before  after", remaining)
	t.Equal(extractJSON1, blocks[0])
}

func (t *jsonExtractorTestSuite) TestTwoJsonToExtractInTheMiddle() {
	jsonText := fmt.Sprintf("before %s after %s end", extractJSON1, extractJSON2)
	blocks, remaining := t.extractor.extractJSON(jsonText)

	t.Len(blocks, 2)
	t.Equal("before  after  end", remaining)
	t.Equal(extractJSON1, blocks[0])
	t.Equal(extractJSON2, blocks[1])
}

// // TODO nested JSON extract does not work
// func (t *jsonExtractorTestSuite) TestExtractNested() {
// 	jsonText := fmt.Sprintf("before %s after", extractNestedJSON)
// 	blocks, remaining := t.extractor.extractJSON(jsonText)

// 	t.Len(blocks, 2)
// 	t.Equal("before  after", remaining)
// 	t.Equal(extractNestedJSON, blocks[0])
// }
