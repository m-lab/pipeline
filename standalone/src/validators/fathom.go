package validators

import (
	"encoding/json"
)

func Fathom(version string, data []byte) error {
	var parsed map[string]interface{}
	return json.Unmarshal(data, &parsed)
}
