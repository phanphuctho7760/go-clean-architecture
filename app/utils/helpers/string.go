package helpers

import "encoding/json"

func ToString(data interface{}) (res string) {
	resByte, err := json.Marshal(data)
	if err != nil {
		return
	}
	return string(resByte)
}
