package formatter

import "encoding/json"

func newJSON() formatter {
	return &fJSON{}
}

type fJSON struct {
	line string
	json interface{}
}

func (f *fJSON) isValid(line string) bool {
	f.line = line
	err := json.Unmarshal([]byte(line), &f.json)
	if err != nil {
		f.json = nil
	}

	return err == nil
}

func (f *fJSON) format() string {
	if f.json == nil {
		return f.line
	}

	res, err := json.MarshalIndent(f.json, "", " ")
	if err != nil {
		return f.line
	}

	return "Json\n" + string(res)
}
