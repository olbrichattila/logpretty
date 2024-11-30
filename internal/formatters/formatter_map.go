package formatter

var extractor = newExtractor()
var jsonFormatter = newJSON()

var formatterMapping = []formatter{
	jsonFormatter,
	newPHP(jsonFormatter, extractor),
	newApache(jsonFormatter),
	newLaravel(jsonFormatter, extractor),
	newYii2(jsonFormatter, extractor),
	newYii1(jsonFormatter, extractor),
	newGeneric(jsonFormatter),
}

var envFormatterMap = map[string]int{
	"json":    0,
	"php":     1,
	"apache":  2,
	"laravel": 3,
	"yii2":    4,
	"yii1":    5,
}
