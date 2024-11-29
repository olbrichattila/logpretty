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
