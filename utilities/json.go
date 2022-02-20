package utilities

import "encoding/json"

func Stringify(object interface{}) ([]byte, error) {
	return json.Marshal(object)
}
