package validators

import (
	"encoding/json"
  "errors"
)

func Fathom(version string, data []byte) error {
	var parsed map[string]interface{}
  if len(data) > 100 * 1024 {
    return errors.New("Too much data.")
  }
	return json.Unmarshal(data, &parsed)
}
