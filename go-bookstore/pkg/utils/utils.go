// package utils

// import (
// 	"encoding/json"
// 	"io/ioutil"
// 	"net/http"
// )

//	func ParseBody(r *http.Request, x interface{}) {
//		if body, err := ioutil.ReadAll(r.Body); err == nil {
//			if err := json.Unmarshal([]byte(body), x); err != nil {
//				return
//			}
//		}
//	}
package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal(body, x); err != nil {
			return
		}
	}
}
