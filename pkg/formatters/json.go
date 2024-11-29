package formatter

import "encoding/json"

func newJSON() formatter {
	return &fJson{}
}

type fJson struct {
	line string
	json interface{}
}

func (f *fJson) isValid(line string) bool {
	f.line = line
	err := json.Unmarshal([]byte(line), &f.json)

	return err == nil
}

func (f *fJson) format() string {
	res, err := json.MarshalIndent(f.json, "", " ")
	if err != nil {
		return f.line
	}

	return string(res)
}
