package health

import "net/http"

func Check(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}
