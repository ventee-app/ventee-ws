package utilities

import "encoding/json"

func Parse(value []byte, object *interface{}) error {
	return json.Unmarshal(value, object)
}

func Stringify(object interface{}) ([]byte, error) {
	return json.Marshal(object)
}
