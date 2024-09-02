package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, post interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		err = json.Unmarshal([]byte(body), &post)
		if err != nil {
			return
		}
	}
}
