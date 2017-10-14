package https

import "net/http"

func EnforceTLS(h http.Handler) http.Handler {
	return enforceTLS{h}
}

type enforceTLS struct {
	http.Handler
}

func (h enforceTLS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	h.Handler.ServeHTTP(w, r)
}
