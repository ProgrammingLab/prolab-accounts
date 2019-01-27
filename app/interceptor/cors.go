package interceptor

import "net/http"

// CORSMiddleware repesents CORS middleware
func CORSMiddleware(hander http.Handler) http.Handler {
	return &corsHander{
		hander: hander,
	}
}

type corsHander struct {
	hander http.Handler
}

func (h *corsHander) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	header := w.Header()
	header.Add("Access-Control-Allow-Origin", "*")
	header.Add("Access-Control-Allow-Methods", "GET,HEAD,POST,PUT,DELETE,PATCH")
	header.Add("Access-Control-Allow-Headers", "Authorization")
	h.hander.ServeHTTP(w, req)
}
