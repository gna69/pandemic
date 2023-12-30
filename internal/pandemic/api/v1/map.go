package v1

import "net/http"

func GetMapHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(`{"key": "value"}`))
	}
}
