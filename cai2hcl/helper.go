package cai2hcl

import (
	"encoding/json"
	"strings"
)

func parseFieldValue(str string, field string) string {
	strList := strings.Split(str, "/")
	for ix, item := range strList {
		if item == field && ix+1 < len(strList) {
			return strList[ix+1]
		}
	}
	return ""
}

// decodeJSON decodes the map object into the target struct.
func decodeJSON(data map[string]interface{}, v interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	return nil
}
