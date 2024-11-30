package formatter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	extractJson1      = "{\"key\": 1}"
	extractJson2      = "{\"key\": 1}"
	extractNestedJson = "{\"key\": \"subKey\": [{\"v1\": 1},{\"v2\": 2}]}"
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
	blocks, remaining := t.extractor.extractJSON(extractJson1)

	t.Len(blocks, 1)
	t.Equal("", remaining)
	t.Equal(extractJson1, blocks[0])
}

func (t *jsonExtractorTestSuite) TestJsonAtTheBeginning() {
	jsonText := fmt.Sprintf("%s after", extractJson1)
	blocks, remaining := t.extractor.extractJSON(jsonText)

	t.Len(blocks, 1)
	t.Equal(" after", remaining)
	t.Equal(extractJson1, blocks[0])
}

func (t *jsonExtractorTestSuite) TestJsonAtTheEnd() {
	jsonText := fmt.Sprintf("before %s", extractJson1)
	blocks, remaining := t.extractor.extractJSON(jsonText)

	t.Len(blocks, 1)
	t.Equal("before ", remaining)
	t.Equal(extractJson1, blocks[0])
}

func (t *jsonExtractorTestSuite) TestOneJsonToExtractInTheMiddle() {
	jsonText := fmt.Sprintf("before %s after", extractJson1)
	blocks, remaining := t.extractor.extractJSON(jsonText)

	t.Len(blocks, 1)
	t.Equal("before  after", remaining)
	t.Equal(extractJson1, blocks[0])
}

func (t *jsonExtractorTestSuite) TestTwoJsonToExtractInTheMiddle() {
	jsonText := fmt.Sprintf("before %s after %s end", extractJson1, extractJson2)
	blocks, remaining := t.extractor.extractJSON(jsonText)

	t.Len(blocks, 2)
	t.Equal("before  after  end", remaining)
	t.Equal(extractJson1, blocks[0])
	t.Equal(extractJson2, blocks[1])
}

// TODO nested JSON extract does not work
// func (t *jsonExtractorTestSuite) TestExtractNested() {
// 	jsonText := fmt.Sprintf("before %s after", extractNestedJson)
// 	blocks, remaining := t.extractor.extractJSON(jsonText)

// 	t.Len(blocks, 2)
// 	t.Equal("before  after", remaining)
// 	t.Equal(extractNestedJson, blocks[0])
// }
