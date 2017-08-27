/**
 * Created by jbisa on 8/20/17.
 */

package handlers

import (
	"net/http"
)

type Response struct {
	Items interface{}
	Err error
}

type Ad struct {
	Id string
	Url string
	Width string
	Height string
	Description string
	Creative string
}

// Set access control response headers, only when the origin is known.
func setAccessControlResponseHeaders (w http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
}
